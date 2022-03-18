package crud

import (
	"reflect"
	"sync"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/sudoblockio/icon-go-worker/models"
)

// TokenHolderCrud - type for tokenHolder table model
type TokenHolderCrud struct {
	db            *gorm.DB
	model         *models.TokenHolder
	modelORM      *models.TokenHolderORM
	LoaderChannel chan *models.TokenHolder
}

var tokenHolderCrud *TokenHolderCrud
var tokenHolderCrudOnce sync.Once

// GetTokenHolderCrud - create and/or return the tokenHolders table model
func GetTokenHolderCrud() *TokenHolderCrud {
	tokenHolderCrudOnce.Do(func() {
		dbConn := getPostgresConn()
		if dbConn == nil {
			zap.S().Fatal("Cannot connect to postgres database")
		}

		tokenHolderCrud = &TokenHolderCrud{
			db:            dbConn,
			model:         &models.TokenHolder{},
			modelORM:      &models.TokenHolderORM{},
			LoaderChannel: make(chan *models.TokenHolder, 1),
		}

		err := tokenHolderCrud.Migrate()
		if err != nil {
			zap.S().Fatal("TokenHolderCrud: Unable migrate postgres table: ", err.Error())
		}

		StartTokenHolderLoader()
	})

	return tokenHolderCrud
}

// Migrate - migrate tokenHolders table
func (m *TokenHolderCrud) Migrate() error {
	// Only using TokenHolderRawORM (ORM version of the proto generated struct) to create the TABLE
	err := m.db.AutoMigrate(m.modelORM) // Migration and Index creation
	return err
}
func (m *TokenHolderCrud) TableName() string {
	return m.modelORM.TableName()
}

func (m *TokenHolderCrud) UpsertOne(
	tokenHolder *models.TokenHolder,
) error {
	db := m.db

	// map[string]interface{}
	updateOnConflictValues := extractFilledFieldsFromModel(
		reflect.ValueOf(*tokenHolder),
		reflect.TypeOf(*tokenHolder),
	)

	// Upsert
	db = db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "token_contract_address"}, {Name: "holder_address"}}, // NOTE set to primary keys for table
		DoUpdates: clause.Assignments(updateOnConflictValues),
	}).Create(tokenHolder)

	return db.Error
}

// StartTokenHolderLoader starts loader
func StartTokenHolderLoader() {
	go func() {

		for {
			// Read tokenHolder
			newTokenHolder := <-GetTokenHolderCrud().LoaderChannel

			//////////////////////
			// Load to postgres //
			//////////////////////
			err := GetTokenHolderCrud().UpsertOne(newTokenHolder)
			zap.S().Debug(
				"Loader=TokenHolder",
				" TokenContractAddress=", newTokenHolder.TokenContractAddress,
				" HolderAddress=", newTokenHolder.HolderAddress,
				" - Upserted",
			)
			if err != nil {
				// Postgres error
				zap.S().Fatal(
					"Loader=TokenHolder",
					" TokenContractAddress=", newTokenHolder.TokenContractAddress,
					" HolderAddress=", newTokenHolder.HolderAddress,
					" - Error: ", err.Error(),
				)
			}
		}
	}()
}
