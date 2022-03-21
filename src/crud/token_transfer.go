package crud

import (
	"reflect"
	"sync"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/sudoblockio/icon-go-worker/models"
)

// TokenTransferCrud - type for tokenTransfer table model
type TokenTransferCrud struct {
	db            *gorm.DB
	model         *models.TokenTransfer
	modelORM      *models.TokenTransferORM
	LoaderChannel chan *models.TokenTransfer
}

var tokenTransferCrud *TokenTransferCrud
var tokenTransferCrudOnce sync.Once

// GetTokenTransferCrud - create and/or return the tokenTransfers table model
func GetTokenTransferCrud() *TokenTransferCrud {
	tokenTransferCrudOnce.Do(func() {
		dbConn := getPostgresConn()
		if dbConn == nil {
			zap.S().Fatal("Cannot connect to postgres database")
		}

		tokenTransferCrud = &TokenTransferCrud{
			db:            dbConn,
			model:         &models.TokenTransfer{},
			modelORM:      &models.TokenTransferORM{},
			LoaderChannel: make(chan *models.TokenTransfer, 1),
		}

		err := tokenTransferCrud.Migrate()
		if err != nil {
			zap.S().Fatal("TokenTransferCrud: Unable migrate postgres table: ", err.Error())
		}

		StartTokenTransferLoader()
	})

	return tokenTransferCrud
}

// Migrate - migrate tokenTransfers table
func (m *TokenTransferCrud) Migrate() error {
	// Only using TokenTransferRawORM (ORM version of the proto generated struct) to create the TABLE
	err := m.db.AutoMigrate(m.modelORM) // Migration and Index creation
	return err
}
func (m *TokenTransferCrud) TableName() string {
	return m.modelORM.TableName()
}

func (m *TokenTransferCrud) UpsertOne(
	tokenTransfer *models.TokenTransfer,
) error {
	db := m.db

	// map[string]interface{}
	updateOnConflictValues := extractFilledFieldsFromModel(
		reflect.ValueOf(*tokenTransfer),
		reflect.TypeOf(*tokenTransfer),
	)

	// Upsert
	db = db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "transaction_hash"}, {Name: "log_index"}}, // NOTE set to primary keys for table
		DoUpdates: clause.Assignments(updateOnConflictValues),
	}).Create(tokenTransfer)

	return db.Error
}

// StartTokenTransferLoader starts loader
func StartTokenTransferLoader() {
	go func() {

		for {
			// Read tokenTransfer
			newTokenTransfer := <-GetTokenTransferCrud().LoaderChannel

			//////////////////////
			// Load to postgres //
			//////////////////////
			err := GetTokenTransferCrud().UpsertOne(newTokenTransfer)
			zap.S().Debug(
				"Loader=TokenTransfer",
				" TransactionHash=", newTokenTransfer.TransactionHash,
				" LogIndex=", newTokenTransfer.LogIndex,
				" - Upserted",
			)
			if err != nil {
				// Postgres error
				zap.S().Fatal(
					"Loader=TokenTransfer",
					" TransactionHash=", newTokenTransfer.TransactionHash,
					" LogIndex=", newTokenTransfer.LogIndex,
					" - Error: ", err.Error(),
				)
			}
		}
	}()
}
