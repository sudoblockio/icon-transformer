package _old

import (
	"errors"
	"github.com/sudoblockio/icon-transformer/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/crud"
)

func initAddressCount() {
	for i := 0; i <= config.Config.RoutinesNumWorkers; i++ {
		go initAddressCounts(i)
	}
}

func initAddressCounts(worker_id int) {
	// Loop through all addresses
	skip := worker_id * config.Config.RoutinesBatchSize
	limit := config.Config.RoutinesBatchSize
	for {
		addresses, err := crud.GetAddressCrud().SelectMany(limit, skip)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			break
		} else if err != nil {
			zap.S().Fatal(err.Error())
		}
		if len(*addresses) == 0 {
			break
		}

		zap.S().Info("Routine=AddressCount", " - Processing ", skip, " addresses...")
		for _, address := range *addresses {

			// Regular Transaction Count //
			count, err := crud.GetTransactionCrud().CountRegularByAddress(address.Address)
			if err != nil {
				zap.S().Warn("Routine=TxCount, Address=", address.Address, " - Error: ", err.Error())
				continue
			}
			address.TransactionCount = count
			err = redis.GetRedisClient().SetCount(
				config.Config.RedisKeyPrefix+"transaction_regular_count_by_address_"+address.Address,
				count,
			)
			if err != nil {
				zap.S().Warn("Routine=TxCount, Address=", address.Address, " - Error: ", err.Error())
				continue
			}

			// Internal Transaction Count //
			count, err = crud.GetTransactionCrud().CountInternalByAddress(address.Address)
			if err != nil {
				zap.S().Warn("Routine=TxInternalCount, Address=", address.Address, " - Error: ", err.Error())
				continue
			}
			address.TransactionCount = count
			err = redis.GetRedisClient().SetCount(
				config.Config.RedisKeyPrefix+"transaction_internal_count_by_address_"+address.Address,
				count,
			)
			if err != nil {
				zap.S().Warn("Routine=TxInternalCount, Address=", address.Address, " - Error: ", err.Error())
				continue
			}

			// Token Transfer Count //
			count, err = crud.GetTokenTransferCrud().CountByAddress(address.Address)
			if err != nil {
				zap.S().Warn("Routine=TokenTxCount, Address=", address.Address, " - Error: ", err.Error())
				continue
			}
			address.TokenTransferCount = count
			err = redis.GetRedisClient().SetCount(
				config.Config.RedisKeyPrefix+"token_transfer_count_by_address_"+address.Address,
				count,
			)
			if err != nil {
				zap.S().Warn("Routine=TokenTxCount, Address=", address.Address, " - Error: ", err.Error())
				continue
			}

			// Log Count //
			count, err = crud.GetLogCrud().CountLogsByAddress(address.Address)
			if err != nil {
				zap.S().Warn("Routine=LogCount, Address=", address.Address, " - Error: ", err.Error())
				continue
			}
			address.TokenTransferCount = count
			err = redis.GetRedisClient().SetCount(
				config.Config.RedisKeyPrefix+"log_count_by_address_"+address.Address,
				count,
			)
			if err != nil {
				zap.S().Warn("Routine=LogCount, Address=", address.Address, " - Error: ", err.Error())
				continue
			}

			crud.GetAddressCrud().UpsertOneCols(&address, []string{"address", "transaction_count", "transaction_internal_count", "log_count", "token_transfer_count"})
		}

		skip += config.Config.RoutinesBatchSize * config.Config.RoutinesNumWorkers
	}
	zap.S().Info("Routine=AddressCount - Completed routine", config.Config.RoutinesSleepDuration.String(), "...")
}
