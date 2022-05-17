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

func tokenAddressBalanceRoutine() {

	// Loop every duration
	for {

		// Loop through all addresses
		skip := 0
		limit := config.Config.RoutinesBatchSize
		for {
			tokenAddresses, err := crud.GetTokenAddressCrud().SelectMany(limit, skip)
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// Sleep
				break
			} else if err != nil {
				zap.S().Fatal(err.Error())
			}
			if len(*tokenAddresses) == 0 {
				// Sleep
				break
			}

			zap.S().Info("Routine=TokenAddressBalanceRoutine", " - Processing ", len(*tokenAddresses), " tokenAddresses...")
			for _, tokenAddress := range *tokenAddresses {

				/////////////
				// Balance //
				/////////////

				// Node call
				balance, err := service.IconNodeServiceGetTokenBalance(tokenAddress.TokenContractAddress, tokenAddress.Address)
				if err != nil {
					// Icon node error
					zap.S().Warn("Routine=TokenAddressBalanceRoutine, Address=", tokenAddress.Address, " - Error: ", err.Error())
					continue
				}

				// Hex -> float64
				decimalBase, err := service.IconNodeServiceGetTokenDecimalBase(tokenAddress.TokenContractAddress)
				if err != nil {
					// Icon node error
					zap.S().Warn("Routine=TokenAddressBalanceRoutine - Error: ", err.Error())
					continue
				}
				tokenAddress.Balance = utils.StringHexToFloat64(balance, decimalBase)

				// Copy struct for pointer conflicts
				//tokenAddressCopy := &models.TokenAddress{}
				//copier.Copy(tokenAddressCopy, &tokenAddress)

				// Insert to database
				//crud.GetTokenAddressCrud().LoaderChannel <- tokenAddressCopy
				crud.GetTokenAddressCrud().UpsertOneCols(&tokenAddress, []string{"address", "balance"})
			}

			skip += limit
		}

		zap.S().Info("Routine=TokenAddressBalanceRoutine - Completed routine, sleeping ", config.Config.RoutinesSleepDuration.String(), "...")
		time.Sleep(config.Config.RoutinesSleepDuration)
	}
}
