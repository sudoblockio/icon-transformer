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
	//if config.Config.RedisRecoveryAddresses {
	//	LoopRoutine(crud.GetCrud(models.Address{}, models.AddressORM{}), addressRoutines)
	//}
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
	skip := 0

	// Loop until there are no more records
	for {
		zap.S().Infof("Fetching batch with skip=%d, limit=%d, table=%s", skip, batchSize, Crud.TableName)
		routineItems, err := Crud.SelectBatchOrder(batchSize, skip)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			zap.S().Warnf("No more records (%s)", err.Error())
			break
		}
		if err != nil {
			zap.S().Warnf("Error fetching records (%s)", err.Error())
			break
		}
		if len(*routineItems) == 0 {
			zap.S().Warn("No more records")
			break
		}

		// Process the fetched records asynchronously.
		// We use a WaitGroup to ensure all processing for the current batch finishes
		// before fetching the next batch.
		var procWg sync.WaitGroup
		for i := range *routineItems {
			item := &(*routineItems)[i]
			// Run each routine concurrently for the current record.
			for _, routine := range routines {
				procWg.Add(1)
				go func(proc func(*M), record *M) {
					defer procWg.Done()
					proc(record)
				}(routine, item)
			}
		}
		// Wait for all asynchronous processing on the batch to complete.
		procWg.Wait()

		// Move to the next batch.
		skip += batchSize
	}

	zap.S().Info("Finished processing table=", Crud.TableName)
}
