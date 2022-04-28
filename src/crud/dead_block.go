package crud

import (
	"reflect"
	"sync"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/sudoblockio/icon-transformer/models"
)

// DeadBlockCrud - type for deadBlock table model
type DeadBlockCrud struct {
	db            *gorm.DB
	model         *models.DeadBlock
	modelORM      *models.DeadBlockORM
	LoaderChannel chan *models.DeadBlock
}

var deadBlockCrud *DeadBlockCrud
var deadBlockCrudOnce sync.Once

// GetDeadBlockCrud - create and/or return the deadBlocks table model
func GetDeadBlockCrud() *DeadBlockCrud {
	deadBlockCrudOnce.Do(func() {
		dbConn := getPostgresConn()
		if dbConn == nil {
			zap.S().Fatal("Cannot connect to postgres database")
		}

		deadBlockCrud = &DeadBlockCrud{
			db:            dbConn,
			model:         &models.DeadBlock{},
			modelORM:      &models.DeadBlockORM{},
			LoaderChannel: make(chan *models.DeadBlock, 1),
		}

		err := deadBlockCrud.Migrate()
		if err != nil {
			zap.S().Fatal("DeadBlockCrud: Unable migrate postgres table: ", err.Error())
		}

		StartDeadBlockLoader()
	})

	return deadBlockCrud
}

// Migrate - migrate deadBlocks table
func (m *DeadBlockCrud) Migrate() error {
	// Only using DeadBlockRawORM (ORM version of the proto generated struct) to create the TABLE
	err := m.db.AutoMigrate(m.modelORM) // Migration and Index creation
	return err
}
func (m *DeadBlockCrud) TableName() string {
	return m.modelORM.TableName()
}

func (m *DeadBlockCrud) UpsertOne(
	deadBlock *models.DeadBlock,
) error {
	db := m.db

	// map[string]interface{}
	updateOnConflictValues := extractFilledFieldsFromModel(
		reflect.ValueOf(*deadBlock),
		reflect.TypeOf(*deadBlock),
	)

	// Upsert
	db = db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "topic"}, {Name: "partition"}, {Name: "offset"}}, // NOTE set to primary keys for table
		DoUpdates: clause.Assignments(updateOnConflictValues),
	}).Create(deadBlock)

	return db.Error
}

// StartDeadBlockLoader starts loader
func StartDeadBlockLoader() {
	go func() {

		for {
			// Read deadBlock
			newDeadBlock := <-GetDeadBlockCrud().LoaderChannel

			//////////////////////
			// Load to postgres //
			//////////////////////
			err := GetDeadBlockCrud().UpsertOne(newDeadBlock)
			zap.S().Debug(
				"Loader=DeadBlock",
				" Topic=", newDeadBlock.Topic,
				" Partition=", newDeadBlock.Partition,
				" Offset=", newDeadBlock.Offset,
				" - Upserted",
			)
			if err != nil {
				// Postgres error
				zap.S().Fatal(
					"Loader=DeadBlock",
					" Topic=", newDeadBlock.Topic,
					" Partition=", newDeadBlock.Partition,
					" Offset=", newDeadBlock.Offset,
					" - Error: ", err.Error(),
				)
			}
		}
	}()
}
