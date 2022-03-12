package transformers

import (
	"encoding/json"

	"github.com/sudoblockio/icon-go-worker/models"
	"go.uber.org/zap"
)

func transformBlockETLToTransactionCreateScores(blockETL *models.BlockETL) []*models.TransactionCreateScore {

	transactionCreateScores := []*models.TransactionCreateScore{}

	for _, transactionETL := range blockETL.Transactions {
		method := extractMethodFromTransactionETL(transactionETL)

		if method == "acceptScore" || method == "rejectScore" {

			// Accept Transaction Hash
			acceptTransactionHash := ""
			if method == "acceptScore" {
				acceptTransactionHash = transactionETL.Hash
			}

			// Reject Transaction Hash
			rejectTransactionHash := ""
			if method == "rejectScore" {
				rejectTransactionHash = transactionETL.Hash
			}

			// Creation Transaction Hash
			creationTransactionHash := ""
			if transactionETL.Data != "" {
				dataJSON := map[string]interface{}{}
				err := json.Unmarshal([]byte(transactionETL.Data), &dataJSON)
				if err == nil {

					paramsInterface, ok := dataJSON["params"]
					if ok {
						// Params field is in dataJSON
						params := paramsInterface.(map[string]interface{})

						creationTransactionHashInterface, ok := params["txHash"]
						if ok {
							// Parsing successful
							creationTransactionHash = creationTransactionHashInterface.(string)
						}
					}
					if ok == false {
						// Parsing error
						zap.S().Warn("Transaction params field parsing error: ", err.Error(), ",Hash=", transactionETL.Hash)
						continue
					}
				} else {
					// Parsing error
					zap.S().Warn("Transaction data field parsing error: ", err.Error(), ",Hash=", transactionETL.Hash)
					continue
				}
			}

			transactionCreateScore := &models.TransactionCreateScore{
				CreationTransactionHash: creationTransactionHash,
				AcceptTransactionHash:   acceptTransactionHash,
				RejectTransactionHash:   rejectTransactionHash,
			}
			transactionCreateScores = append(transactionCreateScores, transactionCreateScore)
		}

	}

	return transactionCreateScores
}
