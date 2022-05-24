package routines

import (
	"errors"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/redis"
)

// Wrapper for generic count routines
func countRoutineCron(routine func() error) {
	for {
		err := routine()
		if err != nil {
			time.Sleep(time.Second * 10)
			continue
		}
		zap.S().Info("Completed routine, sleeping...")
		time.Sleep(config.Config.RoutinesSleepDuration)
	}
}

// Takes a crud count method, calls it, takes the count and puts it into a redis countKey.
func crudCountSetRedis(c func() (int64, error), countKey string) error {
	count, err := c()
	if err != nil {
		// Postgres error
		zap.S().Warn(err)
		return err
	}
	err = redis.GetRedisClient().SetCount(countKey, count)
	if err != nil {
		// Redis error
		zap.S().Warn(err)
		return err
	}
	return nil
}

func blockCountRoutine() {
	countRoutineCron(blockCountExec)
}

// Count blocks from PG and set in redis
func blockCountExec() error {
	err := crudCountSetRedis(
		crud.GetBlockIndexCrud().Count,
		config.Config.RedisKeyPrefix+"block_count",
	)
	if err != nil {
		zap.S().Warn(err.Error())
		return err
	}
	return nil
}

func transactionRegularCountRoutine() {
	countRoutineCron(transactionRegularCountExec)
}

// Count transactions from PG and set in redis
func transactionRegularCountExec() error {
	zap.S().Info("Starting transaction_regular_count routine...")
	err := crudCountSetRedis(
		crud.GetTransactionCrud().CountRegular,
		config.Config.RedisKeyPrefix+"transaction_regular_count",
	)
	if err != nil {
		zap.S().Warn(err.Error())
		return err
	}
	return nil
}

//func transactionInternalCountRoutine() {
//	countRoutineCron(transactionInternalCountExec)
//}
//
//// TODO: This is currently not used in API as there is no table with only internal Txs
//// Count transactions from PG and set in redis
//func transactionInternalCountExec() error {
//	err := crudCountSetRedis(
//		crud.GetTransactionCrud().CountInternal,
//		config.Config.RedisKeyPrefix+"transaction_internal_count",
//	)
//	if err != nil {
//		zap.S().Warn(err.Error())
//		return err
//	}
//	return nil
//}

func tokenTransferCountRoutine() {
	countRoutineCron(tokenTransferCountExec)
}

// Count token_transfers from PG and set in redis
func tokenTransferCountExec() error {
	zap.S().Info("Starting transaction_regular_count routine...")
	err := crudCountSetRedis(
		crud.GetTokenTransferCrud().Count,
		config.Config.RedisKeyPrefix+"token_transfer_count",
	)
	if err != nil {
		zap.S().Warn(err.Error())
		return err
	}
	return nil
}

// TODO: Mistake / didn't realize this is already done in another routine
//  https://github.com/sudoblockio/icon-transformer/issues/22
// Takes a crud count method, calls it, takes the count and puts it into a redis countKey.
func crudCountSetRedisByAddress(addresses []string, c func(a string) (int64, error), countKey string) error {
	for _, address := range addresses {
		count, err := c(address)
		if err != nil {
			// Postgres error
			zap.S().Warn(err)
			return err
		}
		err = redis.GetRedisClient().SetCount(countKey, count)
		if err != nil {
			// Redis error
			zap.S().Warn(err)
			return err
		}
	}
	return nil
}

func byAddressCountRoutine() {
	countRoutineCron(byAddressCountExec)
}

// Count by address from PG and set in redis
func byAddressCountExec() error {
	zap.S().Info("Starting transaction_regular_count_by_address_* routine...")

	skip := 0
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
		zap.S().Info("Routine=AddressBalance", " - Processing ", len(*addresses), " addresses...")
		for _, address := range *addresses {
			err = redis.GetRedisClient().SetCount(
				config.Config.RedisKeyPrefix+"transaction_regular_count_by_address_"+address.Address,
				address.TransactionCount,
			)
			// Internal transfers
			err = redis.GetRedisClient().SetCount(
				config.Config.RedisKeyPrefix+"transaction_internal_count_by_address_"+address.Address,
				address.TransactionInternalCount,
			)
			// Token transfers
			err = redis.GetRedisClient().SetCount(
				config.Config.RedisKeyPrefix+"token_transfer_count_by_address_"+address.Address,
				address.TokenTransferCount,
			)
			// Log count
			err = redis.GetRedisClient().SetCount(
				config.Config.RedisKeyPrefix+"log_count_by_address_"+address.Address,
				address.LogCount,
			)
			if err != nil {
				// Redis error
				zap.S().Warn(err)
				return err
			}
		}
		skip += limit
	}
	return nil
}

//func byTokenContractCountRoutine() {
//	countRoutineCron(byAddressCountExec)
//}

//// Count by address from PG and set in redis
//func byTokenContractCountExec() error {
//	token_contracts, err := crud.GetAddressCrud().SelectAllTokenContracts()
//	if err != nil {
//		// Redis error
//		zap.S().Warn(err)
//		return err
//	}
//	println(token_contracts)
//	for _, token_contract := range *token_contracts {
//		count, err := crud.GetTokenTransferCrud().CountByTokenContract(token_contract.Address)
//		if err != nil {
//			// Redis error
//			zap.S().Warn(err)
//			return err
//		}
//
//		err = redis.GetRedisClient().SetCount(
//			config.Config.RedisKeyPrefix+"token_transfer_count_by_token_contract_"+token_contract.Address,
//			count,
//		)
//		if err != nil {
//			// Redis error
//			zap.S().Warn(err)
//			return err
//		}
//	}
//	return nil
//}
