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

	crud.TransactionTypeCrud.LoaderChannel <- transaction
}

func updateTransactionTypes(
	blockETL *models.BlockETL,
	transactionETL *models.TransactionETL,
	transactionTypeCreationEvents *[]models.Transaction,
	scoreAddress string,
	newStatus int32,
	oldStatus int32,
) {
	// Filter out Txs without some fields as some erroneous data has made it into DB
	// Could only be a couple records but still should be run for time being as not sure how many exactly.
	// TODO: RM this in later sync
	newTxTypeList := []models.Transaction{}
	for _, v := range *transactionTypeCreationEvents {
		if v.Hash[2:] == "0x" && v.ToAddress != "" {
			newTxTypeList = append(newTxTypeList, v)
		}
	}

	// Classify Tx based on history of other Txs
	if len(newTxTypeList) == 0 || blockETL.Number == (newTxTypeList)[0].BlockNumber {
		upsertTransactionType(transactionETL.Hash, scoreAddress, newStatus)
	} else if blockETL.Number < (newTxTypeList)[0].BlockNumber {
		upsertTransactionType(transactionETL.Hash, scoreAddress, newStatus)
		for _, v := range newTxTypeList {
			upsertTransactionType(v.Hash, scoreAddress, oldStatus)
		}
	} else {
		upsertTransactionType(transactionETL.Hash, scoreAddress, oldStatus)
	}

	//if len(*transactionTypeCreationEvents) == 0 || blockETL.Number == (*transactionTypeCreationEvents)[0].BlockNumber {
	//	upsertTransactionType(blockETL.Hash, scoreAddress, newStatus)
	//} else if blockETL.Number < (*transactionTypeCreationEvents)[0].BlockNumber {
	//	upsertTransactionType(blockETL.Hash, scoreAddress, newStatus)
	//	for _, v := range *transactionTypeCreationEvents {
	//		upsertTransactionType(v.Hash, scoreAddress, oldStatus)
	//	}
	//} else {
	//	upsertTransactionType(blockETL.Hash, scoreAddress, oldStatus)
	//}
}

func getTransactionTypes(scoreAddress string, transactionTypes []int32) *[]models.Transaction {
	transactions, err := crud.GetTransactionCrud().SelectManyContractCreations(scoreAddress, transactionTypes)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return transactions
}

// This is for python scores where the Tx for approval needs to be extracted out of the Tx hash by calling the Tx result
//
//	manually. This is because the score address is not the one actually creating the Tx.
func transactionByAddressCreateScores(blockETL *models.BlockETL) {
	//transactionByAddresses := []*models.TransactionByAddress{}

	for _, transactionETL := range blockETL.Transactions {
		method := extractMethodFromTransactionETL(transactionETL)

		if method == "acceptScore" || method == "rejectScore" {

			if transactionETL.Status != "0x1" {
				// Failed Txs will not have event logs that can be indexed
				continue
			}
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

			crud.TransactionByAddressCreateScoreCrud.LoaderChannel <- transactionByAddress
			transactionTypeCreateScores := getTransactionTypes(scoreAddress, []int32{5, 6, 7, 8, 9})

			if method == "acceptScore" {
				updateTransactionTypes(
					blockETL,
					transactionETL,
					transactionTypeCreateScores,
					scoreAddress,
					5,
					7,
				)
			} else if method == "rejectScore" {
				updateTransactionTypes(
					blockETL,
					transactionETL,
					transactionTypeCreateScores,
					scoreAddress,
					6,
					8,
				)
			}
		} else if method == "" {
			// Here we need to perform additional logic to get contract creation events. These events have been
			// inconsistent over time where in some cases the to-address is the contract address and in others it is
			// the cx0 (cx0000000000000000000000000000000000000000) address.
			// To fix this, we are going to manually extract the scoreAddress by making a getTransactionResult call
			// which will return the score address and persist this Tx in the transactions_by_address table which is
			// used to get the transaction list tables.
			content, ok := extractContentFromTranactionETL(transactionETL)

			if ok && (content == "application/zip" || content == "application/java") {
				if transactionETL.Status != "0x1" {
					//skip failed Txs as they won't have responses from API
					continue
				}

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

				//transactionByAddresses = append(transactionByAddresses, transactionByAddress)

				crud.TransactionByAddressCreateScoreCrud.LoaderChannel <- transactionByAddress
				transactionTypeCreateScores := getTransactionTypes(scoreAddress, []int32{3, 4})

				updateTransactionTypes(
					blockETL,
					transactionETL,
					transactionTypeCreateScores,
					scoreAddress,
					3,
					4,
				)
			}
		}
	}
}
