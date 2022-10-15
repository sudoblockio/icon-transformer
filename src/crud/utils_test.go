package crud

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/sudoblockio/icon-transformer/models"
	"gorm.io/gorm/clause"
	"reflect"
	"regexp"
	"testing"
)

func TestCrudUtilsExtractFilledFieldsFromModel(t *testing.T) {
	model := models.Address{Address: "foo"}
	fields := extractFilledFieldsFromModel(
		reflect.ValueOf(model),
		reflect.TypeOf(model),
	)
	assert.Equal(t, fields["address,omitempty"], "foo")
	//assert.Equal(t, fields, map[string]interface{}{"address": "foo", "is_contract": false, "is_prep": false, "is_token": false})
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
	model := models.LogORM{}
	keys := getModelPrimaryKeys(model)
	assert.Equal(t, keys[0], clause.Column{Name: "log_index"})
	for i := range keys {
		// Make sure it does not have any of the default proto struct fields
		assert.NotEqual(t, "sizeOf", keys[i])
	}
}

type TestsRegex struct {
	regexp *regexp.Regexp
	input  string
	output bool
}

func TestCrudUtilsRegex(t *testing.T) {
	for _, test := range []TestsRegex{
		{
			regexp: matchPrimaryKey,
			input:  "primary_key",
			output: true,
		},
		{
			regexp: matchPrimaryKey,
			input:  "primary_",
			output: false,
		},
		{
			regexp: matchPrimaryKey,
			input:  "primary_key;index:log_foo",
			output: true,
		},
	} {
		assert.Equal(
			t,
			test.regexp.MatchString(test.input),
			test.output,
			fmt.Sprintf("input: %v, regex: %v", test.input, test.regexp),
		)
	}
}
