package crud

import (
	"sync"

	"github.com/sudoblockio/icon-transformer/models"
)

var addressTokenCrudOnce sync.Once
var TokenAddressCrud *Crud[models.TokenAddress, models.TokenAddressORM]

// GetTokenAddressCrud - create and/or return the addressToken table Model
func GetTokenAddressCrud() *Crud[models.TokenAddress, models.TokenAddressORM] {
	addressTokenCrudOnce.Do(func() {
		TokenAddressCrud = GetCrud(models.TokenAddress{}, models.TokenAddressORM{})

		TokenAddressCrud.Migrate()
		TokenAddressCrud.removeColumnNames([]string{"balance"})

		TokenAddressCrud.MakeStartLoaderChannel()
	})
	return TokenAddressCrud
}

func InitTokenAddressCrud() {
	GetTokenAddressCrud()
}

var addressTokenBalanceCrudOnce sync.Once
var TokenAddressBalanceCrud *Crud[models.TokenAddress, models.TokenAddressORM]

// GetTokenAddressCrud - create and/or return the addressToken table Model
func GetTokenAddressBalanceCrud() *Crud[models.TokenAddress, models.TokenAddressORM] {
	addressTokenBalanceCrudOnce.Do(func() {
		TokenAddressBalanceCrud = GetCrud(models.TokenAddress{}, models.TokenAddressORM{})
		TokenAddressBalanceCrud.columns = []string{"address", "token_contract_address", "balance"}
		TokenAddressBalanceCrud.metrics.Name = "address_balance_routine"
		TokenAddressBalanceCrud.MakeStartLoaderChannel()
	})

	return TokenAddressBalanceCrud
}
