package crud

import (
	"sync"

	"github.com/sudoblockio/icon-transformer/models"
	"go.uber.org/zap"
)

//// TokenTransferCrud - type for tokenTransfer table model
//type TokenTransferCrud struct {
//	db            *gorm.DB
//	model         *models.TokenTransfer
//	modelORM      *models.TokenTransferORM
//	LoaderChannel chan *models.TokenTransfer
//}

var tokenTransferCrud *Crud[models.TokenTransfer, models.TokenTransferORM]
var tokenTransferCrudOnce sync.Once

// GetTokenTransferCrud - create and/or return the tokenTransfers table model
func GetTokenTransferCrud() *Crud[models.TokenTransfer, models.TokenTransferORM] {
	tokenTransferCrudOnce.Do(func() {
		dbConn := getPostgresConn()
		if dbConn == nil {
			zap.S().Fatal("Cannot connect to postgres database")
		}

		tokenTransferCrud = &Crud[models.TokenTransfer, models.TokenTransferORM]{
			db:            dbConn,
			model:         &models.TokenTransfer{},
			modelORM:      &models.TokenTransferORM{},
			LoaderChannel: make(chan *models.TokenTransfer, 1000),
			columns:       getModelColumnNames(models.TokenTransfer{}),
			primaryKeys:   getModelPrimaryKeys(models.TokenTransferORM{}),
		}

		tokenTransferCrud.Migrate()
		tokenTransferCrud.CreateIndexes("CREATE INDEX IF NOT EXISTS token_transfer_idx_token_contract_address_block_number ON public.token_transfers USING btree (token_contract_address, block_number)")

		startBatchLoader(
			tokenTransferCrud.LoaderChannel,
			tokenTransferCrud.UpsertMany,
			tokenTransferCrud.columns,
		)

	})

	return tokenTransferCrud
}

//// Count - count all entries in token_transfers table
//// NOTE this function will take a long time
//func (m *Crud[models.TokenTransfer, models.TokenTransferORM]) Count() (int64, error) {
//	db := m.db
//
//	// Set table
//	db = db.Model(&models.TokenTransfer{})
//
//	// Count
//	var count int64
//	db = db.Count(&count)
//
//	return count, db.Error
//}

//// CountByToAddress - count all entries in token_transfers table to address
//func (m *Crud[Model, ModelOrm]) CountByToAddress(address string) (int64, error) {
//	db := m.db
//	db = db.Model(&models.TokenTransfer{})
//	db = db.Where("to_address = ?", address)
//	var count int64
//	db = db.Count(&count)
//	return count, db.Error
//}
//
//// CountByAddress - count all entries in token_transfers table
//func (m *Crud[Model, ModelOrm]) CountByAddress(address string) (int64, error) {
//	db := m.db
//	db = db.Model(&models.TokenTransfer{})
//	db = db.Where("to_address = ? or from_address = ?", address, address)
//	var count int64
//	db = db.Count(&count)
//	return count, db.Error
//}
//
//// Count all entries in token_transfers table by token contract
//func (m *Crud[Model, ModelOrm]) CountByTokenContract(address string) (int64, error) {
//	db := m.db
//	db = db.Model(&models.TokenTransfer{})
//	db = db.Where("token_contract_address = ?", address)
//	var count int64
//	db = db.Count(&count)
//	return count, db.Error
//}

//// Migrate - migrate tokenTransfers table
//func (m *TokenTransferCrud) Migrate() error {
//	// Only using TokenTransferRawORM (ORM version of the proto generated struct) to create the TABLE
//	err := m.db.AutoMigrate(m.modelORM) // Migration and Index creation
//	return err
//}
//func (m *TokenTransferCrud) TableName() string {
//	return m.modelORM.TableName()
//}

//func (m *TokenTransferCrud) CreateIndices() error {
//	db := m.db
//
//	// Create indices
//	db = db.Exec("CREATE INDEX IF NOT EXISTS token_transfer_idx_token_contract_address_block_number ON public.token_transfers USING btree (token_contract_address, block_number)")
//
//	return db.Error
//}

//func (m *TokenTransferCrud) UpsertOne(
//	tokenTransfer *models.TokenTransfer,
//) error {
//	db := m.db
//
//	// map[string]interface{}
//	updateOnConflictValues := extractFilledFieldsFromModel(
//		reflect.ValueOf(*tokenTransfer),
//		reflect.TypeOf(*tokenTransfer),
//	)
//
//	// Upsert
//	db = db.Clauses(clause.OnConflict{
//		Columns:   []clause.Column{{Name: "transaction_hash"}, {Name: "log_index"}}, // NOTE set to primary keys for table
//		DoUpdates: clause.Assignments(updateOnConflictValues),
//	}).Create(tokenTransfer)
//
//	return db.Error
//}

//// StartTokenTransferLoader starts loader
//func StartTokenTransferLoader() {
//	go func() {
//
//		for {
//			// Read tokenTransfer
//			newTokenTransfer := <-GetTokenTransferCrud().LoaderChannel
//
//			err := retryLoader(
//				newTokenTransfer,
//				GetTokenTransferCrud().UpsertOne,
//				5,
//				config.Config.DbRetrySleep,
//			)
//			if err != nil {
//				// Postgres error
//				zap.S().Fatal(err.Error())
//			}
//		}
//	}()
//}
