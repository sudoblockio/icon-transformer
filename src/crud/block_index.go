package crud

import (
	"github.com/sudoblockio/icon-transformer/models"
	"sync"
)

import (
	"errors"
)

var blockIndexCrudOnce sync.Once
var BlockIndexCrud *Crud[models.BlockIndex, models.BlockIndexORM]

// GetBlockIndexCrud - create and/or return the blockIndexs table Model
func GetBlockIndexCrud() *Crud[models.BlockIndex, models.BlockIndexORM] {
	blockIndexCrudOnce.Do(func() {
		BlockIndexCrud = GetCrud(models.BlockIndex{}, models.BlockIndexORM{})

		BlockIndexCrud.Migrate()

		BlockIndexCrud.MakeStartLoaderChannel()
	})

	return BlockIndexCrud
}

func SelectOneBlockIndex(c *Crud[models.BlockIndex, models.BlockIndexORM], number int64) (*models.BlockIndex, error) {
	db := c.db
	blockIndex := &models.BlockIndex{}
	//db = db.Model(&models.BlockIndex{})
	db = db.Where("number = ?", number)
	db = db.First(blockIndex)
	return blockIndex, db.Error
}

func InsertOneBlockIndex(c *Crud[models.BlockIndex, models.BlockIndexORM], number int64) error {
	db := c.db
	blockIndex := &models.BlockIndex{Number: number}
	db = db.Model(&models.BlockIndex{})
	db = db.Create(blockIndex)
	return db.Error
}

//func (m *blockIndexCrud) FindMissing(lowestNumber int64, highestNumber int64) ([]int64, error) {
//	db := m.db
//
//	// https://stackoverflow.com/a/12444165
//	// https://stackoverflow.com/a/32072586
//	responseInterface := []interface{}{}
//	db = db.Raw(
//		`SELECT s.i AS missing_numbers
//		FROM generate_series(?::int,?::int) s(i)
//		WHERE NOT EXISTS (SELECT 1 FROM block_indices WHERE number = s.i);`,
//		lowestNumber,
//		highestNumber,
//	).Scan(&responseInterface)
//	if db.Error != nil {
//		return []int64{}, db.Error
//	}
//
//	numbers := []int64{}
//
//	for _, r := range responseInterface {
//		number, ok := r.(int64)
//		if ok == false {
//			return []int64{}, errors.New("Could not cast int64")
//		}
//
//		numbers = append(numbers, number)
//	}
//
//	return numbers, nil
//}

func FindMissingBlocks(m *Crud[models.BlockIndex, models.BlockIndexORM], lowestNumber int64, highestNumber int64) ([]int64, error) {
	db := m.db
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

//func (m *BlockIndexCrud) SelectHighestNumber() (int64, error) {
//	db := m.db
//
//	// Set table
//	db = db.Model(&[]models.BlockIndex{})
//
//	// Latest blocks first
//	db = db.Order("number desc")
//
//	block := &models.BlockIndex{}
//	db = db.First(block)
//
//	return block.Number, db.Error
//}

func SelectHighestNumber(m *Crud[models.BlockIndex, models.BlockIndexORM]) (int64, error) {
	db := m.db
	db = db.Model(&[]models.BlockIndex{})
	db = db.Order("number desc")
	block := &models.BlockIndex{}
	db = db.First(block)
	return block.Number, db.Error
}

//func (m *BlockIndexCrud) SelectLowestNumber() (int64, error) {
//	db := m.db
//
//	// Set table
//	db = db.Model(&[]models.BlockIndex{})
//
//	// Latest blocks first
//	db = db.Order("number asc")
//
//	block := &models.BlockIndex{}
//	db = db.First(block)
//
//	return block.Number, db.Error
//}

func SelectLowestNumber(m *Crud[models.BlockIndex, models.BlockIndexORM]) (int64, error) {
	db := m.db
	db = db.Model(&[]models.BlockIndex{})
	db = db.Order("number asc")
	block := &models.BlockIndex{}
	db = db.First(block)
	return block.Number, db.Error
}
