package crud

import "fmt"

// Count - count all entries in token_transfers table
// NOTE this function will take a long time
func (m *Crud[Model, ModelOrm]) Count() (int64, error) {
	db := m.db
	db = db.Model(&m.model)
	var count int64
	db = db.Count(&count)
	return count, db.Error
}

// CountByAddress - count all entries in token_transfers table
func (m *Crud[Model, ModelOrm]) CountByAddress(address string) (int64, error) {
	db := m.db
	db = db.Model(&m.model)
	db = db.Where("to_address = ? or from_address = ?", address, address)
	var count int64
	db = db.Count(&count)
	return count, db.Error
}

// Count for regular transactions.
// NOTE this function will take a long time
func (m *Crud[Model, ModelOrm]) CountWhere(field string, equals string) (int64, error) {
	db := m.db
	//db = db.Model(&m.model).Where("? = ?", field, equals)
	db = db.Model(&m.model).Where(fmt.Sprintf("%s = ?", field), equals)
	var count int64
	db = db.Count(&count)
	return count, db.Error
}

// CountByTokenContractAddress - select many from addreses table
func (m *Crud[Model, ModelOrm]) CountByTokenContractAddress() (map[string]int64, error) {
	db := m.db

	// Set table
	db = db.Model(&m.model)

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

// Count all entries in token_transfers table by token contract
func (m *Crud[Model, ModelOrm]) CountByTokenContract(address string) (int64, error) {
	db := m.db
	db = db.Model(&m.model)
	db = db.Where("token_contract_address = ?", address)
	var count int64
	db = db.Count(&count)
	return count, db.Error
}

// Count for regular transactions.
// NOTE this function will take a long time
func (m *Crud[Model, ModelOrm]) CountTransactionsRegular() (int64, error) {
	db := m.db
	db = db.Model(&m.model).Where("type = 'transaction'")
	var count int64
	db = db.Count(&count)
	return count, db.Error
}

// Count for internal transactions.
// NOTE this function will take a long time
func (m *Crud[Model, ModelOrm]) CountTransactionsInternal() (int64, error) {
	db := m.db
	db = db.Model(&m.model).Where("type = 'log'")
	var count int64
	db = db.Count(&count)
	return count, db.Error
}

// Count for internal transactions.
// NOTE this function will take a long time
func (m *Crud[Model, ModelOrm]) CountTransactionsInternalByAddress(address string) (int64, error) {
	db := m.db
	db = db.Model(&m.model).Where("type = 'log'")
	db = db.Model(&m.model).Where("to_address = ? or from_address = ?", address, address)
	var count int64
	db = db.Count(&count)
	return count, db.Error
}

// Count for regular transactions.
// NOTE this function will take a long time
func (m *Crud[Model, ModelOrm]) CountTransactionsRegularByAddress(address string) (int64, error) {
	db := m.db
	db = db.Model(&m.model).Where("type='transaction'")
	db = db.Model(&m.model).Where("to_address = ? or from_address = ?", address, address)
	var count int64
	db = db.Count(&count)
	return count, db.Error
}

// TODO: RM?
// Count for regular transactions.
// NOTE this function will take a long time
func (m *Crud[Model, ModelOrm]) CountTransactionIcxByAddress(address string) (int64, error) {
	db := m.db
	db = db.Model(&m.model).Where("type='transaction'")
	db = db.Where("value_decimal != 0")
	db = db.Model(&m.model).Where("to_address = ? or from_address = ?", address, address)
	var count int64
	db = db.Count(&count)
	return count, db.Error
}
