package crud

import (
	"github.com/sudoblockio/icon-transformer/models"
	"sync"
)

var transactionByAddressCrudOnce sync.Once
var TransactionByAddressCrud *Crud[models.TransactionByAddress, models.TransactionByAddressORM]
var TransactionByAddressCreateScoreCrud *Crud[models.TransactionByAddress, models.TransactionByAddressORM]

// InitTransactionByAddressCrud - create and/or return the transactionByAddresss table model
func GetTransactionByAddressCrud() *Crud[models.TransactionByAddress, models.TransactionByAddressORM] {
	transactionByAddressCrudOnce.Do(func() {
		TransactionByAddressCrud = GetCrud(models.TransactionByAddress{}, models.TransactionByAddressORM{})

		TransactionByAddressCrud.Migrate()
		TransactionByAddressCrud.CreateIndexes("CREATE INDEX IF NOT EXISTS transaction_by_addresses_idx_address_block_number_index ON public.transaction_by_addresses (address asc, block_number desc)")

		TransactionByAddressCrud.MakeStartLoaderChannel()

		TransactionByAddressCreateScoreCrud = GetCrud(models.TransactionByAddress{}, models.TransactionByAddressORM{})
		TransactionByAddressCreateScoreCrud.columns = []string{"transaction_hash", "address", "block_number"}
		TransactionByAddressCreateScoreCrud.metrics.Name = TransactionByAddressCrud.TableName + "types"
		TransactionByAddressCreateScoreCrud.MakeStartLoaderChannel()
	})

	return TransactionByAddressCrud
}

// InitTransactionByAddressCrud - initialize the transactionByAddresss table model and loaders
func InitTransactionByAddressCrud() {
	GetTransactionByAddressCrud()
}
