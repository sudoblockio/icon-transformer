package crud

import (
	"reflect"
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
			LoaderChannel: make(chan *models.TokenTransferByAddress, 1),
		}

		err := tokenTransferByAddressCrud.Migrate()
		if err != nil {
			zap.S().Fatal("TokenTransferByAddressCrud: Unable migrate postgres table: ", err.Error())
		}

		StartTokenTransferByAddressLoader()
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

func (m *TokenTransferByAddressCrud) UpsertOne(
	tokenTransferByAddress *models.TokenTransferByAddress,
) error {
	db := m.db

	// map[string]interface{}
	updateOnConflictValues := extractFilledFieldsFromModel(
		reflect.ValueOf(*tokenTransferByAddress),
		reflect.TypeOf(*tokenTransferByAddress),
	)

	// Upsert
	db = db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "transaction_hash"}, {Name: "log_index"}, {Name: "address"}}, // NOTE set to primary keys for table
		DoUpdates: clause.Assignments(updateOnConflictValues),
	}).Create(tokenTransferByAddress)

	return db.Error
}

// StartTokenTransferByAddressLoader starts loader
func StartTokenTransferByAddressLoader() {
	go func() {

		for {
			// Read tokenTransferByAddress
			newTokenTransferByAddress := <-GetTokenTransferByAddressCrud().LoaderChannel

			//////////////////////
			// Load to postgres //
			//////////////////////
			err := GetTokenTransferByAddressCrud().UpsertOne(newTokenTransferByAddress)
			zap.S().Debug(
				"Loader=TokenTransferByAddress",
				" TransactionHash=", newTokenTransferByAddress.TransactionHash,
				" LogIndex=", newTokenTransferByAddress.LogIndex,
				" Address=", newTokenTransferByAddress.Address,
				" - Upserted",
			)
			if err != nil {
				// Postgres error
				zap.S().Fatal(
					"Loader=TokenTransferByAddress",
					" TransactionHash=", newTokenTransferByAddress.TransactionHash,
					" LogIndex=", newTokenTransferByAddress.LogIndex,
					" Address=", newTokenTransferByAddress.Address,
					" - Error: ", err.Error(),
				)
			}
		}
	}()
}
