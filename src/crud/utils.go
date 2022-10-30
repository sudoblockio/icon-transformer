package crud

import (
	"encoding/json"
	"fmt"
	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/metrics"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"
	"gorm.io/gorm/clause"
	"log"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// getModelColumnNames - Get a slice of column names from a Model's tags
func getModelColumnNames[T any](model T) []string {
	var fields []string
	vals := reflect.ValueOf(model)
	for i, f := range reflect.VisibleFields(vals.Type()) {
		if f.IsExported() {
			jsonField, _ := vals.Type().Field(i).Tag.Lookup("json")
			fields = append(fields, jsonField)
		}
	}
	return fields
}

func removeColumnNames(cols []string, rmCols []string) []string {
	var output []string
	for _, v := range cols {
		if !slices.Contains(rmCols, v) {
			output = append(output, v)
		}
	}
	return output
}

func (m *Crud[M, O]) removeColumnNames(rmCols []string) {
	var output []string
	for _, v := range m.columns {
		if !slices.Contains(rmCols, v) {
			output = append(output, v)
		}
	}
	m.columns = output
}

// TODO: Replace with parser from Model tags that gets the actual column Name instead of just assuming it is snake
//
//	https://github.com/sudoblockio/icon-transformer/issues/40
var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

var matchPrimaryKey = regexp.MustCompile("primary_key")

// getModelPrimaryKeys - Get a slice of primary keys from the ORM Model's tags
func getModelPrimaryKeys[T any](model T) []clause.Column {
	var fields []clause.Column
	vals := reflect.ValueOf(model)
	for i, _ := range reflect.VisibleFields(vals.Type()) {
		gormField, _ := vals.Type().Field(i).Tag.Lookup("gorm")
		//if gormField == "primary_key" {
		if matchPrimaryKey.MatchString(gormField) {
			// Assuming keys are in snake case -> Should probably do another lookup
			fields = append(fields, clause.Column{Name: ToSnakeCase(vals.Type().Field(i).Name)})
		}
	}
	return fields
}

// getModelPrimaryKeyFields - Get a slice of primary keys in pascal case from the ORM Model's tags
func getModelPrimaryKeyFields[T any](model T) []string {
	var fields []string
	vals := reflect.ValueOf(model)
	for i, _ := range reflect.VisibleFields(vals.Type()) {
		gormField, _ := vals.Type().Field(i).Tag.Lookup("gorm")
		//if gormField == "primary_key" {
		if matchPrimaryKey.MatchString(gormField) {
			// Assuming keys are in snake case -> Should probably do another lookup
			fields = append(fields, vals.Type().Field(i).Name)
		}
	}
	return fields
}

// columnToString - return slice of strings from slice of columns
func columnToString(cols []clause.Column) []string {
	var columnString []string
	for _, v := range cols {
		columnString = append(columnString, v.Name)
	}
	return columnString
}

// startCustomBatchUpsertLoader - upsert loader with customizations
func (m *Crud[M, O]) startBatchUpsertLoader() {

	batchCounter := m.batchSampleInterval
	var channelOutputs []*M

	go func() {
		for {
			time.Sleep(m.dbBufferWait)
			for i := 0; i < len(m.LoaderChannel); i++ {
				channelOutputs = append(channelOutputs, <-m.LoaderChannel)
			}

			// Sleep if queue is empty / at head
			numOutputs := len(channelOutputs)
			if numOutputs == 0 {
				time.Sleep(config.Config.DbIdleChannelWait)
				continue
			}

			channelOutputs = removeDuplicatePrimaryKeys[M](channelOutputs, m.primaryKeyFields)

			// Metrics
			batchCounter++
			if batchCounter > m.batchSampleInterval {
				//fmt.Println("Loader channel queue=", len(channelOutputs))
				m.metrics.loaderChannelLength.Set(float64(numOutputs))
				batchCounter = 0
			}

			// Retry on failure
			err := m.retryBatchLoader(
				channelOutputs,
				m.UpsertMany,
			)
			if err != nil {
				// Postgres error
				zap.S().Fatal(err.Error())
			}
			// Clear the channel
			channelOutputs = nil
		}
	}()
}

func (m *Crud[M, O]) createCrudMetrics() {
	// Metrics
	m.metrics.loaderChannelLength = metrics.CreateGuage(
		"loader_batch_size",
		"loader channel batch size",
		map[string]string{"table_name": m.TableName, "series": m.metrics.Name},
	)

	m.metrics.loaderChannelDuplicateErrors = metrics.CreateCounter(
		"loader_duplicate_errors",
		"number of duplicate msg errors sql 21000",
		map[string]string{"table_name": m.TableName, "series": m.metrics.Name},
	)

	m.metrics.loaderChannelDeadlockErrors = metrics.CreateCounter(
		"loader_deadlock_errors",
		"number of deadlock msg errors sql 40P01",
		map[string]string{"table_name": m.TableName, "series": m.metrics.Name},
	)
}

func (m *Crud[M, O]) MakeStartLoaderChannel() {
	m.LoaderChannel = make(chan *M, m.loaderChannelBuffer)
	m.createCrudMetrics()
	m.startBatchUpsertLoader()
}

// TODO: See below - https://github.com/sudoblockio/icon-transformer/issues/38
// retryBatchLoader - retry a function until it returns a non-nil error
func (m *Crud[M, O]) retryBatchLoader(
	input []*M,
	f func(values []*M) error,
) (err error) {
	var sleep = config.Config.DbRetrySleep

	for i := 0; i < m.retryAttempts; i++ {
		if i > 0 {
			log.Println("retrying after error:", err)
			time.Sleep(sleep)
			sleep *= 2
		}
		if m.batchErrorHandler == nil {
			err = m.DefaultRetryHandler(f(input), input)
		} else {
			err = m.batchErrorHandler(f(input), input)
		}
		if err == nil {
			return nil
		}

	}
	return fmt.Errorf("after %d attempts, last error: %s", m.retryAttempts, err)
}

type GormErr struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
}

