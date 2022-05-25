package crud

import (
	"reflect"
	"sync"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/sudoblockio/icon-transformer/models"
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

		err = tokenTransferCrud.CreateIndices()
		if err != nil {
			zap.S().Warn("TokenTransferCrud: Unable migrate postgres table: ", err.Error())
		}

		StartTokenTransferLoader()
	})

	return tokenTransferCrud
}

// Count - count all entries in token_transfers table
// NOTE this function will take a long time
func (m *TokenTransferCrud) Count() (int64, error) {
	db := m.db

	// Set table
	db = db.Model(&models.TokenTransfer{})

	// Count
	var count int64
	db = db.Count(&count)

	return count, db.Error
}

// CountByToAddress - count all entries in token_transfers table to address
func (m *TokenTransferCrud) CountByToAddress(address string) (int64, error) {
	db := m.db
	db = db.Model(&models.TokenTransfer{})
	db = db.Where("to_address = ?", address)
	var count int64
	db = db.Count(&count)
	return count, db.Error
}

// CountByAddress - count all entries in token_transfers table
func (m *TokenTransferCrud) CountByAddress(address string) (int64, error) {
	db := m.db
	db = db.Model(&models.TokenTransfer{})
	db = db.Where("to_address = ? or from_address = ?", address, address)
	var count int64
	db = db.Count(&count)
	return count, db.Error
}

// Count all entries in token_transfers table by token contract
func (m *TokenTransferCrud) CountByTokenContract(address string) (int64, error) {
	db := m.db
	db = db.Model(&models.TokenTransfer{})
	db = db.Where("token_contract_address = ?", address)
	var count int64
	db = db.Count(&count)
	return count, db.Error
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

func (m *TokenTransferCrud) CreateIndices() error {
	db := m.db

	// Create indices
	db = db.Exec("CREATE INDEX IF NOT EXISTS token_transfer_idx_token_contract_address_block_number ON public.token_transfers USING btree (token_contract_address, block_number)")

	return db.Error
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
