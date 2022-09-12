package crud

import (
	"github.com/sudoblockio/icon-transformer/models"
	"sync"
)

var transactionInternalByAddressCrudOnce sync.Once
var TransactionInternalByAddressCrud *Crud[models.TransactionInternalByAddress, models.TransactionInternalByAddressORM]

// GetTransactionInternalByAddressCrud - create and/or return the transactionInternalByAddresss table model
func GetTransactionInternalByAddressCrud() *Crud[models.TransactionInternalByAddress, models.TransactionInternalByAddressORM] {
	transactionInternalByAddressCrudOnce.Do(func() {
		TransactionInternalByAddressCrud = GetCrud(models.TransactionInternalByAddress{}, models.TransactionInternalByAddressORM{})

		TransactionInternalByAddressCrud.Migrate()

		TransactionInternalByAddressCrud.MakeStartLoaderChannel()
	})

	return TransactionInternalByAddressCrud
}

func InitTransactionInternalByAddressCrud() {
	GetTransactionInternalByAddressCrud()
}
