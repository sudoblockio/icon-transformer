package crud

import "github.com/sudoblockio/icon-transformer/models"

func (m *Crud[Model, ModelOrm]) SelectMany(
	limit int,
	skip int,
) (*[]Model, error) {
	db := m.db
	db = db.Model(&m.Model)
	db = db.Limit(limit)
	if skip != 0 {
		db = db.Offset(skip)
	}
	output := &[]Model{}
	db = db.Find(output)
	return output, db.Error
}

// SelectWhere - select with one where
func (m *Crud[Model, ModelOrm]) SelectWhere(column string, equals string) ([]*Model, error) {
	db := m.db
	db = db.Model(&m.Model)
	db = db.Where("? = ?", column, equals)

	var output []*Model
	db = db.Find(&output)
	return output, db.Error
}

// SelectWhere - select with one where
func (m *Crud[Model, ModelOrm]) SelectPrep() ([]*Model, error) {
	db := m.db
	db = db.Model(&m.Model)
	db = db.Where("is_prep = ?", true)

	var output []*Model
	db = db.Find(&output)
	return output, db.Error
}

// SelectOneWhere - select one address from addresses table
func (m *Crud[Model, ModelOrm]) SelectOneWhere(column string, equals string) (*Model, error) {
	db := m.db
	db = db.Model(&m.Model)
	db = db.Where("? = ?", column, equals)

	var output *Model
	db = db.First(&output)
	return output, db.Error
}

// SelectTransactionOne - select from transactions table
func (m *Crud[Model, ModelOrm]) SelectTransactionOne(
	hash string,
	logIndex int32, // Used for internal transactions
) (*models.Transaction, error) {
	db := m.db

	// Set table
	db = db.Model(&[]models.Transaction{})

	// Hash
	db = db.Where("hash = ?", hash)

	// Log Index
	db = db.Where("log_index = ?", logIndex)

	transaction := &models.Transaction{}
	db = db.First(transaction)

	return transaction, db.Error
}

// SelectManyContractCreations - select all the contract creation events
// Returns: models, error (if present)
func (m *Crud[Model, ModelOrm]) SelectManyContractCreations(
	address string,
	tranaction_types []int32,
) (*[]models.Transaction, error) {
	db := m.db

	// Set table
	db = db.Model(&[]models.Transaction{})

	// Latest transactions first
	db = db.Order("block_number ASC")

	// Address
	db = db.Where("score_address = ?", address)

	// Type
	db = db.Where("transaction_type IN ?", tranaction_types)

	transactions := &[]models.Transaction{}
	db = db.Find(transactions)

	return transactions, db.Error
}

func (m *Crud[Model, ModelOrm]) SelectBatchOrder(
	limit int,
	skip int,
	order string,
) (*[]Model, error) {
	db := m.db
	db = db.Model(&m.Model)
	db = db.Limit(limit)
	if skip != 0 {
		db = db.Offset(skip)
	}
	db = db.Order(order)
	output := &[]Model{}
	db = db.Find(output)
	return output, db.Error
}
