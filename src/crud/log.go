package crud

import (
	"github.com/sudoblockio/icon-transformer/models"
	"sync"
)

var logCrudOnce sync.Once
var LogCrud *Crud[models.Log, models.LogORM]

// GetLogCrud - create and/or return the logs table Model
func GetLogCrud() *Crud[models.Log, models.LogORM] {
	logCrudOnce.Do(func() {
		LogCrud = GetCrud(models.Log{}, models.LogORM{})

		LogCrud.Migrate()
		LogCrud.CreateIndexes("CREATE INDEX IF NOT EXISTS log_idx_address_block_number ON public.logs USING btree (address, block_number)")

		LogCrud.MakeStartLoaderChannel()
	})

	return LogCrud
}

func InitLogCrud() {
	GetLogCrud()
}

//// NOTE this function will take a long time
//func (m *Crud[M, O]) CountLogsByAddress(address string) (int64, error) {
//	db := m.db
//	db = db.Model(&m.Model).Where("address = ?", address)
//	var count int64
//	db = db.Count(&count)
//	return count, db.Error
//}

//// Migrate - migrate logs table
//func (m *LogCrud) Migrate() error {
//	// Only using LogRawORM (ORM version of the proto generated struct) to create the TABLE
//	err := m.db.AutoMigrate(m.ModelORM) // Migration and Index creation
//	return err
//}
//
//func (m *LogCrud) TableName() string {
//	return m.ModelORM.TableName()
//}
//
//func (m *LogCrud) CreateIndices() error {
//	db := m.db
//
//	// Create indices
//	db = db.Exec("CREATE INDEX IF NOT EXISTS log_idx_address_block_number ON public.logs USING btree (address, block_number)")
//
//	return db.Error
//}

//func (m *LogCrud) UpsertOne(
//	log *models.Log,
//) error {
//	db := m.db
//
//	// map[string]interface{}
//	updateOnConflictValues := extractFilledFieldsFromModel(
//		reflect.ValueOf(*log),
//		reflect.TypeOf(*log),
//	)
//
//	// Upsert
//	db = db.Clauses(clause.OnConflict{
//		Columns:   []clause.Column{{Name: "transaction_hash"}, {Name: "log_index"}}, // NOTE set to primary keys for table
//		DoUpdates: clause.Assignments(updateOnConflictValues),
//	}).Create(log)
//
//	return db.Error
//}

//// StartLogLoader starts loader
//func StartLogLoader() {
//	go func() {
//
//		for {
//			// Read log
//			newLog := <-GetLogCrud().LoaderChannel
//			err := retryLoader(
//				newLog,
//				GetLogCrud().UpsertOne,
//				5,
//				config.Config.DbRetrySleep,
//			)
//			if err != nil {
//				zap.S().Fatal(err.Error())
//			}
//		}
//	}()
//}
