package crud

import (
	"reflect"
	"sync"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/sudoblockio/icon-go-worker/models"
)

// LogCrud - type for log table model
type LogCrud struct {
	db            *gorm.DB
	model         *models.Log
	modelORM      *models.LogORM
	LoaderChannel chan *models.Log
}

var logCrud *LogCrud
var logCrudOnce sync.Once

// GetLogCrud - create and/or return the logs table model
func GetLogCrud() *LogCrud {
	logCrudOnce.Do(func() {
		dbConn := getPostgresConn()
		if dbConn == nil {
			zap.S().Fatal("Cannot connect to postgres database")
		}

		logCrud = &LogCrud{
			db:            dbConn,
			model:         &models.Log{},
			modelORM:      &models.LogORM{},
			LoaderChannel: make(chan *models.Log, 1),
		}

		err := logCrud.Migrate()
		if err != nil {
			zap.S().Fatal("LogCrud: Unable migrate postgres table: ", err.Error())
		}

		StartLogLoader()
	})

	return logCrud
}

// Migrate - migrate logs table
func (m *LogCrud) Migrate() error {
	// Only using LogRawORM (ORM version of the proto generated struct) to create the TABLE
	err := m.db.AutoMigrate(m.modelORM) // Migration and Index creation
	return err
}
func (m *LogCrud) TableName() string {
	return m.modelORM.TableName()
}

func (m *LogCrud) UpsertOne(
	log *models.Log,
) error {
	db := m.db

	// map[string]interface{}
	updateOnConflictValues := extractFilledFieldsFromModel(
		reflect.ValueOf(*log),
		reflect.TypeOf(*log),
	)

	// Upsert
	db = db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "transaction_hash"}, {Name: "log_index"}}, // NOTE set to primary keys for table
		DoUpdates: clause.Assignments(updateOnConflictValues),
	}).Create(log)

	return db.Error
}

// StartLogLoader starts loader
func StartLogLoader() {
	go func() {

		for {
			// Read log
			newLog := <-GetLogCrud().LoaderChannel

			//////////////////////
			// Load to postgres //
			//////////////////////
			err := GetLogCrud().UpsertOne(newLog)
			zap.S().Debug(
				"Loader=Log",
				" TransactionHash=", newLog.TransactionHash,
				" LogIndex=", newLog.LogIndex,
				" - Upserted",
			)
			if err != nil {
				// Postgres error
				zap.S().Fatal(
					"Loader=Log",
					" TransactionHash=", newLog.TransactionHash,
					" LogIndex=", newLog.LogIndex,
					" - Error: ", err.Error(),
				)
			}
		}
	}()
}
