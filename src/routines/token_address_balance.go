package routines

import (
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/models"
	"github.com/sudoblockio/icon-transformer/service"
	"github.com/sudoblockio/icon-transformer/utils"
	"go.uber.org/zap"
)

func setTokenAddressBalances(tokenAddress *models.TokenAddress) {
	// Token Balance //
	balance, err := service.IconNodeServiceGetTokenBalance(tokenAddress.TokenContractAddress, tokenAddress.Address)
	if err != nil {
		// Icon node error
		zap.S().Warn("Routine=TokenBalance, Address=", tokenAddress.Address, " - Error: ", err.Error())
		return
	}

	// TODO: This may all be wrong as we might need to be storing decimals somewhere
	// If this is the case, then we'll need to rethink all of this as it might be then worth it to make another
	// table so just store address details for contracts
	// There are ~1000x more addresses than contracts so putting more info in addresses table might be not worth it
	// Also begs the question if splitting up the contracts service was a good idea
	// Hex -> float64
	tokenAddress.Balance = utils.StringHexToFloat64(balance, 18)

	crud.GetTokenAddressCrud().UpsertOneCols(tokenAddress, []string{"address", "balance", "token_address"})
}
