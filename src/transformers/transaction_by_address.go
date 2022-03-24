package transformers

import (
	"github.com/sudoblockio/icon-transformer/models"
)

func transformBlockETLToTransactionByAddresses(blockETL *models.BlockETL) []*models.TransactionByAddress {

	transactionByAddresses := []*models.TransactionByAddress{}

	//////////////////
	// Transactions //
	//////////////////
	for _, transactionETL := range blockETL.Transactions {

		// From address
		if transactionETL.FromAddress != "" {
			transactionByFromAddress := &models.TransactionByAddress{
				TransactionHash: transactionETL.Hash,
				Address:         transactionETL.FromAddress,
				BlockNumber:     blockETL.Number,
			}

			transactionByAddresses = append(transactionByAddresses, transactionByFromAddress)
		}

		// To address
		if transactionETL.ToAddress != "" {
			transactionByToAddress := &models.TransactionByAddress{
				TransactionHash: transactionETL.Hash,
				Address:         transactionETL.ToAddress,
				BlockNumber:     blockETL.Number,
			}

			transactionByAddresses = append(transactionByAddresses, transactionByToAddress)
		}
	}

	return transactionByAddresses
}
