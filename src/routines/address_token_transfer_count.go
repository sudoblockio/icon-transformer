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

func addressTokenTransferCountRoutine() {

	// Loop every duration
	for {

		// Get all keys
		keys, err := redis.GetRedisClient().GetAllKeys(config.Config.RedisKeyPrefix + "token_transfer_count_by_address_*")
		if err != nil {
			zap.S().Warn(
				"Routine=AddressTokenTransferCount",
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
					"Routine=AddressTokenTransferCount",
					" Step=", "get redis value",
					" Key=", key,
					" Error=", err.Error(),
				)
				continue
			}

			addressString := key[len(config.Config.RedisKeyPrefix+"token_transfer_count_by_address_"):]
			valueInt, _ := strconv.Atoi(valueString)

			address := &models.Address{
				Address:            addressString,
				TokenTransferCount: int64(valueInt),
			}

			crud.GetAddressCrud().UpsertOneCols(address, []string{"address", "token_transfer_count"})
		}

		zap.S().Info("Routine=AddressTokenTransferCount - Completed routine, sleeping ", config.Config.RoutinesSleepDuration.String(), "...")
		time.Sleep(config.Config.RoutinesSleepDuration)
	}
}
