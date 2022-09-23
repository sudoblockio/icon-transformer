package routines

import (
	"github.com/sudoblockio/icon-transformer/crud"
	"go.uber.org/zap"
)

func setTransactionCounts() {
	// Regular Txs
	count, err := crud.GetTransactionCrud().CountTransactionsRegular()
	if err != nil {
		zap.S().Fatal(err.Error())
	}
	setRedisKey(count, "transaction_regular_count")

	// Internal Txs
	count, err = crud.GetTransactionCrud().CountWhere("type", "log")
	if err != nil {
		zap.S().Fatal(err.Error())
	}
	setRedisKey(count, "transaction_internal_count")

	// Token transfer Txs
	count, err = crud.GetTokenTransferCrud().Count()
	if err != nil {
		zap.S().Fatal(err.Error())
	}
	setRedisKey(count, "token_transfer_count")
}
