package transformers

import (
	"github.com/sudoblockio/icon-transformer/models"
)

func transformBlockETLToAddresses(blockETL *models.BlockETL) []*models.Address {

	addresses := []*models.Address{}

	//////////////////
	// Transactions //
	//////////////////
	for _, transactionETL := range blockETL.Transactions {

		// From address
		if transactionETL.FromAddress != "" {

			// Is contract
			isContract := transactionETL.FromAddress[:2] == "cx"

			address := &models.Address{
				Address:    transactionETL.FromAddress,
				IsContract: isContract,
			}
			addresses = append(addresses, address)
		}

		// To address
		if transactionETL.ToAddress != "" {

			// Is contract
			isContract := transactionETL.ToAddress[:2] == "cx"

			address := &models.Address{
				Address:    transactionETL.ToAddress,
				IsContract: isContract,
			}
			addresses = append(addresses, address)
		}
	}

	//////////
	// Logs //
	//////////
	for _, transactionETL := range blockETL.Transactions {
		for _, logETL := range transactionETL.Logs {

			method := extractMethodFromLogETL(logETL)

			// NOTE 'ICXTransfer' is a protected name in Icon
			if method == "ICXTransfer" {
				// Internal Transaction

				// From Address
				fromAddress := logETL.Indexed[1]
				if fromAddress != "" {

					// Is contract
					isContract := fromAddress[:2] == "cx"

					address := &models.Address{
						Address:    fromAddress,
						IsContract: isContract,
					}
					addresses = append(addresses, address)
				}

				// To Address
				toAddress := logETL.Indexed[2]
				if toAddress != "" {

					// Is contract
					isContract := toAddress[:2] == "cx"

					address := &models.Address{
						Address:    toAddress,
						IsContract: isContract,
					}
					addresses = append(addresses, address)
				}
			} else if logETL.Indexed[0] == "Transfer(Address,Address,int,bytes)" && len(logETL.Indexed) == 4 {
				// Token Transfer

				// From Address
				fromAddress := logETL.Indexed[1]
				if fromAddress != "" {

					// Is contract
					isContract := fromAddress[:2] == "cx"

					address := &models.Address{
						Address:    fromAddress,
						IsContract: isContract,
					}
					addresses = append(addresses, address)
				}

				// To Address
				toAddress := logETL.Indexed[2]
				if toAddress != "" {

					// Is contract
					isContract := toAddress[:2] == "cx"

					address := &models.Address{
						Address:    toAddress,
						IsContract: isContract,
					}
					addresses = append(addresses, address)
				}
			}
		}
	}

	return addresses
}

func transformContractToAddress(contract *models.ContractProcessed) *models.Address {

	address := &models.Address{
		Address:              contract.Address,
		Name:                 contract.Name,
		CreatedTimestamp:     contract.CreatedTimestamp,
		Status:               contract.Status,
		IsToken:              contract.IsToken,
		IsContract:           true,
		ContractUpdatedBlock: contract.ContractUpdatedBlock,
	}

	return address
}
