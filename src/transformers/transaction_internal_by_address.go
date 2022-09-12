package transformers

import (
	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/models"
)

func transactionInternalByAddresses(blockETL *models.BlockETL) {

	transactionInternalByAddresses := []*models.TransactionInternalByAddress{}

	for _, transactionETL := range blockETL.Transactions {
		for iL, logETL := range transactionETL.Logs {

			method := extractMethodFromLogETL(logETL)

			// NOTE 'ICXTransfer' is a protected name in Icon
			if method == "ICXTransfer" {
				// From Address
				fromAddress := logETL.Indexed[1]
				transactionInternalByFromAddress := &models.TransactionInternalByAddress{
					TransactionHash: transactionETL.Hash,
					LogIndex:        int64(iL), // NOTE logs will always be in the same order from ETL
					Address:         fromAddress,
					BlockNumber:     blockETL.Number,
				}

				// To Address
				toAddress := logETL.Indexed[2]
				transactionInternalByToAddress := &models.TransactionInternalByAddress{
					TransactionHash: transactionETL.Hash,
					LogIndex:        int64(iL), // NOTE logs will always be in the same order from ETL
					Address:         toAddress,
					BlockNumber:     blockETL.Number,
				}

				if config.Config.ProcessCounts {
					transactionInternalByAddresses = append(transactionInternalByAddresses, transactionInternalByFromAddress)
					transactionInternalByAddresses = append(transactionInternalByAddresses, transactionInternalByToAddress)
				}
				crud.TransactionInternalByAddressCrud.LoaderChannel <- transactionInternalByToAddress
				crud.TransactionInternalByAddressCrud.LoaderChannel <- transactionInternalByFromAddress
				broadcastToWebsocketRedisChannel(blockETL, transactionInternalByAddresses, config.Config.RedisTransactionsChannel)
			}
		}
	}
}