// getGormError - Unmarshall the go error
func getGormError(err error) GormErr {
	var newError GormErr
	if err != nil {
		byteErr, ok := json.Marshal(err)
		if ok != nil {
			zap.S().Info(ok)
			return newError
		}
		ok = json.Unmarshal(byteErr, &newError)
		if ok != nil {
			zap.S().Info(ok)
			return newError
		}
		return newError
	}
	return newError
}

// TODO: Rm when batch is implemented -> Only for upsert one / very expensive to be done on each Tx
func extractFilledFieldsFromModel(modelValueOf reflect.Value, modelTypeOf reflect.Type) map[string]interface{} {
	// Helper for combining structs by giving a struct that represents what is in the DB and another struct that is
	//  partially instantiated.  Uses reflection to check if field is non-nill.  Purpose is to make sure you don't
	//  overwrite existing data with new data. Used everywhere an upsert is called.
	fields := map[string]interface{}{}

	for i := 0; i < modelValueOf.NumField(); i++ {
		modelField := modelValueOf.Field(i)
		modelType := modelTypeOf.Field(i)

		modelTypeJSONTag := modelType.Tag.Get("json")
		if modelTypeJSONTag != "" {
			// exported field

			// Check if field if filled
			modelFieldKind := modelField.Kind()
			isFieldFilled := true
			switch modelFieldKind {
			case reflect.String:
				v := modelField.Interface().(string)
				if v == "" {
					isFieldFilled = false
				}
			case reflect.Int:
				v := modelField.Interface().(int)
				if v == 0 {
					isFieldFilled = false
				}
			case reflect.Int8:
				v := modelField.Interface().(int8)
				if v == 0 {
					isFieldFilled = false
				}
			case reflect.Int16:
				v := modelField.Interface().(int16)
				if v == 0 {
					isFieldFilled = false
				}
			case reflect.Int32:
				v := modelField.Interface().(int32)
				if v == 0 {
					isFieldFilled = false
				}
			case reflect.Int64:
				v := modelField.Interface().(int64)
				if v == 0 {
					isFieldFilled = false
				}
			case reflect.Uint:
				v := modelField.Interface().(uint)
				if v == 0 {
					isFieldFilled = false
				}
			case reflect.Uint8:
				v := modelField.Interface().(uint8)
				if v == 0 {
					isFieldFilled = false
				}
			case reflect.Uint16:
				v := modelField.Interface().(uint16)
				if v == 0 {
					isFieldFilled = false
				}
			case reflect.Uint32:
				v := modelField.Interface().(uint32)
				if v == 0 {
					isFieldFilled = false
				}
			case reflect.Uint64:
				v := modelField.Interface().(uint64)
				if v == 0 {
					isFieldFilled = false
				}
			case reflect.Float32:
				v := modelField.Interface().(float32)
				if v == 0 {
					isFieldFilled = false
				}
			case reflect.Float64:
				v := modelField.Interface().(float64)
				if v == 0 {
					isFieldFilled = false
				}
			}

			if isFieldFilled == true {
				fields[modelTypeJSONTag] = modelField.Interface()
			}
		}
	}

	return fields
}
