package transformers

import (
	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/models"
	"github.com/sudoblockio/icon-transformer/service"
	"github.com/sudoblockio/icon-transformer/utils"
	"go.uber.org/zap"
	"time"
)

var allTokenAddresses = make(map[string]bool)

func LoadTokenContractCheckDuplicates(tokenAddress *models.TokenAddress) {
	if _, ok := allTokenAddresses[tokenAddress.Address+tokenAddress.TokenContractAddress]; !ok {

		allTokenAddresses[tokenAddress.Address+tokenAddress.TokenContractAddress] = true
		crud.TokenAddressCrud.LoaderChannel <- tokenAddress

		// Metrics
		metricsBlockTransformer.tokenAddressesSeen.Inc()
		return
	}
	metricsBlockTransformer.tokenAddressesIgnored.Inc()
}

func tokenAddresses(blockETL *models.BlockETL) {

	tokenAddresses := []*models.TokenAddress{}
	tokenAddress := &models.TokenAddress{}

	for _, transactionETL := range blockETL.Transactions {
		for _, logETL := range transactionETL.Logs {

			if logETL.Indexed[0] == "Transfer(Address,Address,int,bytes)" && len(logETL.Indexed) == 4 {
				// Token Transfer

				// From Address
				fromAddress := logETL.Indexed[1]
				if fromAddress != "" {

					tokenAddress := &models.TokenAddress{
						Address:              fromAddress,
						TokenContractAddress: logETL.Address,
					}
					if config.Config.ProcessCounts {
						tokenAddresses = append(tokenAddresses, tokenAddress)
					}
					LoadTokenContractCheckDuplicates(tokenAddress)
				}

				// To Address
				toAddress := logETL.Indexed[2]
				if toAddress != "" && tokenAddress.Address != toAddress {

					tokenAddress := &models.TokenAddress{
						Address:              toAddress,
						TokenContractAddress: logETL.Address,
					}
					if config.Config.ProcessCounts {
						tokenAddresses = append(tokenAddresses, tokenAddress)
					}
					LoadTokenContractCheckDuplicates(tokenAddress)
				}
			}
		}
	}
	if config.Config.ProcessCounts {
		tokenAddressBalances(blockETL, tokenAddresses)
	}
}

func tokenAddressBalances(blockETL *models.BlockETL, tokenAddresses []*models.TokenAddress) {

	blockTimestamp := time.Unix(blockETL.Timestamp/1000000, 0)

	if time.Since(blockTimestamp) <= config.Config.TransformerServiceCallThreshold {
		for _, tokenAddress := range tokenAddresses {
			balance, err := service.IconNodeServiceGetTokenBalance(tokenAddress.TokenContractAddress, tokenAddress.Address)
			if err != nil {
				// Icon node error
				zap.S().Warn("Routine=TokenAddressBalanceRoutine, Address=", tokenAddress.Address, " - Error: ", err.Error())
				continue
			}

			decimalBase, err := service.IconNodeServiceGetTokenDecimalBase(tokenAddress.TokenContractAddress)
			if err != nil {
				// Icon node error
				zap.S().Warn("Routine=TokenAddressBalanceRoutine - Error: ", err.Error())
				continue
			}
			tokenAddress.Balance = utils.StringHexToFloat64(balance, decimalBase)

			// Insert to database
			crud.GetTokenAddressCrud().UpsertOneColumns(tokenAddress, []string{"address", "balance", "token_contract_address"})
		}
	}
}
