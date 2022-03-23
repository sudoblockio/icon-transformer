package routines

import (
	"time"

	"go.uber.org/zap"

	"github.com/sudoblockio/icon-go-worker/config"
)

func transactionCreateScoreRoutine() {

	// TODO
	// Loop every duration
	for {

		zap.S().Info("Routine=TransactionCreateScoreRoutine - Completed routine, sleeping ", config.Config.RoutinesSleepDuration.String(), "...")
		time.Sleep(config.Config.RoutinesSleepDuration)
	}
}
