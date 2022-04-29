package crud

import (
	"reflect"
	"sync"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/sudoblockio/icon-transformer/models"
)

// MissingBlockCrud - type for missingBlock table model
type MissingBlockCrud struct {
	db            *gorm.DB
	model         *models.MissingBlock
	modelORM      *models.MissingBlockORM
	LoaderChannel chan *models.MissingBlock
}

var missingBlockCrud *MissingBlockCrud
var missingBlockCrudOnce sync.Once

// GetMissingBlockCrud - create and/or return the missingBlocks table model
func GetMissingBlockCrud() *MissingBlockCrud {
	missingBlockCrudOnce.Do(func() {
		dbConn := getPostgresConn()
		if dbConn == nil {
			zap.S().Fatal("Cannot connect to postgres database")
		}

		missingBlockCrud = &MissingBlockCrud{
			db:            dbConn,
			model:         &models.MissingBlock{},
			modelORM:      &models.MissingBlockORM{},
			LoaderChannel: make(chan *models.MissingBlock, 1),
		}

		err := missingBlockCrud.Migrate()
		if err != nil {
			zap.S().Fatal("MissingBlockCrud: Unable migrate postgres table: ", err.Error())
		}

		StartMissingBlockLoader()
	})

	return missingBlockCrud
}

// Migrate - migrate missingBlocks table
func (m *MissingBlockCrud) Migrate() error {
	// Only using MissingBlockRawORM (ORM version of the proto generated struct) to create the TABLE
	err := m.db.AutoMigrate(m.modelORM) // Migration and Index creation
	return err
}
func (m *MissingBlockCrud) TableName() string {
	return m.modelORM.TableName()
}

func (m *MissingBlockCrud) DeleteAll() error {
	db := m.db

	// Set table
	db = db.Model(&models.MissingBlock{})

	// Delete
	db = db.Delete(&models.MissingBlock{})

	return db.Error
}

func (m *MissingBlockCrud) UpsertOne(
	missingBlock *models.MissingBlock,
) error {
	db := m.db

	// map[string]interface{}
	updateOnConflictValues := extractFilledFieldsFromModel(
		reflect.ValueOf(*missingBlock),
		reflect.TypeOf(*missingBlock),
	)

	// Upsert
	db = db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "number"}}, // NOTE set to primary keys for table
		DoUpdates: clause.Assignments(updateOnConflictValues),
	}).Create(missingBlock)

	return db.Error
}

// StartMissingBlockLoader starts loader
func StartMissingBlockLoader() {
	go func() {

		for {
			// Read missingBlock
			newMissingBlock := <-GetMissingBlockCrud().LoaderChannel

			//////////////////////
			// Load to postgres //
			//////////////////////
			err := GetMissingBlockCrud().UpsertOne(newMissingBlock)
			zap.S().Debug(
				"Loader=MissingBlock",
				" Number=", newMissingBlock.Number,
				" - Upserted",
			)
			if err != nil {
				// Postgres error
				zap.S().Fatal(
					"Loader=MissingBlock",
					" Number=", newMissingBlock.Number,
					" - Error: ", err.Error(),
				)
			}
		}
	}()
}
