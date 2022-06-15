package routines

import (
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/models"
	"github.com/sudoblockio/icon-transformer/service"
	"github.com/sudoblockio/icon-transformer/utils"
	"go.uber.org/zap"
)

func setAddressBalances(address *models.Address) {

	// Balance //
	balance, err := service.IconNodeServiceGetBalance(address.Address)
	if err != nil {
		// Icon node error
		zap.S().Warn("Routine=Balance, Address=", address.Address, " - Error: ", err.Error())
		return
	}
	// Hex -> float64
	address.Balance = utils.StringHexToFloat64(balance, 18)

	// Staked Balance //
	stakedBalance, err := service.IconNodeServiceGetStakedBalance(address.Address)
	if err != nil {
		// Icon node error
		zap.S().Warn("Routine=Balance, Address=", address.Address, " - Error: ", err.Error())
		return
	}

	// Hex -> float64
	address.Balance += utils.StringHexToFloat64(stakedBalance, 18)

	crud.GetAddressCrud().UpsertOneCols(address, []string{"address", "balance"})
}
