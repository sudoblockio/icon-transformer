package crud

import (
	"reflect"
	"sync"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/sudoblockio/icon-transformer/models"
)

// TransactionCreateScoreCrud - type for transactionCreateScore table model
type TransactionCreateScoreCrud struct {
	db            *gorm.DB
	model         *models.TransactionCreateScore
	modelORM      *models.TransactionCreateScoreORM
	LoaderChannel chan *models.TransactionCreateScore
}

var transactionCreateScoreCrud *TransactionCreateScoreCrud
var transactionCreateScoreCrudOnce sync.Once

// GetTransactionCreateScoreCrud - create and/or return the transactionCreateScores table model
func GetTransactionCreateScoreCrud() *TransactionCreateScoreCrud {
	transactionCreateScoreCrudOnce.Do(func() {
		dbConn := getPostgresConn()
		if dbConn == nil {
			zap.S().Fatal("Cannot connect to postgres database")
		}

		transactionCreateScoreCrud = &TransactionCreateScoreCrud{
			db:            dbConn,
			model:         &models.TransactionCreateScore{},
			modelORM:      &models.TransactionCreateScoreORM{},
			LoaderChannel: make(chan *models.TransactionCreateScore, 1),
		}

		err := transactionCreateScoreCrud.Migrate()
		if err != nil {
			zap.S().Fatal("TransactionCreateScoreCrud: Unable migrate postgres table: ", err.Error())
		}

		StartTransactionCreateScoreLoader()
	})

	return transactionCreateScoreCrud
}

// Migrate - migrate transactionCreateScores table
func (m *TransactionCreateScoreCrud) Migrate() error {
	// Only using TransactionCreateScoreRawORM (ORM version of the proto generated struct) to create the TABLE
	err := m.db.AutoMigrate(m.modelORM) // Migration and Index creation
	return err
}
func (m *TransactionCreateScoreCrud) TableName() string {
	return m.modelORM.TableName()
}

// SelectMany - select many from addreses table
func (m *TransactionCreateScoreCrud) SelectMany(
	limit int,
	skip int,
) (*[]models.TransactionCreateScore, error) {
	db := m.db

	// Set table
	db = db.Model(&models.TransactionCreateScore{})

	// Limit
	db = db.Limit(limit)

	// Skip
	if skip != 0 {
		db = db.Offset(skip)
	}

	transactionCreateScores := &[]models.TransactionCreateScore{}
	db = db.Find(transactionCreateScores)

	return transactionCreateScores, db.Error
}

func (m *TransactionCreateScoreCrud) UpsertOne(
	transactionCreateScore *models.TransactionCreateScore,
) error {
	db := m.db

	// map[string]interface{}
	updateOnConflictValues := extractFilledFieldsFromModel(
		reflect.ValueOf(*transactionCreateScore),
		reflect.TypeOf(*transactionCreateScore),
	)

	// Upsert
	db = db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "creation_transaction_hash"}}, // NOTE set to primary keys for table
		DoUpdates: clause.Assignments(updateOnConflictValues),
	}).Create(transactionCreateScore)

	return db.Error
}

// StartTransactionCreateScoreLoader starts loader
func StartTransactionCreateScoreLoader() {
	go func() {

		for {
			// Read transactionCreateScore
			newTransactionCreateScore := <-GetTransactionCreateScoreCrud().LoaderChannel

			//////////////////////
			// Load to postgres //
			//////////////////////
			err := GetTransactionCreateScoreCrud().UpsertOne(newTransactionCreateScore)
			zap.S().Debug("Loader=TransactionCreateScore, CreationTransactionHash=", newTransactionCreateScore.CreationTransactionHash, " - Upserted")
			if err != nil {
				// Postgres error
				zap.S().Fatal("Loader=TransactionCreateScore, CreationTransactionHash=", newTransactionCreateScore.CreationTransactionHash, " - Error: ", err.Error())
			}
		}
	}()
}
