package transformers

import (
	"errors"
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/models"
	"github.com/sudoblockio/icon-transformer/service"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func getScoreAddressFromTransactionResult(result map[string]interface{}) (string, error) {
	scoreAddress, ok := result["scoreAddress"].(string)
	if ok == false {
		return "", errors.New("Invalid response")
	}

	return scoreAddress, nil
}

func upsertTransactionType(transactionHash string, scoreAddress string, transactionType int32) {
	transaction := &models.Transaction{
		Hash:            transactionHash,
		TransactionType: transactionType,
		ScoreAddress:    scoreAddress,
		LogIndex:        -1,
	}

	crud.GetTransactionCrud().UpsertOneCols(transaction, []string{"hash", "transaction_type", "score_address"})
}

func updateTransactionTypes(
	blockETL *models.BlockETL,
	transactionTypeCreationEvents *[]models.Transaction,
	scoreAddress string,
	newStatus int32,
	oldStatus int32,
) {
	if len(*transactionTypeCreationEvents) == 0 || blockETL.Number == (*transactionTypeCreationEvents)[0].BlockNumber {
		upsertTransactionType(blockETL.Hash, scoreAddress, newStatus)
	} else if blockETL.Number < (*transactionTypeCreationEvents)[0].BlockNumber {
		upsertTransactionType(blockETL.Hash, scoreAddress, newStatus)
		for _, v := range *transactionTypeCreationEvents {
			upsertTransactionType(v.Hash, scoreAddress, oldStatus)
		}
	} else {
		upsertTransactionType(blockETL.Hash, scoreAddress, oldStatus)
	}
}

func getTransactionTypes(scoreAddress string, transactionTypes []int32) *[]models.Transaction {
	transactions, err := crud.GetTransactionCrud().SelectManyContractCreations(scoreAddress, transactionTypes)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return transactions
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

			creationHash := transactionETL.Logs[0].Indexed[1]

			result, err := service.IconNodeServiceGetTransactionResult(creationHash)
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

			transactionTypeCreateScores := getTransactionTypes(scoreAddress, []int32{5, 6, 7, 8, 9})

			if method == "acceptScore" {
				updateTransactionTypes(blockETL, transactionTypeCreateScores, scoreAddress, 5, 7)
			} else if method == "rejectScore" {
				updateTransactionTypes(blockETL, transactionTypeCreateScores, scoreAddress, 6, 8)
			}
		} else if method == "" {
			// Here we need to perform additional logic to get contract creation events. These events have been
			// inconsistent over time where in some cases the to-address is the contract address and in others it is
			// the cx0 (cx0000000000000000000000000000000000000000) address.
			// To fix this, we are going to manually extract the scoreAddress by making a getTransactionResult call
			// which will return the score address and persist this Tx in the transactions_by_address table which is
			// used to get the transaction list tables.

			// TODO: Need to classify transaction type
			// Relates to https://github.com/sudoblockio/icon-tracker-frontend/issues/19

			content, ok := extractContentFromTranactionETL(transactionETL)

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

				transactionTypeCreateScores := getTransactionTypes(scoreAddress, []int32{3, 4})

				updateTransactionTypes(blockETL, transactionTypeCreateScores, scoreAddress, 3, 4)
			}
		}
	}
	return transactionByAddresses
}
