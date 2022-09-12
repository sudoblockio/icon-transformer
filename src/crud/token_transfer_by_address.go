package crud

import (
	"sync"

	"github.com/sudoblockio/icon-transformer/models"
)

var TokenTransferByAddressCrud *Crud[models.TokenTransferByAddress, models.TokenTransferByAddressORM]
var tokenTransferByAddressCrudOnce sync.Once

// GetTokenTransferByAddressCrud - create and/or return the tokenTransferByAddresss table model
func GetTokenTransferByAddressCrud() *Crud[models.TokenTransferByAddress, models.TokenTransferByAddressORM] {
	tokenTransferByAddressCrudOnce.Do(func() {
		TokenTransferByAddressCrud = GetCrud(models.TokenTransferByAddress{}, models.TokenTransferByAddressORM{})

		TokenTransferByAddressCrud.Migrate()

		TokenTransferByAddressCrud.MakeStartLoaderChannel()
	})

	return TokenTransferByAddressCrud
}

func InitTokenTransferByAddressCrud() {
	GetTokenTransferByAddressCrud()
}
