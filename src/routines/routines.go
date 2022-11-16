package routines

import (
	"errors"
	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
	"sync"
	"time"
)

// Functions that are used in recovery that output a model that needs to be upserted into the address table
var addressRoutines = []func(a *models.Address){
	setAddressBalances,
	SetAddressTxCounts,
	setAddressContractMeta,
}

var tokenAddressRoutines = []func(t *models.TokenAddress){
	setTokenAddressBalances,
	// TODO: Fix this
	//setTokenAddressTxCounts,
}

func StartRecovery() {
	zap.S().Warn("Init recovery...")
	// Global count
	setTransactionCounts()

	// One shot
	if config.Config.NetworkName == "mainnet" {
		addressTypeRoutine()
	}
	countAddressesToRedisRoutine()

	// By address
	LoopRoutine(crud.GetCrud(models.Address{}, models.AddressORM{}), addressRoutines)
	// By token address
	LoopRoutine(crud.GetCrud(models.TokenAddress{}, models.TokenAddressORM{}), tokenAddressRoutines)

	zap.S().Info("finished recovery. exiting..")
	os.Exit(0)
}

var cronRoutines = []func(){
	addressIsPrep,
	tokenAddressCountRoutine, // Isn't used - RM?
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

func LoopRoutine[M any, O any](Crud *crud.Crud[M, O], routines []func(*M)) {
	var wg sync.WaitGroup
	for i := 0; i <= (config.Config.RoutinesNumWorkers - 1); i++ {
		wg.Add(1)

		i := i
		go func() {
			defer wg.Done()
			// Loop through all addresses
			skip := i * config.Config.RoutinesBatchSize
			limit := config.Config.RoutinesBatchSize

			zap.S().Info("Starting worker on table=", Crud.TableName, " with skip=", skip, " with workerId=", i)
			// Run loop until addresses have all been iterated over
			for {
				wg.Add(1)

				routineItems, err := Crud.SelectMany(limit, skip)
				if errors.Is(err, gorm.ErrRecordNotFound) {
					zap.S().Warn("Ending address routing with error=", err.Error())
					break
				} else if err != nil {
					zap.S().Warn("Ending address routing with error=", err.Error())
					break
				}
				if len(*routineItems) == 0 {
					zap.S().Warn("Ending address routing, no more addresses")
					break
				}

				zap.S().Info("Starting skip=", skip, " limit=", limit, " table=", Crud.TableName, " workerId=", i)
				for i := 0; i < len(*routineItems); i++ {
					var item *M
					item = &(*routineItems)[i]
					for _, r := range routines {
						r(item)
					}
				}

				zap.S().Info("Finished skip=", skip, " limit=", limit, " table=", Crud.TableName, " workerId=", i)

				skip += config.Config.RoutinesBatchSize * config.Config.RoutinesNumWorkers
				wg.Done()
			}
		}()
	}

	wg.Wait()
}
