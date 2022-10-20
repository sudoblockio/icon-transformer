package crud

import (
	"github.com/sudoblockio/icon-transformer/models"
	"sync"
)

var blockCrudOnce sync.Once
var BlockCrud *Crud[models.Block, models.BlockORM]

// GetBlockCrud - create and/or return the blocks table Model
func GetBlockCrud() *Crud[models.Block, models.BlockORM] {
	blockCrudOnce.Do(func() {
		BlockCrud = GetCrud(models.Block{}, models.BlockORM{})

		BlockCrud.Migrate()

		BlockCrud.MakeStartLoaderChannel()
	})
	return BlockCrud
}

func InitBlockCrud() {
	GetBlockCrud()
}
