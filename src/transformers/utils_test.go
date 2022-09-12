package transformers

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

import (
	"encoding/json"
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/models"
	"io"
	"os"
	"path/filepath"
)

type TestCase[T any] struct {
	blockFile string
	Expected  *T
}

func testReadBlock[T any](t *testing.T, c TestCase[T]) *models.BlockETL {
	jsonFile, err := os.Open(filepath.Join("testdata", c.blockFile))
	assert.Nil(t, err)

	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		jsonFile, err = os.Open(filepath.Join("testdata", c.blockFile))
		assert.Nil(t, err)
	}(jsonFile)

	jsonTransaction, _ := io.ReadAll(jsonFile)

	var block *models.BlockETL
	err = json.Unmarshal(jsonTransaction, &block)
	return block
}

func testTransformerEquals[M any, _ any](
	block *models.BlockETL,
	Crud *crud.Crud[M, crud.ModelOrm],
	model M,
	modelOrm crud.ModelOrm,
	transformer func(etl *models.BlockETL),
) *M {
	Crud = crud.GetCrud(model, modelOrm)
	Crud.LoaderChannel = make(chan *M, 2)
	transformer(block)
	return <-Crud.LoaderChannel
}

func testCompareStructValues(t *testing.T, expectedStruct interface{}, actualStruct interface{}) {
	expectedValue := reflect.Indirect(reflect.ValueOf(expectedStruct))
	outputValue := reflect.Indirect(reflect.ValueOf(actualStruct))

	expectedFields := reflect.VisibleFields(expectedValue.Type())

	for i := 0; i < expectedValue.NumField(); i++ {

		if !expectedFields[i].IsExported() {
			continue
		}

		expected := expectedValue.Field(i).Interface()
		if expected == nil || expected == "" {
			continue
		}
		actualName := expectedValue.Type().Field(i).Name
		actual := outputValue.FieldByName(actualName).Interface()

		assert.Equal(t, expected, actual, actualName)
	}
}

func TestUtilsGetFunctionName(t *testing.T) {
	function := getFunctionName(getFunctionName)

	assert.Equal(t, function, "getFunctionName")
}
