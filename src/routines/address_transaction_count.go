package routines

import (
	"strconv"
	"time"

	"go.uber.org/zap"

	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/models"
	"github.com/sudoblockio/icon-transformer/redis"
)

func addressTransactionCountRoutine() {

	// Loop every duration
	for {

		// Get all keys
		keys, err := redis.GetRedisClient().GetAllKeys(config.Config.RedisKeyPrefix + "transaction_regular_count_by_address_*")
		if err != nil {
			zap.S().Warn(
				"Routine=AddressTransactionCount",
				" Step=", "get redis keys",
				" Error=", err.Error(),
			)
			time.Sleep(1 * time.Second)
			continue
		}

		for _, key := range keys {
			valueString, err := redis.GetRedisClient().GetValue(key)
			if err != nil {
				zap.S().Warn(
					"Routine=AddressTransactionCount",
					" Step=", "get redis value",
					" Key=", key,
					" Error=", err.Error(),
				)
				continue
			}

			addressString := key[len(config.Config.RedisKeyPrefix+"transaction_regular_count_by_address_"):]
			valueInt, _ := strconv.Atoi(valueString)

			address := &models.Address{
				Address:          addressString,
				TransactionCount: int64(valueInt),
			}

			crud.GetAddressCrud().LoaderChannel <- address
		}

		zap.S().Info("Routine=AddressTransactionCount - Completed routine, sleeping ", config.Config.RoutinesSleepDuration.String(), "...")
		time.Sleep(config.Config.RoutinesSleepDuration)
	}
}
