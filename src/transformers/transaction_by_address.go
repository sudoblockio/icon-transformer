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

		// This is for speeding up queries for single page views so that not many where conditions need to be used.
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
