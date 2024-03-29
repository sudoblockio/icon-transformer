package crud

import (
	"github.com/sudoblockio/icon-transformer/models"
	"sync"
)

var transactionInternalByAddressCrudOnce sync.Once
var TransactionInternalByAddressCrud *Crud[models.TransactionInternalByAddress, models.TransactionInternalByAddressORM]

// GetTransactionInternalByAddressCrud - create and/or return the transactionInternalByAddresss table Model
func GetTransactionInternalByAddressCrud() *Crud[models.TransactionInternalByAddress, models.TransactionInternalByAddressORM] {
	transactionInternalByAddressCrudOnce.Do(func() {
		TransactionInternalByAddressCrud = GetCrud(models.TransactionInternalByAddress{}, models.TransactionInternalByAddressORM{})

		TransactionInternalByAddressCrud.Migrate()
		TransactionInternalByAddressCrud.CreateIndexes(`
		create index if not exists transaction_by_address_idx_address_block_number
		on transaction_internal_by_addresses
		(address, block_number DESC);
		`)

		TransactionInternalByAddressCrud.MakeStartLoaderChannel()
	})

	return TransactionInternalByAddressCrud
}

func InitTransactionInternalByAddressCrud() {
	GetTransactionInternalByAddressCrud()
}
