package crud

import (
	"reflect"
	"sync"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/sudoblockio/icon-transformer/models"
)

// TransactionInternalByAddressCrud - type for transactionInternalByAddress table model
type TransactionInternalByAddressCrud struct {
	db            *gorm.DB
	model         *models.TransactionInternalByAddress
	modelORM      *models.TransactionInternalByAddressORM
	LoaderChannel chan *models.TransactionInternalByAddress
}

var transactionInternalByAddressCrud *TransactionInternalByAddressCrud
var transactionInternalByAddressCrudOnce sync.Once

// GetTransactionInternalByAddressCrud - create and/or return the transactionInternalByAddresss table model
func GetTransactionInternalByAddressCrud() *TransactionInternalByAddressCrud {
	transactionInternalByAddressCrudOnce.Do(func() {
		dbConn := getPostgresConn()
		if dbConn == nil {
			zap.S().Fatal("Cannot connect to postgres database")
		}

		transactionInternalByAddressCrud = &TransactionInternalByAddressCrud{
			db:            dbConn,
			model:         &models.TransactionInternalByAddress{},
			modelORM:      &models.TransactionInternalByAddressORM{},
			LoaderChannel: make(chan *models.TransactionInternalByAddress, 1),
		}

		err := transactionInternalByAddressCrud.Migrate()
		if err != nil {
			zap.S().Fatal("TransactionInternalByAddressCrud: Unable migrate postgres table: ", err.Error())
		}

		StartTransactionInternalByAddressLoader()
	})

	return transactionInternalByAddressCrud
}

// Migrate - migrate transactionInternalByAddresss table
func (m *TransactionInternalByAddressCrud) Migrate() error {
	// Only using TransactionInternalByAddressRawORM (ORM version of the proto generated struct) to create the TABLE
	err := m.db.AutoMigrate(m.modelORM) // Migration and Index creation
	return err
}
func (m *TransactionInternalByAddressCrud) TableName() string {
	return m.modelORM.TableName()
}

func (m *TransactionInternalByAddressCrud) UpsertOne(
	transactionInternalByAddress *models.TransactionInternalByAddress,
) error {
	db := m.db

	// map[string]interface{}
	updateOnConflictValues := extractFilledFieldsFromModel(
		reflect.ValueOf(*transactionInternalByAddress),
		reflect.TypeOf(*transactionInternalByAddress),
	)

	// Upsert
	db = db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "transaction_hash"}, {Name: "log_index"}, {Name: "address"}}, // NOTE set to primary keys for table
		DoUpdates: clause.Assignments(updateOnConflictValues),
	}).Create(transactionInternalByAddress)

	return db.Error
}

// StartTransactionInternalByAddressLoader starts loader
func StartTransactionInternalByAddressLoader() {
	go func() {

		for {
			// Read transactionInternalByAddress
			newTransactionInternalByAddress := <-GetTransactionInternalByAddressCrud().LoaderChannel

			//////////////////////
			// Load to postgres //
			//////////////////////
			err := GetTransactionInternalByAddressCrud().UpsertOne(newTransactionInternalByAddress)
			zap.S().Debug(
				"Loader=TransactionInternalByAddress",
				" TransactionHash=", newTransactionInternalByAddress.TransactionHash,
				" LogIndex=", newTransactionInternalByAddress.LogIndex,
				" Address=", newTransactionInternalByAddress.Address,
				" - Upserted",
			)
			if err != nil {
				// Postgres error
				zap.S().Fatal(
					"Loader=TransactionInternalByAddress",
					" TransactionHash=", newTransactionInternalByAddress.TransactionHash,
					" LogIndex=", newTransactionInternalByAddress.LogIndex,
					" Address=", newTransactionInternalByAddress.Address,
					" - Error: ", err.Error(),
				)
			}
		}
	}()
}
