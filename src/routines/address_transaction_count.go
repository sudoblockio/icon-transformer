package routines

import (
	"time"

	"go.uber.org/zap"

	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/models"
	"github.com/sudoblockio/icon-transformer/redis"
)

func addressTransactionCountRoutine() {
	countRoutineCron(addressTransactionCountExec)
}

func addressTransactionCountExec() error {
	// Get all keys
	keys, err := redis.GetRedisClient().GetAllKeys(config.Config.RedisKeyPrefix + "transaction_regular_count_by_address_*")
	if err != nil {
		zap.S().Warn(
			"Routine=AddressTransactionCount",
			" Step=", "get redis keys",
			" Error=", err.Error(),
		)
		time.Sleep(1 * time.Second)
		return err
	}

	for _, key := range keys {
		addressString := key[len(config.Config.RedisKeyPrefix+"transaction_regular_count_by_address_"):]
		count, err := crud.GetTransactionCrud().CountRegularByAddress(addressString)
		if err != nil {
			zap.S().Warn(
				"Routine=AddressTransactionCount",
				" Address=", addressString,
				" Error=", err.Error(),
			)
			return err
		}

		address := &models.Address{
			Address:          addressString,
			TransactionCount: count,
		}

		crud.GetAddressCrud().UpsertOneCols(address, []string{"address", "transaction_count"})
	}
	return nil
}
