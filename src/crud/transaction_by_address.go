package crud

import (
	"github.com/sudoblockio/icon-transformer/models"
	"sync"
)

var transactionByAddressCrudOnce sync.Once
var TransactionByAddressCrud *Crud[models.TransactionByAddress, models.TransactionByAddressORM]
var TransactionByAddressCreateScoreCrud *Crud[models.TransactionByAddress, models.TransactionByAddressORM]

// InitTransactionByAddressCrud - create and/or return the transactionByAddresss table Model
func GetTransactionByAddressCrud() *Crud[models.TransactionByAddress, models.TransactionByAddressORM] {
	transactionByAddressCrudOnce.Do(func() {
		TransactionByAddressCrud = GetCrud(models.TransactionByAddress{}, models.TransactionByAddressORM{})

		TransactionByAddressCrud.Migrate()
		TransactionByAddressCrud.CreateIndexes(`
		create index if not exists transaction_by_addresses_idx_address_block_number_index 
		on public.transaction_by_addresses 
		(address asc, block_number desc)
		`)

		TransactionByAddressCrud.MakeStartLoaderChannel()

		TransactionByAddressCreateScoreCrud = GetCrud(models.TransactionByAddress{}, models.TransactionByAddressORM{})
		TransactionByAddressCreateScoreCrud.columns = []string{"transaction_hash", "address", "block_number"}
		TransactionByAddressCreateScoreCrud.metrics.Name = TransactionByAddressCrud.TableName + "types"
		TransactionByAddressCreateScoreCrud.MakeStartLoaderChannel()
	})

	return TransactionByAddressCrud
}

// InitTransactionByAddressCrud - initialize the transactionByAddresss table Model and loaders
func InitTransactionByAddressCrud() {
	GetTransactionByAddressCrud()
}
