package crud

import (
	"github.com/sudoblockio/icon-transformer/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Generic struct to hold common objects when doing crud
type Crud[Model any, ModelOrm any] struct {
	db            *gorm.DB
	model         *Model
	modelORM      *ModelOrm
	LoaderChannel chan *Model
	columns       []string
	primaryKeys   []clause.Column
}

// Migrate - migrate table
func (m *Crud[M, O]) Migrate() {
	err := m.db.AutoMigrate(m.modelORM)
	if err != nil {
		zap.S().Fatal("TokenTransferCrud: Unable migrate postgres table: ", err.Error())
	}
}

// TODO: Implement wrapper to pull indexes from gorm tags and code gen this statement

// CreateIndexes - Create indexes
func (m *Crud[Model, ModelOrm]) CreateIndexes(statement string) {
	db := m.db
	db = db.Exec(statement)
	if db.Error != nil {
		zap.S().Warn("TokenTransferCrud: Unable migrate postgres table: ", db.Error.Error())
	}
}

func (m *Crud[Model, ModelOrm]) UpsertOne(
	block *models.Block,
	cols []string,
) error {
	db := m.db
	db = db.Clauses(clause.OnConflict{
		Columns:   m.primaryKeys,
		DoUpdates: clause.AssignmentColumns(cols),
	}).Create(block)
	return db.Error
}

func (m *Crud[Model, ModelOrm]) UpsertMany(
	values []*Model,
	cols []string,
) error {
	db := m.db
	db = db.Clauses(clause.OnConflict{
		Columns:   m.primaryKeys,
		DoUpdates: clause.AssignmentColumns(cols),
	}).Create(values)

	if db.Error == nil {
		return nil
	}

	gormErr := getGormError(db)
	if gormErr.Code == "21000" {
		for _, v := range values {
			db := m.db
			db = db.Clauses(clause.OnConflict{
				Columns:   m.primaryKeys,
				DoUpdates: clause.AssignmentColumns(cols),
			}).Create(v)

			if db.Error != nil {
				zap.S().Fatal(db.Error.Error())
			}
			return nil
		}
	}

	return db.Error
}

// Count - count all entries in token_transfers table
// NOTE this function will take a long time
func (m *Crud[Model, ModelOrm]) Count() (int64, error) {
	db := m.db
	db = db.Model(&Crud[Model, ModelOrm]{})
	var count int64
	db = db.Count(&count)
	return count, db.Error
}

// CountByToAddress - count all entries in token_transfers table to address
func (m *Crud[Model, ModelOrm]) CountByToAddress(address string) (int64, error) {
	db := m.db
	db = db.Model(&models.TokenTransfer{})
	db = db.Where("to_address = ?", address)
	var count int64
	db = db.Count(&count)
	return count, db.Error
}

// CountByAddress - count all entries in token_transfers table
func (m *Crud[Model, ModelOrm]) CountByAddress(address string) (int64, error) {
	db := m.db
	db = db.Model(&models.TokenTransfer{})
	db = db.Where("to_address = ? or from_address = ?", address, address)
	var count int64
	db = db.Count(&count)
	return count, db.Error
}

// Count all entries in token_transfers table by token contract
func (m *Crud[Model, ModelOrm]) CountByTokenContract(address string) (int64, error) {
	db := m.db
	db = db.Model(&models.TokenTransfer{})
	db = db.Where("token_contract_address = ?", address)
	var count int64
	db = db.Count(&count)
	return count, db.Error
}
