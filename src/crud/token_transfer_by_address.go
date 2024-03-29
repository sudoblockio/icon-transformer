package crud

import (
	"sync"

	"github.com/sudoblockio/icon-transformer/models"
)

var TokenTransferByAddressCrud *Crud[models.TokenTransferByAddress, models.TokenTransferByAddressORM]
var tokenTransferByAddressCrudOnce sync.Once

// GetTokenTransferByAddressCrud - create and/or return the tokenTransferByAddresss table Model
func GetTokenTransferByAddressCrud() *Crud[models.TokenTransferByAddress, models.TokenTransferByAddressORM] {
	tokenTransferByAddressCrudOnce.Do(func() {
		TokenTransferByAddressCrud = GetCrud(models.TokenTransferByAddress{}, models.TokenTransferByAddressORM{})

		TokenTransferByAddressCrud.Migrate()
		TokenTransferByAddressCrud.CreateIndexes(`
		create index if not exists token_transfer_by_addresses_idx_address_block_number_index
		on public.token_transfer_by_addresses
		(address asc, block_number desc)
		`)

		TokenTransferByAddressCrud.MakeStartLoaderChannel()
	})

	return TokenTransferByAddressCrud
}

func InitTokenTransferByAddressCrud() {
	GetTokenTransferByAddressCrud()
}
