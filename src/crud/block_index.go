package crud

import (
	"errors"
	"sync"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/sudoblockio/icon-transformer/models"
)

// BlockIndexCrud - type for blockIndex table model
type BlockIndexCrud struct {
	db       *gorm.DB
	model    *models.BlockIndex
	modelORM *models.BlockIndexORM
}

var blockIndexCrud *BlockIndexCrud
var blockIndexCrudOnce sync.Once

// GetBlockIndexCrud - create and/or return the blockIndexs table model
func GetBlockIndexCrud() *BlockIndexCrud {
	blockIndexCrudOnce.Do(func() {
		dbConn := getPostgresConn()
		if dbConn == nil {
			zap.S().Fatal("Cannot connect to postgres database")
		}

		blockIndexCrud = &BlockIndexCrud{
			db:       dbConn,
			model:    &models.BlockIndex{},
			modelORM: &models.BlockIndexORM{},
		}

		err := blockIndexCrud.Migrate()
		if err != nil {
			zap.S().Fatal("BlockIndexCrud: Unable migrate postgres table: ", err.Error())
		}
	})

	return blockIndexCrud
}

// Count - count all entries in blocks table
func (m *BlockIndexCrud) Count() (int64, error) {
	db := m.db

	// Set table
	db = db.Model(&models.BlockIndex{})

	// Count
	var count int64
	db = db.Count(&count)

	return count, db.Error
}

// Migrate - migrate blockIndexs table
func (m *BlockIndexCrud) Migrate() error {
	// Only using BlockIndexRawORM (ORM version of the proto generated struct) to create the TABLE
	err := m.db.AutoMigrate(m.modelORM) // Migration and Index creation
	return err
}

func (m *BlockIndexCrud) TableName() string {
	return m.modelORM.TableName()
}

// SelectOne - Select one from blockIndex table
func (m *BlockIndexCrud) SelectOne(number int64) (*models.BlockIndex, error) {
	db := m.db

	// Set table
	db = db.Model(&models.BlockIndex{})

	// Number
	db = db.Where("number = ?", number)

	blockIndex := &models.BlockIndex{}
	db = db.First(blockIndex)

	return blockIndex, db.Error
}

// Insert - Insert blockIndex into table
func (m *BlockIndexCrud) InsertOne(blockIndex *models.BlockIndex) error {
	db := m.db

	// Set table
	db = db.Model(&models.BlockIndex{})

	db = db.Create(blockIndex)

	return db.Error
}

func (m *BlockIndexCrud) FindMissing() ([]int64, error) {
	db := m.db

	highestNumber, err := m.SelectHighestNumber()
	if err != nil {
		return []int64{}, err
	}

	lowestNumber, err := m.SelectLowestNumber()
	if err != nil {
		return []int64{}, err
	}

	// https://stackoverflow.com/a/12444165
	// https://stackoverflow.com/a/32072586
	responseInterface := []interface{}{}
	db = db.Raw(
		`SELECT s.i AS missing_numbers
		FROM generate_series(?::int,?::int) s(i)
		WHERE NOT EXISTS (SELECT 1 FROM block_indices WHERE number = s.i);`,
		lowestNumber,
		highestNumber,
	).Scan(&responseInterface)
	if db.Error != nil {
		return []int64{}, db.Error
	}

	numbers := []int64{}

	for _, r := range responseInterface {
		number, ok := r.(int64)
		if ok == false {
			return []int64{}, errors.New("Could not cast int64")
		}

		numbers = append(numbers, number)
	}

	return numbers, nil
}

func (m *BlockIndexCrud) SelectHighestNumber() (int64, error) {
	db := m.db

	// Set table
	db = db.Model(&[]models.BlockIndex{})

	// Latest blocks first
	db = db.Order("number desc")

	block := &models.BlockIndex{}
	db = db.First(block)

	return block.Number, db.Error
}

func (m *BlockIndexCrud) SelectLowestNumber() (int64, error) {
	db := m.db

	// Set table
	db = db.Model(&[]models.BlockIndex{})

	// Latest blocks first
	db = db.Order("number asc")

	block := &models.BlockIndex{}
	db = db.First(block)

	return block.Number, db.Error
}
