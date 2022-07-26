package crud

import (
	"github.com/sudoblockio/icon-transformer/config"
	"reflect"
	"sync"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/sudoblockio/icon-transformer/models"
)

// TokenAddressCrud - type for addressToken table model
type TokenAddressCrud struct {
	db            *gorm.DB
	model         *models.TokenAddress
	modelORM      *models.TokenAddressORM
	LoaderChannel chan *models.TokenAddress
}

var addressTokenCrud *TokenAddressCrud
var addressTokenCrudOnce sync.Once

// GetTokenAddressCrud - create and/or return the addressToken table model
func GetTokenAddressCrud() *TokenAddressCrud {
	addressTokenCrudOnce.Do(func() {
		dbConn := getPostgresConn()
		if dbConn == nil {
			zap.S().Fatal("Cannot connect to postgres database")
		}

		addressTokenCrud = &TokenAddressCrud{
			db:            dbConn,
			model:         &models.TokenAddress{},
			modelORM:      &models.TokenAddressORM{},
			LoaderChannel: make(chan *models.TokenAddress, 1),
		}

		err := addressTokenCrud.Migrate()
		if err != nil {
			zap.S().Fatal("TokenAddressCrud: Unable migrate postgres table: ", err.Error())
		}

		StartTokenAddressLoader()
	})

	return addressTokenCrud
}

// Migrate - migrate addressToken table
func (m *TokenAddressCrud) Migrate() error {
	// Only using TokenAddressRawORM (ORM version of the proto generated struct) to create the TABLE
	err := m.db.AutoMigrate(m.modelORM) // Migration and Index creation
	return err
}
func (m *TokenAddressCrud) TableName() string {
	return m.modelORM.TableName()
}

// SelectMany - select many from addreses table
func (m *TokenAddressCrud) SelectMany(
	limit int,
	skip int,
) (*[]models.TokenAddress, error) {
	db := m.db

	// Set table
	db = db.Model(&models.TokenAddress{})

	// Limit
	db = db.Limit(limit)

	// Skip
	if skip != 0 {
		db = db.Offset(skip)
	}

	tokenAddresses := &[]models.TokenAddress{}
	db = db.Find(tokenAddresses)

	return tokenAddresses, db.Error
}

// CountByTokenContractAddress - select many from addreses table
func (m *TokenAddressCrud) CountByTokenContractAddress() (map[string]int64, error) {
	db := m.db

	// Set table
	db = db.Model(&models.TokenAddress{})

	// Count
	db = db.Raw("SELECT COUNT(address) as count, token_contract_address FROM token_addresses GROUP BY token_contract_address")

	// Rows
	rows, err := db.Rows()
	if err != nil {
		return nil, err
	}
	countByTokenContractAddress := map[string]int64{}
	for rows.Next() {
		count := int64(0)
		tokenContractAddress := ""
		rows.Scan(&count, &tokenContractAddress)

		countByTokenContractAddress[tokenContractAddress] = count
	}

	return countByTokenContractAddress, nil
}

func (m *TokenAddressCrud) UpsertOne(
	addressToken *models.TokenAddress,
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

func (m *TokenAddressCrud) UpsertOneColsE(
	addressToken *models.TokenAddress, cols []string,
) error {
	db := m.db

	// Upsert
	db = db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "address"}, {Name: "token_contract_address"}},
		DoUpdates: clause.AssignmentColumns(cols),
	}).Create(addressToken)

	return db.Error
}

func (m *TokenAddressCrud) UpsertOneCols(addressToken *models.TokenAddress, cols []string) {
	//err := GetTokenAddressCrud().UpsertOneColsE(addressToken, cols)
	err := retryCrudColumns(
		addressToken,
		cols,
		GetTokenAddressCrud().UpsertOneColsE,
		5,
		config.Config.DbRetrySleep,
	)
	zap.S().Debug(
		"Loader=Address",
		" Address=", addressToken.Address,
		" - Upserted",
	)
	if err != nil {
		// Postgres error
		zap.S().Fatal(
			"Loader=Address",
			" Address=", addressToken.Address,
			" - Error: ", err.Error(),
		)
	}
}

// StartTokenAddressLoader starts loader
func StartTokenAddressLoader() {
	go func() {

		for {
			// Read addressToken
			newTokenAddress := <-GetTokenAddressCrud().LoaderChannel

			err := retryLoader(
				newTokenAddress,
				GetTokenAddressCrud().UpsertOne,
				5,
				config.Config.DbRetrySleep,
			)
			if err != nil {
				zap.S().Fatal(err.Error())
			}
		}
	}()
}
