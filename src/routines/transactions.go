package routines

import (
	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/redis"
	"go.uber.org/zap"
)

func setTransactionCounts() {
	count, err := crud.GetTransactionCrud().CountRegular()
	if err != nil {
		zap.S().Fatal(err.Error())
	}
	err = redis.GetRedisClient().SetCount(
		config.Config.RedisKeyPrefix+"transaction_regular_count",
		count,
	)
	if err != nil {
		zap.S().Fatal(err.Error())
	}

	count, err = crud.GetTransactionCrud().CountInternal()
	if err != nil {
		zap.S().Fatal(err.Error())
	}
	err = redis.GetRedisClient().SetCount(
		config.Config.RedisKeyPrefix+"transaction_internal_count",
		count,
	)
	if err != nil {
		zap.S().Fatal(err.Error())
	}
}
