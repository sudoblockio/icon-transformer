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
	balance, err := service.IconNodeServiceGetBalance(address.Address)
	if err != nil && err.Error() != "Invalid response" {
		// Icon node error
		// Can happen when an address has only failed Txs in history and hence will have 0 / nil balance
		zap.S().Debug("Routine=Balance, Address=", address.Address, " - Error: ", err.Error())
	}
	// Hex -> float64
	address.Balance = utils.StringHexToFloat64(balance, 18)

	// Staked Balance //
	stakedBalance, err := service.IconNodeServiceGetStakedBalance(address.Address)
	if err != nil {
		// Icon node error
		zap.S().Debug("Routine=Balance, Address=", address.Address, " - Error: ", err.Error())
		return address
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
	}
}
