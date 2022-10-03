package crud

import (
	"sync"
	"time"

	"github.com/sudoblockio/icon-transformer/models"
)

var blockCrudOnce sync.Once
var BlockCrud *Crud[models.Block, models.BlockORM]

// GetBlockCrud - create and/or return the blocks table model
func GetBlockCrud() *Crud[models.Block, models.BlockORM] {
	blockCrudOnce.Do(func() {
		BlockCrud = GetCrud(models.Block{}, models.BlockORM{})

		BlockCrud.Migrate()
		BlockCrud.dbBufferWait = 10 * time.Millisecond

		BlockCrud.MakeStartLoaderChannel()
	})
	return BlockCrud
}

func InitBlockCrud() {
	GetBlockCrud()
}
