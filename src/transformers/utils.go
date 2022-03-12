package transformers

import (
	"encoding/json"
	"strings"

	"github.com/sudoblockio/icon-go-worker/models"
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
