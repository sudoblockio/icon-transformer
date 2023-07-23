package routines

import (
	"time"

	"go.uber.org/zap"

	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/redis"
)

func countAddressesToRedisRoutine() {
	// All
	countAll, err := crud.GetAddressCrud().Count()
	if err != nil {
		// Try again
		zap.S().Warn("Routine=AddressCount - ERROR: ", err.Error())
		time.Sleep(1 * time.Second)
		return
	}

	countContracts, err := crud.GetAddressCrud().CountWhere("is_contract", "true")
	if err != nil {
		// Try again
		zap.S().Warn("Routine=AddressCount - ERROR: ", err.Error())
		time.Sleep(1 * time.Second)
		return
	}

	countTokenAddresses, err := crud.GetTokenAddressCrud().CountByTokenContractAddress()
	if err != nil {
		// Try again
		zap.S().Warn("Routine=AddressCount - ERROR: ", err.Error())
		time.Sleep(1 * time.Second)
		return
	}

	err = redis.GetRedisClient().SetCount(config.Config.RedisKeyPrefix+"address_count", countAll)
	err = redis.GetRedisClient().SetCount(config.Config.RedisKeyPrefix+"address_contract_count", countContracts)

	for address, count := range countTokenAddresses {
		err = redis.GetRedisClient().SetCount(config.Config.RedisKeyPrefix+"token_address_count_by_token_contract_"+address, count)
		if err != nil {
			// Try again
			zap.S().Warn("Routine=AddressCount Redis - ERROR: ", err.Error())
		}
	}

	zap.S().Info("Routine=AddressCount - Completed routine...")
}
