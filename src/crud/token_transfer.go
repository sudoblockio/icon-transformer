package crud

import (
	"sync"

	"github.com/sudoblockio/icon-transformer/models"
)

var tokenTransferCrudOnce sync.Once
var TokenTransferCrud *Crud[models.TokenTransfer, models.TokenTransferORM]

// GetTokenTransferCrud - create and/or return the tokenTransfers table Model
func GetTokenTransferCrud() *Crud[models.TokenTransfer, models.TokenTransferORM] {
	tokenTransferCrudOnce.Do(func() {
		TokenTransferCrud = GetCrud(models.TokenTransfer{}, models.TokenTransferORM{})

		TokenTransferCrud.Migrate()
		TokenTransferCrud.CreateIndexes("CREATE INDEX IF NOT EXISTS token_transfer_idx_token_contract_address_block_number ON public.token_transfers USING btree (token_contract_address, block_number)")

		TokenTransferCrud.MakeStartLoaderChannel()
	})

	return TokenTransferCrud
}

func InitTokenTransferCrud() {
	GetTokenTransferCrud()
}
