package crud

import (
	"errors"
	"reflect"
	"sync"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/sudoblockio/icon-transformer/models"
)

// BlockCrud - type for block table model
type BlockCrud struct {
	db            *gorm.DB
	model         *models.Block
	modelORM      *models.BlockORM
	LoaderChannel chan *models.Block
}

var blockCrud *BlockCrud
var blockCrudOnce sync.Once

// GetBlockCrud - create and/or return the blocks table model
func GetBlockCrud() *BlockCrud {
	blockCrudOnce.Do(func() {
		dbConn := getPostgresConn()
		if dbConn == nil {
			zap.S().Fatal("Cannot connect to postgres database")
		}

		blockCrud = &BlockCrud{
			db:            dbConn,
			model:         &models.Block{},
			modelORM:      &models.BlockORM{},
			LoaderChannel: make(chan *models.Block, 1),
		}

		err := blockCrud.Migrate()
		if err != nil {
			zap.S().Fatal("BlockCrud: Unable migrate postgres table: ", err.Error())
		}

		StartBlockLoader()
	})

	return blockCrud
}

// Migrate - migrate blocks table
func (m *BlockCrud) Migrate() error {
	// Only using BlockRawORM (ORM version of the proto generated struct) to create the TABLE
	err := m.db.AutoMigrate(m.modelORM) // Migration and Index creation
	return err
}

func (m *BlockCrud) TableName() string {
	return m.modelORM.TableName()
}

func (m *BlockCrud) FindMissing() ([]int64, error) {
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
		WHERE NOT EXISTS (SELECT 1 FROM blocks WHERE number = s.i);`,
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

func (m *BlockCrud) SelectHighestNumber() (int64, error) {
	db := m.db

	// Set table
	db = db.Model(&[]models.Block{})

	// Latest blocks first
	db = db.Order("number desc")

	block := &models.Block{}
	db = db.First(block)

	return block.Number, db.Error
}

func (m *BlockCrud) SelectLowestNumber() (int64, error) {
	db := m.db

	// Set table
	db = db.Model(&[]models.Block{})

	// Latest blocks first
	db = db.Order("number asc")

	block := &models.Block{}
	db = db.First(block)

	return block.Number, db.Error
}

func (m *BlockCrud) UpsertOne(
	block *models.Block,
) error {
	db := m.db

	// map[string]interface{}
	updateOnConflictValues := extractFilledFieldsFromModel(
		reflect.ValueOf(*block),
		reflect.TypeOf(*block),
	)

	// Upsert
	db = db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "number"}}, // NOTE set to primary keys for table
		DoUpdates: clause.Assignments(updateOnConflictValues),
	}).Create(block)

	return db.Error
}

// StartBlockLoader starts loader
func StartBlockLoader() {
	go func() {

		for {
			// Read block
			newBlock := <-GetBlockCrud().LoaderChannel

			//////////////////////
			// Load to postgres //
			//////////////////////
			err := GetBlockCrud().UpsertOne(newBlock)
			zap.S().Debug(
				"Loader=Block",
				" Number=", newBlock.Number,
				" - Upserted",
			)
			if err != nil {
				// Postgres error
				zap.S().Fatal(
					"Loader=Block",
					" Number=", newBlock.Number,
					" - Error: ", err.Error(),
				)
			}
		}
	}()
}
