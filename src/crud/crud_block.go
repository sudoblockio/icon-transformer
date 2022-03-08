package crud

import (
	"reflect"
	"sync"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/sudoblockio/icon-go-worker/models"
)

// BlockCrud - type for block table model
type BlockCrud struct {
	db            *gorm.DB
	model         *models.Block
	modelORM      *models.BlockORM
	LoaderChannel chan interface{}
}

var blockCrud *BlockCrud
var blockCrudOnce sync.Once

// GetBlockCrud - create and/or return the blocks table model
func GetBlockCrud() *BlockCrud {
	blockCrudOnce.Do(func() {
		dbConn := getPostgresConn()
		if dbConn == nil {
			zap.S().Fatal("Cannot connect to postgres database")
		}

		blockCrud = &BlockCrud{
			db:            dbConn,
			model:         &models.Block{},
			modelORM:      &models.BlockORM{},
			LoaderChannel: make(chan interface{}, 1),
		}

		err := blockCrud.Migrate()
		if err != nil {
			zap.S().Fatal("BlockCrud: Unable migrate postgres table: ", err.Error())
		}

		StartBlockLoader()
	})

	return blockCrud
}

// Migrate - migrate blocks table
func (m *BlockCrud) Migrate() error {
	// Only using BlockRawORM (ORM version of the proto generated struct) to create the TABLE
	err := m.db.AutoMigrate(m.modelORM) // Migration and Index creation
	return err
}
func (m *BlockCrud) TableName() string {
	return m.modelORM.TableName()
}

func (m *BlockCrud) UpsertOne(
	block *models.Block,
) error {
	db := m.db

	// map[string]interface{}
	updateOnConflictValues := extractFilledFieldsFromModel(
		reflect.ValueOf(*block),
		reflect.TypeOf(*block),
	)

	// Upsert
	db = db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "number"}}, // NOTE set to primary keys for table
		DoUpdates: clause.Assignments(updateOnConflictValues),
	}).Create(block)

	return db.Error
}

// StartBlockLoader starts loader
func StartBlockLoader() {
	go func() {

		for {
			// Read block
			newBlockInterface := <-GetBlockCrud().LoaderChannel
			newBlock, ok := newBlockInterface.(models.Block)
			if ok == false {
				zap.S().Warn("Loader=Block - Error: Invalid type")
				continue
			}

			//////////////////////
			// Load to postgres //
			//////////////////////
			err := GetBlockCrud().UpsertOne(&newBlock)
			zap.S().Debug("Loader=Block, Number=", newBlock.Number, " - Upserted")
			if err != nil {
				// Postgres error
				zap.S().Fatal("Loader=Block, Number=", newBlock.Number, " - Error: ", err.Error())
			}
		}
	}()
}
