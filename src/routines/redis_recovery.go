package routines

import (
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/redis"
	"go.uber.org/zap"
)

func StartRedisRecovery() {

	// Missing Keys
	go redisRecovery()

}

func redisRecovery() {

	redisKeys, err := crud.GetRedisKeyCrud().SelectAll()
	if err != nil {
		zap.S().Fatal(err.Error())
	}

	for _, redisKey := range *redisKeys {
		err = redis.GetRedisClient().SetValue(redisKey.Key, redisKey.Value)
		if err != nil {
			zap.S().Fatal(err.Error())
		}
	}

	zap.S().Info("Routine=redisRecovery - Recovered redis")
}
