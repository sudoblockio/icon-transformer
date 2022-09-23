package routines

import (
	"go.uber.org/zap"

	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/redis"
)

// TODO: RM? This sets a key that isn't used in API

// Sets a redis key for the number
func tokenAddressCountRoutine() {
	tokenAddressCounts, err := crud.GetTokenAddressCrud().CountByTokenContractAddress()
	if err != nil {
		zap.S().Fatal(err.Error())
	}

	zap.S().Info("Routine=TokenAddressCountRoutine", " - Processing tokenAddressCounts...")
	for tokenContractAddress, count := range tokenAddressCounts {
		countKey := config.Config.RedisKeyPrefix + "token_address_count_by_token_contract_" + tokenContractAddress
		err = redis.GetRedisClient().SetCount(countKey, count)
		if err != nil {
			zap.S().Warn(err.Error())
		}
	}

	zap.S().Info("Routine=TokenAddressCountRoutine - Completed routine...")
}
