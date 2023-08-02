package routines

import (
	"errors"
	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
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

	//Global count
	setTransactionCounts()
	countAddressesToRedisRoutine()

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

	skip := 0
	limit := config.Config.RoutinesBatchSize

	zap.S().Info("Starting worker on table=", Crud.TableName, " with skip=", skip, " with limit=", limit)
	// Run loop until addresses have all been iterated over
	for {
		// Grabs a set of addresses or just contracts
		routineItems, err := Crud.SelectBatchOrder(limit, skip, "address", config.Config.RedisRecoveryContractsOnly)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			zap.S().Warn("Ending address routine with error=", err.Error())
			break
		}
		if err != nil {
			zap.S().Warn("Ending address routine with error=", err.Error())
			break
		}
		if len(*routineItems) == 0 {
			zap.S().Warn("Ending address routine, no more addresses")
			break
		}

		zap.S().Info("Starting skip=", skip, " limit=", limit, " table=", Crud.TableName)
		for i := 0; i < len(*routineItems); i++ {
			item := &(*routineItems)[i]
			for _, r := range routines {
				r(item)
			}
		}
		zap.S().Info("Finished skip=", skip, " limit=", limit, " table=", Crud.TableName)

		skip += config.Config.RoutinesBatchSize * config.Config.RoutinesNumWorkers
	}
}
