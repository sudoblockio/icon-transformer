package crud

import (
	"fmt"
	"reflect"
	"strconv"
)

func removeDuplicatePrimaryKeys[m any](models []*m, primaryKeys []string) []*m {
	var key string
	unique := make(map[string]bool)
	var result []*m

	for _, v := range models {
		for _, pk := range primaryKeys {
			r := reflect.ValueOf(*v).FieldByName(pk)

			switch r.Type().String() {
			default:
				error.Error(fmt.Errorf("Unknown type of field %s", key))
			case "string":
				key += r.String()
			case "int64", "int", "int32":
				key += strconv.Itoa(int(r.Int()))
			}
		}

		if _, ok := unique[key]; !ok {
			unique[key] = true
			result = append(result, v)
		}
		key = ""
	}
	return result
}
