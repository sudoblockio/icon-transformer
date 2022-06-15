package _old

import (
	"time"

	"go.uber.org/zap"

	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/redis"
)

// TODO: rm

func addressCountRoutine() {
	for {
		// All
		countAll, err := crud.GetAddressCrud().CountAll()
		if err != nil {
			// Try again
			zap.S().Warn("Routine=AddressCount - ERROR: ", err.Error())
			time.Sleep(1 * time.Second)
			continue
		}

		countContracts, err := crud.GetAddressCrud().CountContracts()
		if err != nil {
			// Try again
			zap.S().Warn("Routine=AddressCount - ERROR: ", err.Error())
			time.Sleep(1 * time.Second)
			continue
		}

		countTokenAddresses, err := crud.GetTokenAddressCrud().CountByTokenContractAddress()
		if err != nil {
			// Try again
			zap.S().Warn("Routine=AddressCount - ERROR: ", err.Error())
			time.Sleep(1 * time.Second)
			continue
		}

		///////////////
		// Set count //
		///////////////
		redis.GetRedisClient().SetCount(config.Config.RedisKeyPrefix+"address_count", countAll)
		redis.GetRedisClient().SetCount(config.Config.RedisKeyPrefix+"address_contract_count", countContracts)

		for address, count := range countTokenAddresses {
			redis.GetRedisClient().SetCount(config.Config.RedisKeyPrefix+"token_address_count_by_token_contract_"+address, count)
		}

		zap.S().Info("Routine=AddressCount - Completed routine, sleeping ", config.Config.RoutinesSleepDuration.String(), "...")
		time.Sleep(config.Config.RoutinesSleepDuration)
	}
}
