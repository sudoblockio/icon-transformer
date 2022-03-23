package crud

import (
	"reflect"
	"sync"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/sudoblockio/icon-go-worker/models"
)

// TransactionCrud - type for transaction table model
type TransactionCrud struct {
	db            *gorm.DB
	model         *models.Transaction
	modelORM      *models.TransactionORM
	LoaderChannel chan *models.Transaction
}

var transactionCrud *TransactionCrud
var transactionCrudOnce sync.Once

// GetTransactionCrud - create and/or return the transactions table model
func GetTransactionCrud() *TransactionCrud {
	transactionCrudOnce.Do(func() {
		dbConn := getPostgresConn()
		if dbConn == nil {
			zap.S().Fatal("Cannot connect to postgres database")
		}

		transactionCrud = &TransactionCrud{
			db:            dbConn,
			model:         &models.Transaction{},
			modelORM:      &models.TransactionORM{},
			LoaderChannel: make(chan *models.Transaction, 1),
		}

		err := transactionCrud.Migrate()
		if err != nil {
			zap.S().Fatal("TransactionCrud: Unable migrate postgres table: ", err.Error())
		}

		StartTransactionLoader()
	})

	return transactionCrud
}

// Migrate - migrate transactions table
func (m *TransactionCrud) Migrate() error {
	// Only using TransactionRawORM (ORM version of the proto generated struct) to create the TABLE
	err := m.db.AutoMigrate(m.modelORM) // Migration and Index creation
	return err
}
func (m *TransactionCrud) TableName() string {
	return m.modelORM.TableName()
}

// SelectOne - select from transactions table
func (m *TransactionCrud) SelectOne(
	hash string,
	logIndex int32, // Used for internal transactions
) (*models.Transaction, error) {
	db := m.db

	// Set table
	db = db.Model(&[]models.Transaction{})

	// Hash
	db = db.Where("hash = ?", hash)

	// Log Index
	db = db.Where("log_index = ?", logIndex)

	transaction := &models.Transaction{}
	db = db.First(transaction)

	return transaction, db.Error
}

func (m *TransactionCrud) UpsertOne(
	transaction *models.Transaction,
) error {
	db := m.db

	// map[string]interface{}
	updateOnConflictValues := extractFilledFieldsFromModel(
		reflect.ValueOf(*transaction),
		reflect.TypeOf(*transaction),
	)

	// Upsert
	db = db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "hash"}, {Name: "log_index"}}, // NOTE set to primary keys for table
		DoUpdates: clause.Assignments(updateOnConflictValues),
	}).Create(transaction)

	return db.Error
}

// StartTransactionLoader starts loader
func StartTransactionLoader() {
	go func() {

		for {
			// Read transaction
			newTransaction := <-GetTransactionCrud().LoaderChannel

			//////////////////////
			// Load to postgres //
			//////////////////////
			err := GetTransactionCrud().UpsertOne(newTransaction)
			zap.S().Debug("Loader=Transaction, Hash=", newTransaction.Hash, ", LogIndex=", newTransaction.LogIndex, " - Upserted")
			if err != nil {
				// Postgres error
				zap.S().Fatal("Loader=Transaction, Hash=", newTransaction.Hash, ", LogIndex=", newTransaction.LogIndex, " - Error: ", err.Error())
			}
		}
	}()
}