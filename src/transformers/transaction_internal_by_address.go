package transformers

import (
	"github.com/sudoblockio/icon-go-worker/models"
)

func transformBlockETLToTransactionInternalByAddresses(blockETL *models.BlockETL) []*models.TransactionInternalByAddress {

	transactionInternalByAddresses := []*models.TransactionInternalByAddress{}

	//////////
	// Logs //
	//////////
	for _, transactionETL := range blockETL.Transactions {
		for iL, logETL := range transactionETL.Logs {

			method := extractMethodFromLogETL(logETL)

			// NOTE 'ICXTransfer' is a protected name in Icon
			if method == "ICXTransfer" {
				// Internal Transaction

				// From Address
				fromAddress := logETL.Indexed[1]
				transactionInternalByFromAddress := &models.TransactionInternalByAddress{
					TransactionHash: transactionETL.Hash,
					LogIndex:        int64(iL), // NOTE logs will always be in the same order from ETL
					Address:         fromAddress,
					BlockNumber:     blockETL.Number,
				}
				transactionInternalByAddresses = append(transactionInternalByAddresses, transactionInternalByFromAddress)

				// To Address
				toAddress := logETL.Indexed[2]
				transactionInternalByToAddress := &models.TransactionInternalByAddress{
					TransactionHash: transactionETL.Hash,
					LogIndex:        int64(iL), // NOTE logs will always be in the same order from ETL
					Address:         toAddress,
					BlockNumber:     blockETL.Number,
				}
				transactionInternalByAddresses = append(transactionInternalByAddresses, transactionInternalByToAddress)
			}
		}
	}

	return transactionInternalByAddresses
}
