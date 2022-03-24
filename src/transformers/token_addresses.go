package transformers

import (
	"github.com/sudoblockio/icon-transformer/models"
)

func transformBlockETLToTokenAddresses(blockETL *models.BlockETL) []*models.TokenAddress {

	tokenAddresses := []*models.TokenAddress{}

	//////////
	// Logs //
	//////////
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
					tokenAddresses = append(tokenAddresses, tokenAddress)
				}

				// To Address
				toAddress := logETL.Indexed[2]
				if toAddress != "" {

					tokenAddress := &models.TokenAddress{
						Address:              toAddress,
						TokenContractAddress: logETL.Address,
					}
					tokenAddresses = append(tokenAddresses, tokenAddress)
				}
			}
		}
	}

	return tokenAddresses
}
