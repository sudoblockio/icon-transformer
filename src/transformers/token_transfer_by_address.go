package transformers

import (
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/models"
)

func transformBlockETLToTokenTransferByAddresses(blockETL *models.BlockETL) {

	//tokenTransferByAddresses := []*models.TokenTransferByAddress{}
	loaderChannel := crud.GetTokenTransferByAddressCrud().LoaderChannel

	//////////
	// Logs //
	//////////
	for _, transactionETL := range blockETL.Transactions {
		for iL, logETL := range transactionETL.Logs {

			// NOTE check for specific log definition
			// NOTE 'Transfer' is not a protected name in Icon
			if logETL.Indexed[0] == "Transfer(Address,Address,int,bytes)" && len(logETL.Indexed) == 4 ||
				logETL.Indexed[0] == "Transfer(Address,Address,int)" && len(logETL.Indexed) == 4 ||
				logETL.Indexed[0] == "TransferSingle(Address,Address,Address,int,int)" && len(logETL.Indexed) == 4 {
				// Token Transfers
				// From Address
				fromAddress := logETL.Indexed[1]
				tokenTransferByFromAddress := &models.TokenTransferByAddress{
					TransactionHash: transactionETL.Hash,
					LogIndex:        int64(iL),
					Address:         fromAddress,
					BlockNumber:     blockETL.Number,
				}
				loaderChannel <- tokenTransferByFromAddress

				// To Address
				toAddress := logETL.Indexed[2]
				tokenTransferByToAddress := &models.TokenTransferByAddress{
					TransactionHash: transactionETL.Hash,
					LogIndex:        int64(iL),
					Address:         toAddress,
					BlockNumber:     blockETL.Number,
				}
				loaderChannel <- tokenTransferByToAddress
			}
		}
	}
}
