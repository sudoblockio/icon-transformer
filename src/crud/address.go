package crud

import (
	"reflect"
	"sync"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/sudoblockio/icon-transformer/models"
)

// AddressCrud - type for address table model
type AddressCrud struct {
	db            *gorm.DB
	model         *models.Address
	modelORM      *models.AddressORM
	LoaderChannel chan *models.Address
}

var addressCrud *AddressCrud
var addressCrudOnce sync.Once

// GetAddressCrud - create and/or return the addresss table model
func GetAddressCrud() *AddressCrud {
	addressCrudOnce.Do(func() {
		dbConn := getPostgresConn()
		if dbConn == nil {
			zap.S().Fatal("Cannot connect to postgres database")
		}

		addressCrud = &AddressCrud{
			db:            dbConn,
			model:         &models.Address{},
			modelORM:      &models.AddressORM{},
			LoaderChannel: make(chan *models.Address, 1),
		}

		err := addressCrud.Migrate()
		if err != nil {
			zap.S().Fatal("AddressCrud: Unable migrate postgres table: ", err.Error())
		}

		StartAddressLoader()
	})

	return addressCrud
}

// Migrate - migrate addresss table
func (m *AddressCrud) Migrate() error {
	// Only using AddressRawORM (ORM version of the proto generated struct) to create the TABLE
	err := m.db.AutoMigrate(m.modelORM) // Migration and Index creation
	return err
}
func (m *AddressCrud) TableName() string {
	return m.modelORM.TableName()
}

// SelectMany - select many from addreses table
func (m *AddressCrud) SelectMany(
	limit int,
	skip int,
) (*[]models.Address, error) {
	db := m.db

	// Set table
	db = db.Model(&models.Address{})

	// Limit
	db = db.Limit(limit)

	// Skip
	if skip != 0 {
		db = db.Offset(skip)
	}

	addresses := &[]models.Address{}
	db = db.Find(addresses)

	return addresses, db.Error
}

// SelectAllAddresses - select all addresses from addreses table - Only for routine / recovery
func (m *AddressCrud) SelectAllAddresses() (*[]models.Address, error) {
	db := m.db
	db = db.Model(&models.Address{})
	addresses := &[]models.Address{}
	db = db.Find(addresses)
	return addresses, db.Error
}

// SelectOneAddress - select one address from addresses table
func (m *AddressCrud) SelectOneAddress(_address string) (*models.Address, error) {
	db := m.db
	db = db.Model(&models.Address{})
	db = db.Where("address = ?", _address)

	address := &models.Address{}

	db = db.First(address)

	return address, db.Error
}

// SelectAllAddresses - select all addresses from addreses table - Only for routine / recovery
func (m *AddressCrud) SelectAllTokenContracts() (*[]models.Address, error) {
	db := m.db
	db = db.Model(&models.Address{})
	addresses := &[]models.Address{}
	db = db.Where("is_token=true")
	db = db.Find(addresses)
	return addresses, db.Error
}

// SelectMany - select many from addreses table
func (m *AddressCrud) SelectPReps() (*[]models.Address, error) {
	db := m.db

	// Set table
	db = db.Model(&models.Address{})

	// Get preps
	db = db.Where("is_prep=true")

	addresses := &[]models.Address{}
	db = db.Find(addresses)

	return addresses, db.Error
}

// SelectCount - select from blockCounts table
// NOTE very slow operation
func (m *AddressCrud) CountAll() (int64, error) {
	db := m.db

	// Set table
	db = db.Model(&models.Address{})

	count := int64(0)
	db = db.Count(&count)

	return count, db.Error
}

// SelectCount - select from blockCounts table
// NOTE very slow operation
func (m *AddressCrud) CountContracts() (int64, error) {
	db := m.db

	// Set table
	db = db.Model(&models.Address{})

	// Is Contract
	db = db.Where("is_contract = ?", true)

	count := int64(0)
	db = db.Count(&count)

	return count, db.Error
}

func (m *AddressCrud) UpsertOne(
	address *models.Address,
) error {
	db := m.db

	// map[string]interface{}
	updateOnConflictValues := extractFilledFieldsFromModel(
		reflect.ValueOf(*address),
		reflect.TypeOf(*address),
	)

	// Upsert
	db = db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "address"}}, // NOTE set to primary keys for table
		DoUpdates: clause.Assignments(updateOnConflictValues),
	}).Create(address)

	return db.Error
}

func (m *AddressCrud) UpsertOneColsE(
	address *models.Address, cols []string,
) error {
	db := m.db

	db = db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "address"}}, // NOTE set to primary keys for table
		DoUpdates: clause.AssignmentColumns(cols),
	}).Create(address)

	return db.Error
}

func (m *AddressCrud) UpsertOneCols(address *models.Address, cols []string) {
	err := GetAddressCrud().UpsertOneColsE(address, cols)
	zap.S().Debug(
		"Loader=Address",
		" Address=", address.Address,
		" - Upserted",
	)
	if err != nil {
		// Postgres error
		zap.S().Fatal(
			"Loader=Address",
			" Address=", address.Address,
			" - Error: ", err.Error(),
		)
	}
}

// StartAddressLoader starts loader
func StartAddressLoader() {
	go func() {

		for {
			// Read address
			newAddress := <-GetAddressCrud().LoaderChannel

			//////////////////////
			// Load to postgres //
			//////////////////////
			err := GetAddressCrud().UpsertOne(newAddress)
			zap.S().Debug(
				"Loader=Address",
				" Address=", newAddress.Address,
				" - Upserted",
			)
			if err != nil {
				// Postgres error
				zap.S().Fatal(
					"Loader=Address",
					" Address=", newAddress.Address,
					" - Error: ", err.Error(),
				)
			}
		}
	}()
}

// StartAddressBalanceLoader starts loader
func StartAddressBalanceLoader() {
	go func() {
		for {
			newAddress := <-GetAddressCrud().LoaderChannel
			GetAddressCrud().UpsertOneCols(newAddress, []string{"address", "balance"})
		}
	}()
}
