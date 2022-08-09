package transformers

import (
	"encoding/json"
	"github.com/sudoblockio/icon-transformer/models"
	"reflect"
	"runtime"
	"strings"
)

func extractMethodFromTransactionETL(transactionETL *models.TransactionETL) string {

	method := ""
	if transactionETL.Data != "" {
		dataJSON := map[string]interface{}{}
		err := json.Unmarshal([]byte(transactionETL.Data), &dataJSON)
		if err == nil {
			// Parsing successful
			if methodInterface, ok := dataJSON["method"]; ok {
				// Method field is in dataJSON
				method = methodInterface.(string)
			}
		} else {
			// Parsing error
			return ""
		}
	}

	return method
}

func extractMethodFromLogETL(logETL *models.LogETL) string {
	return strings.Split(logETL.Indexed[0], "(")[0]
}

func extractContentFromTranactionETL(transactionETL *models.TransactionETL) (string, bool) {
	dataJSON := map[string]interface{}{}
	err := json.Unmarshal([]byte(transactionETL.Data), &dataJSON)
	if err != nil {
		//zap.S().Warn("Transaction data field parsing error: ", err.Error(), ",Hash=", transactionETL.Hash)
		return "", false
	}
	content, ok := dataJSON["contentType"].(string)

	return content, ok
}

// Used to get the function name so that we can put that into a config file so that only parts of the trnasformation
// can be processed.
func getFunctionName(i interface{}) string {
	functionParts := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	functionNames := strings.Split(functionParts, ".")
	return functionNames[len(functionNames)-1]
}
