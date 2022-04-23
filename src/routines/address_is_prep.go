package routines

import (
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/models"
	"github.com/sudoblockio/icon-transformer/service"
	"time"

	"github.com/sudoblockio/icon-transformer/config"
	"go.uber.org/zap"
)

func addressIsPrep() {

	// Loop every duration
	for {
		pRepAddressesDb, err := crud.GetAddressCrud().SelectPReps()
		pRepAddressesRpc, err := service.IconNodeServiceGetPreps()

		if err != nil {
			zap.S().Warn("Routine=AddressIsPrep - Error routine, sleeping ", config.Config.RoutinesSleepDuration.String(), err.Error())
			time.Sleep(config.Config.RoutinesSleepDuration)
		}

		// Add new preps
		for _, pRepAddressRpc := range pRepAddressesRpc {

			isInDb := false

			for _, pRepAddressDb := range *pRepAddressesDb {
				if pRepAddressRpc == pRepAddressDb.Address {
					isInDb = true
					break
				}
			}

			// Add to DB
			if isInDb == false {
				address := &models.Address{
					Address: pRepAddressRpc,
					IsPrep:  true,
				}
				crud.GetAddressCrud().LoaderChannel <- address
			}
		}

		// Remove old preps
		for _, pRepAddressDb := range *pRepAddressesDb {

			isInRpc := false

			for _, pRepAddressRpc := range pRepAddressesRpc {
				if pRepAddressRpc == pRepAddressDb.Address {
					isInRpc = true
					break
				}
			}

			// Update in DB
			if isInRpc == false {
				address := &models.Address{
					Address: pRepAddressDb.Address,
					IsPrep:  false,
				}
				crud.GetAddressCrud().LoaderChannel <- address
			}
		}

		zap.S().Info("Routine=AddressIsPrep - Completed routine, sleeping ", config.Config.RoutinesSleepDuration.String(), "...")
		time.Sleep(config.Config.RoutinesSleepDuration)
	}
}
