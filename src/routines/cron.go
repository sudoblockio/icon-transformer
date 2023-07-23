package routines

import (
	"github.com/sudoblockio/icon-transformer/config"
	"go.uber.org/zap"
	"time"
)

var cronRoutines = []func(){
	addressIsPrep,
	tokenAddressCountRoutine, // Isn't used - RM?
	countAddressesToRedisRoutine,
}

func CronStart() {

	zap.S().Warn("Init cron...")
	// Init - Jobs that run once on startup
	if config.Config.NetworkName == "mainnet" {
		addressTypeRoutine()
	}

	// Short
	go RoutinesCron(cronRoutines, config.Config.RoutinesSleepDuration)

	// Long
	// TODO: These were snubbed because they stress DB and should be run with something to slow them down
	//go AddressRoutinesCron(addressRoutines, 6*time.Hour)
}

// Wrapper for generic routines
func RoutinesCron(routines []func(), sleepDuration time.Duration) {
	for {
		zap.S().Warn("Starting cron...")
		for _, r := range routines {
			r()
		}
		zap.S().Info("Completed routine, sleeping...")
		time.Sleep(sleepDuration)
	}
}
