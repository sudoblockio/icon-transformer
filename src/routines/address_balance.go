package routines

import (
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/models"
	"github.com/sudoblockio/icon-transformer/service"
	"github.com/sudoblockio/icon-transformer/utils"
	"go.uber.org/zap"
)

func getAddressBalances(address *models.Address) *models.Address {

	// Balance //
	// Assume error means zero balance
	//balance, _ := service.IconNodeServiceGetBalance(address.Address)
	balance, err := service.IconNodeServiceGetBalance(address.Address)
	if err != nil && err.Error() != "Invalid response" {
		// Icon node error
		// Can happen when a node has only failed Txs in history
		zap.S().Warn("Routine=Balance, Address=", address.Address, " - Error: ", err.Error())
		//return nil
	}
	// Hex -> float64
	address.Balance = utils.StringHexToFloat64(balance, 18)

	// Staked Balance //
	stakedBalance, err := service.IconNodeServiceGetStakedBalance(address.Address)
	if err != nil {
		// Icon node error
		zap.S().Warn("Routine=Balance, Address=", address.Address, " - Error: ", err.Error())
		return nil
	}

	// Hex -> float64
	address.Balance += utils.StringHexToFloat64(stakedBalance, 18)
	return address
}

func setAddressBalances(address *models.Address) {
	//zap.S().Info(address.Address)
	addressNew := getAddressBalances(address)
	if address != nil {
		crud.GetAddressRoutineCruds()["address_balance"].LoaderChannel <- addressNew
	} else {

		println()
	}
}
