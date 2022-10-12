package crud

import (
	"reflect"
)

func removeDuplicatePrimaryKeys[m any](models []*m, primaryKeys []string) []*m {
	var key string
	unique := make(map[string]bool)
	var result []*m

	for _, v := range models {
		for _, pk := range primaryKeys {
			key += reflect.ValueOf(*v).FieldByName(pk).String()
		}

		if _, ok := unique[key]; !ok {
			unique[key] = true
			result = append(result, v)
		}
		key = ""
	}
	return result
}
