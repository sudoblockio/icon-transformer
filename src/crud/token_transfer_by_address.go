package crud

import (
	"sync"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/sudoblockio/icon-transformer/models"
)

// TokenTransferByAddressCrud - type for tokenTransferByAddress table model
type TokenTransferByAddressCrud struct {
	db            *gorm.DB
	model         *models.TokenTransferByAddress
	modelORM      *models.TokenTransferByAddressORM
	LoaderChannel chan *models.TokenTransferByAddress
	columns       []string
	primaryKeys   []clause.Column
}

var tokenTransferByAddressCrud *TokenTransferByAddressCrud
var tokenTransferByAddressCrudOnce sync.Once

// GetTokenTransferByAddressCrud - create and/or return the tokenTransferByAddresss table model
func GetTokenTransferByAddressCrud() *TokenTransferByAddressCrud {
	tokenTransferByAddressCrudOnce.Do(func() {
		dbConn := getPostgresConn()
		if dbConn == nil {
			zap.S().Fatal("Cannot connect to postgres database")
		}

		tokenTransferByAddressCrud = &TokenTransferByAddressCrud{
			db:            dbConn,
			model:         &models.TokenTransferByAddress{},
			modelORM:      &models.TokenTransferByAddressORM{},
			LoaderChannel: make(chan *models.TokenTransferByAddress, 100),
			columns:       getModelColumnNames(models.TokenTransferByAddress{}),
			primaryKeys:   getModelPrimaryKeys(models.TokenTransferByAddressORM{}),
		}

		err := tokenTransferByAddressCrud.Migrate()
		if err != nil {
			zap.S().Fatal("TokenTransferByAddressCrud: Unable migrate postgres table: ", err.Error())
		}

		//StartTokenTransferByAddressLoader()
		startBatchLoader(
			tokenTransferByAddressCrud.LoaderChannel,
			tokenTransferByAddressCrud.UpsertMany,
			tokenTransferByAddressCrud.columns,
		)
	})

	return tokenTransferByAddressCrud
}

// Migrate - migrate tokenTransferByAddresss table
func (m *TokenTransferByAddressCrud) Migrate() error {
	// Only using TokenTransferByAddressRawORM (ORM version of the proto generated struct) to create the TABLE
	err := m.db.AutoMigrate(m.modelORM) // Migration and Index creation
	return err
}
func (m *TokenTransferByAddressCrud) TableName() string {
	return m.modelORM.TableName()
}

// CountByTokenTransferByAddress - select one from token_transfers_by_addres table
func (m *TokenTransferByAddressCrud) CountByTokenTransfersByAddress(address string) (int64, error) {
	db := m.db
	db = db.Model(&models.TokenTransferByAddress{}).Where("address = ?", address)
	var count int64
	db = db.Count(&count)
	return count, db.Error
}

func (m *TokenTransferByAddressCrud) UpsertMany(
	tokenTransferByAddress []*models.TokenTransferByAddress,
	cols []string,
) error {
	db := m.db
	db = db.Clauses(clause.OnConflict{
		Columns:   m.primaryKeys,
		DoUpdates: clause.AssignmentColumns(cols),
	}).Create(tokenTransferByAddress)

	if db.Error == nil {
		return nil
	}

	gormErr := getGormError(db)
	if gormErr.Code == "21000" {
		for _, v := range tokenTransferByAddress {
			db := m.db
			db = db.Clauses(clause.OnConflict{
				Columns:   m.primaryKeys,
				DoUpdates: clause.AssignmentColumns(cols),
			}).Create(v)

			if db.Error != nil {
				zap.S().Fatal(db.Error.Error())
			}
			return nil
		}
	}

	return db.Error
}

//func (m *TokenTransferByAddressCrud) UpsertOne(
//	tokenTransferByAddress *models.TokenTransferByAddress,
//) error {
//	db := m.db
//
//	// map[string]interface{}
//	updateOnConflictValues := extractFilledFieldsFromModel(
//		reflect.ValueOf(*tokenTransferByAddress),
//		reflect.TypeOf(*tokenTransferByAddress),
//	)
//
//	// Upsert
//	db = db.Clauses(clause.OnConflict{
//		Columns:   []clause.Column{{Name: "transaction_hash"}, {Name: "log_index"}, {Name: "address"}}, // NOTE set to primary keys for table
//		DoUpdates: clause.Assignments(updateOnConflictValues),
//	}).Create(tokenTransferByAddress)
//
//	return db.Error
//}

//// StartTokenTransferByAddressLoader starts loader
//func StartTokenTransferByAddressLoader() {
//	go func() {
//		for {
//			// Read tokenTransferByAddress
//			newTokenTransferByAddress := <-GetTokenTransferByAddressCrud().LoaderChannel
//			err := retryLoader(
//				newTokenTransferByAddress,
//				GetTokenTransferByAddressCrud().UpsertOne,
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
