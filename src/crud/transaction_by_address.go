package crud

import (
	"reflect"
	"sync"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/sudoblockio/icon-transformer/models"
)

// TransactionByAddressCrud - type for transactionByAddress table model
type TransactionByAddressCrud struct {
	db            *gorm.DB
	model         *models.TransactionByAddress
	modelORM      *models.TransactionByAddressORM
	LoaderChannel chan *models.TransactionByAddress
}

var transactionByAddressCrud *TransactionByAddressCrud
var transactionByAddressCrudOnce sync.Once

// GetTransactionByAddressCrud - create and/or return the transactionByAddresss table model
func GetTransactionByAddressCrud() *TransactionByAddressCrud {
	transactionByAddressCrudOnce.Do(func() {
		dbConn := getPostgresConn()
		if dbConn == nil {
			zap.S().Fatal("Cannot connect to postgres database")
		}

		transactionByAddressCrud = &TransactionByAddressCrud{
			db:            dbConn,
			model:         &models.TransactionByAddress{},
			modelORM:      &models.TransactionByAddressORM{},
			LoaderChannel: make(chan *models.TransactionByAddress, 1),
		}

		err := transactionByAddressCrud.Migrate()
		if err != nil {
			zap.S().Fatal("TransactionByAddressCrud: Unable migrate postgres table: ", err.Error())
		}

		err = transactionByAddressCrud.CreateIndices()
		if err != nil {
			zap.S().Warn("TransactionByAddressCrud: Unable to create indices: ", err.Error())
		}

		StartTransactionByAddressLoader()
	})

	return transactionByAddressCrud
}

// Migrate - migrate transactionByAddresss table
func (m *TransactionByAddressCrud) Migrate() error {
	// Only using TransactionByAddressRawORM (ORM version of the proto generated struct) to create the TABLE
	err := m.db.AutoMigrate(m.modelORM) // Migration and Index creation
	return err
}

func (m *TransactionByAddressCrud) TableName() string {
	return m.modelORM.TableName()
}

func (m *TransactionByAddressCrud) CreateIndices() error {
	db := m.db

	// Create indices
	db = db.Exec("CREATE INDEX IF NOT EXISTS transaction_by_addresses_idx_address_block_number_index ON public.transaction_by_addresses (address asc, block_number desc)")

	return db.Error
}

// Count for regular transactions.
// NOTE this function will take a long time
func (m *TransactionByAddressCrud) CountByAddress(address string) (int64, error) {
	db := m.db
	db = db.Model(&models.TransactionByAddress{}).Where("address = ?", address)
	var count int64
	db = db.Count(&count)
	return count, db.Error
}

func (m *TransactionByAddressCrud) UpsertOne(
	transactionByAddress *models.TransactionByAddress,
) error {
	db := m.db

	// map[string]interface{}
	updateOnConflictValues := extractFilledFieldsFromModel(
		reflect.ValueOf(*transactionByAddress),
		reflect.TypeOf(*transactionByAddress),
	)

	// Upsert
	db = db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "transaction_hash"}, {Name: "address"}}, // NOTE set to primary keys for table
		DoUpdates: clause.Assignments(updateOnConflictValues),
	}).Create(transactionByAddress)

	return db.Error
}

// StartTransactionByAddressLoader starts loader
func StartTransactionByAddressLoader() {
	go func() {

		for {
			// Read transactionByAddress
			newTransactionByAddress := <-GetTransactionByAddressCrud().LoaderChannel

			//////////////////////
			// Load to postgres //
			//////////////////////
			err := GetTransactionByAddressCrud().UpsertOne(newTransactionByAddress)
			zap.S().Debug(
				"Loader=TransactionByAddress",
				" TransactionHash=", newTransactionByAddress.TransactionHash,
				" Address=", newTransactionByAddress.Address,
				" - Upserted",
			)
			if err != nil {
				// Postgres error
				zap.S().Fatal(
					"Loader=TransactionByAddress",
					" TransactionHash=", newTransactionByAddress.TransactionHash,
					" Address=", newTransactionByAddress.Address,
					" - Error: ", err.Error(),
				)
			}
		}
	}()
}
