package crud

import (
	"github.com/stretchr/testify/assert"
	"github.com/sudoblockio/icon-transformer/models"
	"gorm.io/gorm/clause"
	"reflect"
	"testing"
)

func TestCrudUtilsExtractFilledFieldsFromModel(t *testing.T) {
	model := models.Address{Address: "foo"}
	fields := extractFilledFieldsFromModel(
		reflect.ValueOf(model),
		reflect.TypeOf(model),
	)
	assert.Equal(t, fields, map[string]interface{}{"address": "foo", "is_contract": false, "is_prep": false, "is_token": false})
}

func TestCrudUtilsGetModelColumnsNames(t *testing.T) {
	model := models.Address{}
	keys := getModelColumnNames(model)
	assert.Greater(t, len(keys), 10)

	for i := range keys {
		// Make sure it does not have any of the default proto struct fields
		assert.NotEqual(t, "sizeOf", keys[i])
	}
}

func TestCrudUtilsGetModelPrimaryKeys(t *testing.T) {
	model := models.AddressORM{}
	keys := getModelPrimaryKeys(model)
	assert.Equal(t, keys[0], clause.Column{Name: "address"})
	for i := range keys {
		// Make sure it does not have any of the default proto struct fields
		assert.NotEqual(t, "sizeOf", keys[i])
	}
}
