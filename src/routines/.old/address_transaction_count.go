package _old

import (
	"github.com/sudoblockio/icon-transformer/routines"
	"time"

	"go.uber.org/zap"

	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/models"
	"github.com/sudoblockio/icon-transformer/redis"
)

func addressTransactionCountRoutine() {
	routines.countRoutineCron(addressTransactionCountExec)
}

func addressTransactionCountExec() error {
	zap.S().Info("Starting addressTransactionCountExec routine.")

	// TODO: RM the redis as it makes no sense to go from redis to PG...
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

//func addressTransactionCountExec() error {
//	// Get all keys
//	keys, err := redis.GetRedisClient().GetAllKeys(config.Config.RedisKeyPrefix + "transaction_regular_count_by_address_*")
//	if err != nil {
//		zap.S().Warn(
//			"Routine=AddressTransactionCount",
//			" Step=", "get redis keys",
//			" Error=", err.Error(),
//		)
//		time.Sleep(1 * time.Second)
//		return err
//	}
//
//	for _, key := range keys {
//		addressString := key[len(config.Config.RedisKeyPrefix+"transaction_regular_count_by_address_"):]
//		count, err := crud.GetTransactionCrud().CountRegularByAddress(addressString)
//		if err != nil {
//			zap.S().Warn(
//				"Routine=AddressTransactionCount",
//				" Address=", addressString,
//				" Error=", err.Error(),
//			)
//			return err
//		}
//
//		address := &models.Address{
//			Address:          addressString,
//			TransactionCount: count,
//		}
//
//		crud.GetAddressCrud().UpsertOneCols(address, []string{"address", "transaction_count"})
//	}
//	return nil
//}
