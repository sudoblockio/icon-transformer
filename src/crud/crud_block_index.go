package crud

import (
	"sync"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/sudoblockio/icon-go-worker/models"
)

// BlockIndexCrud - type for blockIndex table model
type BlockIndexCrud struct {
	db       *gorm.DB
	model    *models.BlockIndex
	modelORM *models.BlockIndexORM
}

var blockIndexCrud *BlockIndexCrud
var blockIndexCrudOnce sync.Once

// GetBlockIndexCrud - create and/or return the blockIndexs table model
func GetBlockIndexCrud() *BlockIndexCrud {
	blockIndexCrudOnce.Do(func() {
		dbConn := getPostgresConn()
		if dbConn == nil {
			zap.S().Fatal("Cannot connect to postgres database")
		}

		blockIndexCrud = &BlockIndexCrud{
			db:       dbConn,
			model:    &models.BlockIndex{},
			modelORM: &models.BlockIndexORM{},
		}

		err := blockIndexCrud.Migrate()
		if err != nil {
			zap.S().Fatal("BlockIndexCrud: Unable migrate postgres table: ", err.Error())
		}
	})

	return blockIndexCrud
}

// Migrate - migrate blockIndexs table
func (m *BlockIndexCrud) Migrate() error {
	// Only using BlockIndexRawORM (ORM version of the proto generated struct) to create the TABLE
	err := m.db.AutoMigrate(m.modelORM) // Migration and Index creation
	return err
}

func (m *BlockIndexCrud) TableName() string {
	return m.modelORM.TableName()
}

// SelectOne - Select one from blockIndex table
func (m *BlockIndexCrud) SelectOne(number int64) (*models.BlockIndex, error) {
	db := m.db

	// Set table
	db = db.Model(&models.BlockIndex{})

	// Number
	db = db.Where("number = ?", number)

	blockIndex := &models.BlockIndex{}
	db = db.First(blockIndex)

	return blockIndex, db.Error
}

// Insert - Insert blockIndex into table
func (m *BlockIndexCrud) InsertOne(blockIndex *models.BlockIndex) error {
	db := m.db

	// Set table
	db = db.Model(&models.BlockIndex{})

	db = db.Create(blockIndex)

	return db.Error
}
