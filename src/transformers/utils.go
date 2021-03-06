package transformers

import (
	"encoding/json"
	"strings"

	"github.com/sudoblockio/icon-transformer/models"
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
