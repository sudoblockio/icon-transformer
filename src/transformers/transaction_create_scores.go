package transformers

import (
	"encoding/json"
	"errors"
	"github.com/sudoblockio/icon-transformer/models"
	"github.com/sudoblockio/icon-transformer/service"
	"go.uber.org/zap"
)

func getScoreAddressFromTransactionResult(result map[string]interface{}) (string, error) {
	scoreAddress, ok := result["scoreAddress"].(string)
	if ok == false {
		return "", errors.New("Invalid response")
	}

	return scoreAddress, nil
}

// This is for python scores where the Tx for approval needs to be extracted out of the Tx hash by calling the Tx result
//  manually. This is because the score address is not the one actually creating the Tx.
func transformBlockETLToTransactionByAddressCreateScores(blockETL *models.BlockETL) []*models.TransactionByAddress {
	transactionByAddresses := []*models.TransactionByAddress{}

	for _, transactionETL := range blockETL.Transactions {
		if transactionETL.Status == "0x0" {
			// Skip failed Txs which could have unaccepted results
			continue
		}

		method := extractMethodFromTransactionETL(transactionETL)

		if method == "acceptScore" || method == "rejectScore" {
			result, err := service.IconNodeServiceGetTransactionResult(transactionETL.Hash)
			if err != nil {
				zap.S().Warn("Could not get tx result for hash: ", err.Error(), ",Hash=", transactionETL.Hash)
				continue
			}

			scoreAddress, err := getScoreAddressFromTransactionResult(result)
			if err != nil {
				zap.S().Warn("Could not get score address from creation tx hash: ", err.Error(), ",Hash=", transactionETL.Hash)
				continue
			}

			transactionByAddress := &models.TransactionByAddress{
				TransactionHash: transactionETL.Hash,
				Address:         scoreAddress,
				BlockNumber:     blockETL.Number,
			}

			transactionByAddresses = append(transactionByAddresses, transactionByAddress)
		} else if method == "" {
			// Here we need to perform additional logic to get contract creation events. These events have been
			// inconsistent over time where in some cases the to-address is the contract address and in others it is
			// the cx0 (cx0000000000000000000000000000000000000000) address.
			// To fix this, we are going to manually extract the scoreAddress by making a getTransactionResult call
			// which will return the score address and persist this Tx in the transactions_by_address table which is
			// used to get the transaction list tables.

			// TODO: Need to classify transaction type
			// Relates to https://github.com/sudoblockio/icon-tracker-frontend/issues/19

			dataJSON := map[string]interface{}{}
			err := json.Unmarshal([]byte(transactionETL.Data), &dataJSON)
			if err != nil {
				//zap.S().Warn("Transaction data field parsing error: ", err.Error(), ",Hash=", transactionETL.Hash)
				continue
			}
			content, ok := dataJSON["contentType"]
			// contentType is in tx data so it must be a contract creation event
			if ok && (content == "application/zip" || content == "application/java") {
				result, err := service.IconNodeServiceGetTransactionResult(transactionETL.Hash)
				if err != nil {
					zap.S().Warn("Could not get tx result for hash: ", err.Error(), ",Hash=", transactionETL.Hash)
					continue
				}

				scoreAddress, err := getScoreAddressFromTransactionResult(result)
				if err != nil {
					zap.S().Warn("Could not get score address from creation tx hash: ", err.Error(), ",Hash=", transactionETL.Hash)
					continue
				}

				transactionByAddress := &models.TransactionByAddress{
					TransactionHash: transactionETL.Hash,
					Address:         scoreAddress,
					BlockNumber:     blockETL.Number,
				}

				transactionByAddresses = append(transactionByAddresses, transactionByAddress)
			}
		}
	}
	return transactionByAddresses
}

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

//func transformBlockETLToTransactionByAddressCreateScores(blockETL *models.BlockETL) []*models.TransactionByAddress {
//	transactionByAddresses := []*models.TransactionByAddress{}
//
//	for _, transactionETL := range blockETL.Transactions {
//		method := extractMethodFromTransactionETL(transactionETL)
//
//		if method == "acceptScore" || method == "rejectScore" {
//			// Creation Transaction Hash
//			creationTransactionHash := ""
//			if transactionETL.Data != "" {
//				dataJSON := map[string]interface{}{}
//				err := json.Unmarshal([]byte(transactionETL.Data), &dataJSON)
//				if err == nil {
//
//					paramsInterface, ok := dataJSON["params"]
//					if ok {
//						// Params field is in dataJSON
//						params := paramsInterface.(map[string]interface{})
//
//						creationTransactionHashInterface, ok := params["txHash"]
//						if ok {
//							// Parsing successful
//							creationTransactionHash = creationTransactionHashInterface.(string)
//						}
//					}
//					if ok == false {
//						// Parsing error
//						zap.S().Warn("Transaction params field parsing error: ", err.Error(), ",Hash=", transactionETL.Hash)
//						continue
//					}
//				} else {
//					// Parsing error
//					zap.S().Warn("Transaction data field parsing error: ", err.Error(), ",Hash=", transactionETL.Hash)
//					continue
//				}
//			}
//
//			result, err := service.IconNodeServiceGetTransactionResult(creationTransactionHash)
//			if err != nil {
//				zap.S().Warn("Could not get score address from creation tx hash: ", err.Error(), ",Hash=", creationTransactionHash)
//				continue
//			}
//
//			getScoreAddressFromTransactionResult(result)
//
//			transactionByAddress := &models.TransactionByAddress{
//				TransactionHash: transactionETL.Hash,
//				Address:         scoreAddress,
//				BlockNumber:     blockETL.Number,
//			}
//
//			transactionByAddresses = append(transactionByAddresses, transactionByAddress)
//		}
//	}
//	return transactionByAddresses
//}
