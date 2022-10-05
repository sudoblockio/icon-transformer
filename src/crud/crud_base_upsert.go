package crud

import "gorm.io/gorm/clause"

func (m *Crud[Model, ModelOrm]) UpsertManyColumns(
	values []*Model,
	cols []string,
) error {
	db := m.db
	db = db.Model(&m.Model)
	db = db.Clauses(clause.OnConflict{
		Columns:   m.primaryKeys,
		DoUpdates: clause.AssignmentColumns(cols),
	}).Create(values)

	if db.Error == nil {
		return nil
	}
	return db.Error
}

func (m *Crud[Model, ModelOrm]) UpsertMany(
	values []*Model,
) error {
	return m.UpsertManyColumns(values, m.columns)
}

func (m *Crud[Model, ModelOrm]) UpsertOneColumns(
	value *Model,
	cols []string,
) error {
	db := m.db
	db = db.Clauses(clause.OnConflict{
		Columns:   m.primaryKeys,
		DoUpdates: clause.AssignmentColumns(cols),
	}).Create(value)
	return db.Error
}

func (m *Crud[Model, ModelOrm]) UpsertOne(
	value *Model,
) error {
	return m.UpsertOneColumns(value, m.columns)
}

func (m *Crud[Model, ModelOrm]) LoopUpsertOne(values []*Model) error {
	for _, v := range values {
		db := m.db
		db = db.Clauses(clause.OnConflict{
			Columns:   m.primaryKeys,
			DoUpdates: clause.AssignmentColumns(m.columns),
		}).Create(v)

		if db.Error != nil {
			return db.Error
		}
	}
	return nil
}
