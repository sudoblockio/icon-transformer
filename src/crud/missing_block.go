package crud

import (
	"sync"

	"github.com/sudoblockio/icon-transformer/models"
)

var missingBlockCrudOnce sync.Once
var MissingBlockCrud *Crud[models.MissingBlock, models.MissingBlockORM]

// GetMissingBlockCrud - create and/or return the missingBlocks table model
func GetMissingBlockCrud() *Crud[models.MissingBlock, models.MissingBlockORM] {
	missingBlockCrudOnce.Do(func() {
		missingBlockCrudOnce.Do(func() {
			MissingBlockCrud = GetCrud(models.MissingBlock{}, models.MissingBlockORM{})

			MissingBlockCrud.Migrate()

			MissingBlockCrud.MakeStartLoaderChannel()
		})
	})

	return MissingBlockCrud
}
