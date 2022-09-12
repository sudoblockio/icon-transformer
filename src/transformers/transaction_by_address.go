package transformers

import (
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/models"
)

func transactionByAddresses(blockETL *models.BlockETL) {

	transactionByFromAddress := &models.TransactionByAddress{}

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

			crud.TransactionByAddressCrud.LoaderChannel <- transactionByFromAddress
		}
		if transactionETL.ToAddress != "" && transactionETL.ToAddress != transactionByFromAddress.Address {
			transactionByToAddress := &models.TransactionByAddress{
				TransactionHash: transactionETL.Hash,
				Address:         transactionETL.ToAddress,
				BlockNumber:     blockETL.Number,
			}

			crud.TransactionByAddressCrud.LoaderChannel <- transactionByToAddress
		}
	}
}
