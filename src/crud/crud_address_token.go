package crud

import (
	"reflect"
	"sync"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/sudoblockio/icon-go-worker/models"
)

// AddressTokenCrud - type for addressToken table model
type AddressTokenCrud struct {
	db            *gorm.DB
	model         *models.AddressToken
	modelORM      *models.AddressTokenORM
	LoaderChannel chan *models.AddressToken
}

var addressTokenCrud *AddressTokenCrud
var addressTokenCrudOnce sync.Once

// GetAddressTokenCrud - create and/or return the addressToken table model
func GetAddressTokenCrud() *AddressTokenCrud {
	addressTokenCrudOnce.Do(func() {
		dbConn := getPostgresConn()
		if dbConn == nil {
			zap.S().Fatal("Cannot connect to postgres database")
		}

		addressTokenCrud = &AddressTokenCrud{
			db:            dbConn,
			model:         &models.AddressToken{},
			modelORM:      &models.AddressTokenORM{},
			LoaderChannel: make(chan *models.AddressToken, 1),
		}

		err := addressTokenCrud.Migrate()
		if err != nil {
			zap.S().Fatal("AddressTokenCrud: Unable migrate postgres table: ", err.Error())
		}

		StartAddressTokenLoader()
	})

	return addressTokenCrud
}

// Migrate - migrate addressToken table
func (m *AddressTokenCrud) Migrate() error {
	// Only using AddressTokenRawORM (ORM version of the proto generated struct) to create the TABLE
	err := m.db.AutoMigrate(m.modelORM) // Migration and Index creation
	return err
}
func (m *AddressTokenCrud) TableName() string {
	return m.modelORM.TableName()
}

func (m *AddressTokenCrud) UpsertOne(
	addressToken *models.AddressToken,
) error {
	db := m.db

	// map[string]interface{}
	updateOnConflictValues := extractFilledFieldsFromModel(
		reflect.ValueOf(*addressToken),
		reflect.TypeOf(*addressToken),
	)

	// Upsert
	db = db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "address"}, {Name: "token_contract_address"}}, // NOTE set to primary keys for table
		DoUpdates: clause.Assignments(updateOnConflictValues),
	}).Create(addressToken)

	return db.Error
}

// StartAddressTokenLoader starts loader
func StartAddressTokenLoader() {
	go func() {

		for {
			// Read addressToken
			newAddressToken := <-GetAddressTokenCrud().LoaderChannel

			//////////////////////
			// Load to postgres //
			//////////////////////
			err := GetAddressTokenCrud().UpsertOne(newAddressToken)
			zap.S().Debug(
				"Loader=AddressToken",
				" Address=", newAddressToken.Address,
				" TokenContractAddress=", newAddressToken.TokenContractAddress,
				" - Upserted",
			)
			if err != nil {
				// Postgres error
				zap.S().Fatal(
					"Loader=AddressToken",
					" Address=", newAddressToken.Address,
					" TokenContractAddress=", newAddressToken.TokenContractAddress,
					" - Error: ", err.Error(),
				)
			}
		}
	}()
}
