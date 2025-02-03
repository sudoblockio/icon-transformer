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
	// TODO: Fix this - Might not need - not used it seems? This is not cached and does not hit db
	//  function signiture would suggest another loop anyways.
	//setTokenAddressTxCounts,
}

func StartRecovery() {
	zap.S().Warn("Init recovery...")
	// One shot
	if config.Config.NetworkName == "mainnet" {
		addressTypeRoutine()
	}

	////Global count
	//setTransactionCounts()
	//countAddressesToRedisRoutine()

	// By address
	if config.Config.RedisRecoveryAddresses {
		LoopRoutine(crud.GetCrud(models.Address{}, models.AddressORM{}), addressRoutines)
	}
	// By token address
	if config.Config.RedisRecoveryTokenAddresses {
		LoopRoutine(crud.GetCrud(models.TokenAddress{}, models.TokenAddressORM{}), tokenAddressRoutines)
	}

	zap.S().Info("finished recovery. exiting..")
	time.Sleep(30 * time.Second)
	os.Exit(0)
}

func LoopRoutine[M any, O any](Crud *crud.Crud[M, O], routines []func(*M)) {
	batchSize := config.Config.RoutinesBatchSize
	numWorkers := config.Config.RoutinesNumWorkers

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for workerID := 0; workerID < numWorkers; workerID++ {
		go func(id int) {
			defer wg.Done()
			// Each worker starts at a different offset
			skip := id * batchSize
			for {
				zap.S().Infof("Worker %d: fetching batch with skip=%d, limit=%d, table=%s", id, skip, batchSize, Crud.TableName)
				routineItems, err := Crud.SelectBatchOrder(
					batchSize,
					skip,
				)
				if errors.Is(err, gorm.ErrRecordNotFound) {
					zap.S().Warnf("Worker %d: no more records (%s)", id, err.Error())
					break
				}
				if err != nil {
					zap.S().Warnf("Worker %d: error fetching records (%s)", id, err.Error())
					break
				}
				if len(*routineItems) == 0 {
					zap.S().Warnf("Worker %d: no more records", id)
					break
				}

				// Process each record in the batch
				for i := 0; i < len(*routineItems); i++ {
					item := &(*routineItems)[i]
					for _, r := range routines {
						r(item)
					}
				}

				// Increment skip by the total number of records processed across all workers per iteration.
				skip += numWorkers * batchSize
			}
		}(workerID)
	}

	wg.Wait()
	zap.S().Info("All workers finished processing table=", Crud.TableName)
}
