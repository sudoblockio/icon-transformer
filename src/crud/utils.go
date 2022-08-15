package crud

import (
	"encoding/json"
	"fmt"
	"github.com/sudoblockio/icon-transformer/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// TODO: Replace when done
// retryLoader - retry a function until it returns a non-nil error
func retryLoader[T any](input T, f func(i T) error, attempts int, sleep time.Duration) (err error) {
	for i := 0; i < attempts; i++ {
		if i > 0 {
			log.Println("retrying after error:", err)
			time.Sleep(sleep)
			sleep *= 2
		}
		err = f(input)
		if err == nil {
			return nil
		}
	}
	return fmt.Errorf("after %d attempts, last error: %s", attempts, err)
}

// TODO: This is going to be replaced
// retryLoader - retry a function until it returns a non-nil error
func retryCrudColumns[T any](input T, columns []string, f func(i T, c []string) error, attempts int, sleep time.Duration) (err error) {
	for i := 0; i < attempts; i++ {
		if i > 0 {
			log.Println("retrying after error:", err)
			time.Sleep(sleep)
			sleep *= 2
		}
		err = f(input, columns)
		if err == nil {
			return nil
		}
	}
	return fmt.Errorf("after %d attempts, last error: %s", attempts, err)
}

// getModelColumnNames - Get a slice of column names from a model's tags
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

// TODO: Replace with parser from model tags that gets the actual column name instead of just assuming it is snake
//
//	https://github.com/sudoblockio/icon-transformer/issues/40
var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

// getModelPrimaryKeys - Get a slice of primary keys from the ORM model's tags
func getModelPrimaryKeys[T any](model T) []clause.Column {
	var fields []clause.Column
	vals := reflect.ValueOf(model)
	for i, _ := range reflect.VisibleFields(vals.Type()) {
		gormField, _ := vals.Type().Field(i).Tag.Lookup("gorm")
		if gormField == "primary_key" {
			// Assuming keys are in snake case -> Should probably do another lookup
			fields = append(fields, clause.Column{Name: ToSnakeCase(vals.Type().Field(i).Name)})
		}
	}
	return fields
}

// TODO: See below - https://github.com/sudoblockio/icon-transformer/issues/38
// retryBatchLoader - retry a function until it returns a non-nil error
func retryBatchLoader[T any](input []T, f func([]T, []string) error, cols []string, attempts int, sleep time.Duration) (err error) {
	for i := 0; i < attempts; i++ {
		if i > 0 {
			log.Println("retrying after error:", err)
			time.Sleep(sleep)
			sleep *= 2
		}
		err = f(input, cols)
		if err == nil {
			return nil
		}
	}
	return fmt.Errorf("after %d attempts, last error: %s", attempts, err)
}

// TODO: This will be replaced / updated when crud functions are implemented
//
//	https://github.com/sudoblockio/icon-transformer/issues/38
type GenericChan[T any] chan T

// startBatchLoader - Starts a loader channel that batches the inputs to the
func startBatchLoader[T any](c GenericChan[T], f func([]T, []string) error, cols []string) {
	var channelOutputs []T
	go func() {
		for {
			for i := 1; i < len(c); i++ {
				channelOutputs = append(channelOutputs, <-c)
			}
			if len(channelOutputs) == 0 {
				time.Sleep(100 * time.Millisecond)
				continue
			}

			fmt.Println(len(channelOutputs))
			// Retry on failure
			err := retryBatchLoader(
				channelOutputs,
				f,
				cols,
				5,
				config.Config.DbRetrySleep,
			)
			// Clear the channel
			channelOutputs = nil
			if err != nil {
				// Postgres error
				zap.S().Fatal(err.Error())
			}
		}
	}()
}

type GormErr struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
}

// getGormError - Unmarshall the go error
func getGormError(db *gorm.DB) GormErr {
	var newError GormErr
	if db.Error != nil {
		byteErr, _ := json.Marshal(db.Error)
		json.Unmarshal(byteErr, &newError)
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
