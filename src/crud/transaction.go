package crud

import (
	"reflect"
	"sync"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/sudoblockio/icon-transformer/models"
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

		err = transactionCrud.CreateIndices()
		if err != nil {
			zap.S().Warn("TransactionCrud: Unable to create indices: ", err.Error())
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

func (m *TransactionCrud) CreateIndices() error {
	db := m.db

	zap.S().Info("Creating transaction_idx_block_number_type_hash index")
	// Create indices
	db = db.Exec("CREATE INDEX IF NOT EXISTS transaction_idx_block_number_type_hash ON public.transactions USING btree (block_number, type, hash)")
	zap.S().Info("Finished creating transaction_idx_block_number_type_hash index")

	return db.Error
}

// Count all entries in transactions table both internal and regular. See CountInternal and CountRegular for those types.
// NOTE this function will take a long time
func (m *TransactionCrud) Count() (int64, error) {
	db := m.db
	db = db.Model(&models.Transaction{})
	var count int64
	db = db.Count(&count)
	return count, db.Error
}

// Count for regular transactions.
// NOTE this function will take a long time
func (m *TransactionCrud) CountRegular() (int64, error) {
	db := m.db
	db = db.Model(&models.Transaction{}).Where("type = 'transaction'")
	var count int64
	db = db.Count(&count)
	return count, db.Error
}

// Count for internal transactions.
// NOTE this function will take a long time
func (m *TransactionCrud) CountInternal() (int64, error) {
	db := m.db
	db = db.Model(&models.Transaction{}).Where("type = 'log'")
	var count int64
	db = db.Count(&count)
	return count, db.Error
}

// Count for internal transactions.
// NOTE this function will take a long time
func (m *TransactionCrud) CountInternalByAddress(address string) (int64, error) {
	db := m.db
	db = db.Model(&models.Transaction{}).Where("type = 'log'")
	db = db.Model(&models.Transaction{}).Where("to_address = ? or from_address = ?", address, address)
	var count int64
	db = db.Count(&count)
	return count, db.Error
}

// Count for regular transactions.
// NOTE this function will take a long time
func (m *TransactionCrud) CountRegularByAddress(address string) (int64, error) {
	db := m.db
	db = db.Model(&models.Transaction{}).Where("type='transaction'")
	db = db.Model(&models.Transaction{}).Where("to_address = ? or from_address = ?", address, address)
	var count int64
	db = db.Count(&count)
	return count, db.Error
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

// SelectManyContractCreations - select all the contract creation events
// Returns: models, error (if present)
func (m *TransactionCrud) SelectManyContractCreations(
	address string,
	tranaction_types []int32,
) (*[]models.Transaction, error) {
	db := m.db

	// Set table
	db = db.Model(&[]models.Transaction{})

	// Latest transactions first
	db = db.Order("block_number ASC")

	// Address
	db = db.Where("score_address = ?", address)

	// Type
	db = db.Where("transaction_type IN ?", tranaction_types)

	transactions := &[]models.Transaction{}
	db = db.Find(transactions)

	return transactions, db.Error
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

func (m *TransactionCrud) UpsertOneColsE(
	transaction *models.Transaction, cols []string,
) error {
	db := m.db

	// Upsert
	db = db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "hash"}, {Name: "log_index"}},
		DoUpdates: clause.AssignmentColumns(cols),
	}).Create(transaction)

	return db.Error
}

func (m *TransactionCrud) UpsertOneCols(transaction *models.Transaction, cols []string) {
	err := GetTransactionCrud().UpsertOneColsE(transaction, cols)
	zap.S().Debug(
		"Loader=Transaction",
		" Hash=", transaction.Hash,
		" - Upserted",
	)
	if err != nil {
		// Postgres error
		zap.S().Fatal(
			"Loader=Address",
			" Hash=", transaction.Hash,
			" - Error: ", err.Error(),
		)
	}
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
