package crud

import "github.com/sudoblockio/icon-transformer/models"

// DeleteMissing - Delete records with a string where condition
func (m *Crud[Model, ModelOrm]) DeleteMissing(where string) error {
	db := m.db
	db = db.Model(&m.Model)
	// NOTE delete needs a WHERE clause
	db = db.Where(where)
	db = db.Delete(&models.MissingBlock{})
	return db.Error
}
