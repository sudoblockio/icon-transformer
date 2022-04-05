package routines

import (
	"time"

	"go.uber.org/zap"

	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/redis"
)

func addressCountRoutine() {

	// Loop every duration
	for {

		///////////////
		// Get count //
		///////////////
		count, err := crud.GetAddressCrud().Count()
		if err != nil {
			// Try again
			zap.S().Warn("Routine=AddressCount - ERROR: ", err.Error())
			time.Sleep(1 * time.Second)
			continue
		}

		///////////////
		// Set count //
		///////////////
		redis.GetRedisClient().SetCount(config.Config.RedisKeyPrefix+"address_count", count)

		zap.S().Info("Routine=AddressCount - Completed routine, sleeping ", config.Config.RoutinesSleepDuration.String(), "...")
		time.Sleep(config.Config.RoutinesSleepDuration)
	}
}
