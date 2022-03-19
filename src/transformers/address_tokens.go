package transformers

import (
	"github.com/sudoblockio/icon-go-worker/models"
)

func transformBlockETLToAddressTokens(blockETL *models.BlockETL) []*models.AddressToken {

	addressTokens := []*models.AddressToken{}

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

					addressToken := &models.AddressToken{
						Address:              fromAddress,
						TokenContractAddress: logETL.Address,
					}
					addressTokens = append(addressTokens, addressToken)
				}

				// To Address
				toAddress := logETL.Indexed[2]
				if toAddress != "" {

					addressToken := &models.AddressToken{
						Address:              toAddress,
						TokenContractAddress: logETL.Address,
					}
					addressTokens = append(addressTokens, addressToken)
				}
			}
		}
	}

	return addressTokens
}
