package routines

import (
	"errors"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/jinzhu/copier"
	"github.com/sudoblockio/icon-go-worker/config"
	"github.com/sudoblockio/icon-go-worker/crud"
	"github.com/sudoblockio/icon-go-worker/models"
	"github.com/sudoblockio/icon-go-worker/service"
	"github.com/sudoblockio/icon-go-worker/utils"
)

func addressBalanceRoutine() {

	// Loop every duration
	for {

		// Loop through all addresses
		skip := 0
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
				addressCopy := &models.Address{}
				copier.Copy(addressCopy, &address)

				// Insert to database
				crud.GetAddressCrud().LoaderChannel <- addressCopy
			}

			skip += limit
		}
		zap.S().Info("Routine=AddressBalance - Completed routine, sleeping ", config.Config.RoutinesSleepDuration.String(), "...")
		time.Sleep(config.Config.RoutinesSleepDuration)
	}
}