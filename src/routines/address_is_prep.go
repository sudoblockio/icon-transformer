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

	//pRepAddressesDb, err := crud.GetAddressCrud().SelectWhere("is_prep", "true")

	pRepAddressesDb, err := crud.GetAddressCrud().SelectPrep()

	if err != nil {
		zap.S().Fatal(err.Error())
	}
	pRepAddressesRpc, err := service.IconNodeServiceGetPreps()
	if err != nil {
		zap.S().Fatal(err.Error())
	}

	if err != nil {
		zap.S().Warn("Routine=AddressIsPrep - Error routine, sleeping ", config.Config.RoutinesSleepDuration.String(), err.Error())
		time.Sleep(config.Config.RoutinesSleepDuration)
	}

	// Add new preps
	for _, pRepAddressRpc := range pRepAddressesRpc {

		isInDb := false

		for _, pRepAddressDb := range pRepAddressesDb {
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

			//err = crud.GetAddressCrud().UpsertOneColumns(address, []string{"address", "is_prep"})
			crud.GetAddressRoutineCruds()["address_is_prep"].LoaderChannel <- address
			if err != nil {
				zap.S().Fatal(err.Error())
			}
		}
	}

	// Remove old preps
	for _, pRepAddressDb := range pRepAddressesDb {

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
			//crud.GetAddressCrud().LoaderChannel <- address
			//err = crud.GetAddressCrud().UpsertOneColumns(address, []string{"address", "is_prep"})
			crud.GetAddressRoutineCruds()["address_is_prep"].LoaderChannel <- address

			if err != nil {
				zap.S().Fatal(err.Error())
			}
		}
	}
	zap.S().Info("Routine=AddressIsPrep - Completed routine...")
}
