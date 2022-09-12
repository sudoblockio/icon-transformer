package routines

import (
	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/redis"
	"go.uber.org/zap"
)

func set[M any, O any](count int64, Crud *crud.Crud[M, O]) {
	err := redis.GetRedisClient().SetCount(
		config.Config.RedisKeyPrefix+"transaction_regular_count",
		count,
	)
	if err != nil {
		zap.S().Fatal(err.Error())
	}

	//Crud.LoaderChannel <- count

}

func setRedisKey(count int64, redisKey string) {
	err := redis.GetRedisClient().SetCount(
		config.Config.RedisKeyPrefix+redisKey,
		count,
	)
	if err != nil {
		zap.S().Fatal(err.Error())
	}
}

// Takes a crud count method, calls it, takes the count and puts it into a redis countKey.
func CrudCountSetRedis(c func() (int64, error), countKey string) error {
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
