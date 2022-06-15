package _old

import (
	"time"

	"go.uber.org/zap"

	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/models"
	"github.com/sudoblockio/icon-transformer/redis"
)

func redisStoreKeysRoutine() {

	// Loop every duration
	for {

		// Get all keys
		keys, err := redis.GetRedisClient().GetAllKeys(config.Config.RedisKeyPrefix + "*")
		if err != nil {
			zap.S().Warn(
				"Routine=RedisStoreKeysRoutine",
				" Step=", "get redis keys",
				" Error=", err.Error(),
			)
			time.Sleep(1 * time.Second)
			continue
		}

		for _, key := range keys {
			value, err := redis.GetRedisClient().GetValue(key)
			if err != nil {
				zap.S().Warn(
					"Routine=RedisStoreKeysRoutine",
					" Step=", "get redis value",
					" Key=", key,
					" Error=", err.Error(),
				)
				continue
			}

			redisKey := &models.RedisKey{
				Key:   key,
				Value: value,
			}

			crud.GetRedisKeyCrud().LoaderChannel <- redisKey
		}

		zap.S().Info("Routine=RedisStoreKeysRoutine - Completed routine, sleeping ", config.Config.RoutinesSleepDuration.String(), "...")
		time.Sleep(config.Config.RoutinesSleepDuration)
	}
}
