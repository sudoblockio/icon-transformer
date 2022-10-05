package crud

import (
	"sync"

	"github.com/sudoblockio/icon-transformer/models"
)

var deadBlockCrudOnce sync.Once
var DeadBlockCrud *Crud[models.DeadBlock, models.DeadBlockORM]

// GetDeadBlockCrud - create and/or return the deadBlocks table Model
func GetDeadBlockCrud() *Crud[models.DeadBlock, models.DeadBlockORM] {
	deadBlockCrudOnce.Do(func() {
		DeadBlockCrud = GetCrud(models.DeadBlock{}, models.DeadBlockORM{})

		DeadBlockCrud.Migrate()

		DeadBlockCrud.MakeStartLoaderChannel()
	})

	return DeadBlockCrud
}

func InitDeadBlockCrud() {
	GetDeadBlockCrud()
}
