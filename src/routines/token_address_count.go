package routines

import (
	"time"

	"go.uber.org/zap"

	"github.com/sudoblockio/icon-go-worker/config"
	"github.com/sudoblockio/icon-go-worker/crud"
	"github.com/sudoblockio/icon-go-worker/redis"
)

func tokenAddressCountRoutine() {

	// Loop every duration
	for {

		tokenAddressCounts, err := crud.GetTokenAddressCrud().CountByTokenContractAddress()
		if err != nil {
			zap.S().Fatal(
				"Routine=TokenAddressCount,",
				" Step=", "Get count from db",
				" Error=", err.Error(),
			)
		}

		zap.S().Info("Routine=TokenAddressCount", " - Processing tokenAddressCounts...")
		for tokenContractAddress, count := range tokenAddressCounts {
			countKey := config.Config.RedisKeyPrefix + "token_address_count_by_token_contract_" + tokenContractAddress
			err = redis.GetRedisClient().SetCount(countKey, count)
			if err != nil {
				zap.S().Warn(
					"Routine=TokenAddressCount,",
					" Step=", "Set count in redis",
					" Error=", err.Error(),
				)
			}
		}

		zap.S().Info("Routine=TokenAddressCount - Completed routine, sleeping ", config.Config.RoutinesSleepDuration.String(), "...")
		time.Sleep(config.Config.RoutinesSleepDuration)
	}
}
