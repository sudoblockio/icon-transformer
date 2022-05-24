package routines

import (
	"errors"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/service"
	"github.com/sudoblockio/icon-transformer/utils"
)

func addressBalanceRoutine() {
	for i := 0; i <= config.Config.RoutinesNumWorkers; i++ {
		go getAddressBalances(i)
	}
}

func getAddressBalances(worker_id int) {

	// Loop every duration
	for {

		// Loop through all addresses
		//skip := 0
		skip := worker_id * config.Config.RoutinesBatchSize
		limit := config.Config.RoutinesBatchSize
		for {
			addresses, err := crud.GetAddressCrud().SelectMany(limit, skip)
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// Sleep
				break
			} else if err != nil {
				zap.S().Fatal(err.Error())
			}
			if len(*addresses) == 0 {
				// Sleep
				break
			}

			zap.S().Info("Routine=AddressBalance", " - Processing ", len(*addresses), " addresses...")
			for _, address := range *addresses {

				/////////////
				// Balance //
				/////////////

				// Node call
				balance, err := service.IconNodeServiceGetBalance(address.Address)
				if err != nil {
					// Icon node error
					zap.S().Warn("Routine=Balance, Address=", address.Address, " - Error: ", err.Error())
					continue
				}

				// Hex -> float64
				address.Balance = utils.StringHexToFloat64(balance, 18)

				////////////////////
				// Staked Balance //
				////////////////////
				stakedBalance, err := service.IconNodeServiceGetStakedBalance(address.Address)
				if err != nil {
					// Icon node error
					zap.S().Warn("Routine=Balance, Address=", address.Address, " - Error: ", err.Error())
					continue
				}

				// Hex -> float64
				address.Balance += utils.StringHexToFloat64(stakedBalance, 18)

				// Copy struct for pointer conflicts
				//addressCopy := &models.Address{}
				//copier.Copy(addressCopy, &address)
				// Insert to database
				//crud.GetAddressCrud().LoaderChannel <- addressCopy

				crud.GetAddressCrud().UpsertOneCols(&address, []string{"address", "balance"})
			}

			skip += skip + config.Config.RoutinesBatchSize*config.Config.RoutinesNumWorkers
		}
		zap.S().Info("Routine=AddressBalance - Completed routine, sleeping ", config.Config.RoutinesSleepDuration.String(), "...")
		time.Sleep(config.Config.RoutinesSleepDuration)
	}
}
