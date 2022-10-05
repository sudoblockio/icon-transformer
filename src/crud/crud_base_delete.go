package crud

import "github.com/sudoblockio/icon-transformer/models"

// DeleteMissing - WIP -> used in missing blocks
func (m *Crud[Model, ModelOrm]) DeleteMissing() error {
	db := m.db
	db = db.Model(&m.Model)

	// NOTE delete needs a WHERE clause
	db = db.Where("number > 0")

	// Delete
	db = db.Delete(&models.MissingBlock{})

	return db.Error
}
