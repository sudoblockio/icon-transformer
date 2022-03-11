package transformers

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"

	"github.com/sudoblockio/icon-go-worker/config"
	"github.com/sudoblockio/icon-go-worker/crud"
	"github.com/sudoblockio/icon-go-worker/kafka"
	"github.com/sudoblockio/icon-go-worker/models"
)

func Start() {
	go start()
}

func start() {
	kafkaBlocksTopic := config.Config.KafkaBlocksTopic

	// Input channels
	kafkaBlocksTopicChannel := kafka.KafkaTopicConsumer.TopicChannels[kafkaBlocksTopic]

	// Output channels
	blockLoaderChannel := crud.GetBlockCrud().LoaderChannel
	transactionLoaderChannel := crud.GetTransactionCrud().LoaderChannel

	zap.S().Debug("Blocks transformer: started working")
	for {

		///////////////////
		// Kafka Message //
		///////////////////

		consumerTopicMsg := <-kafkaBlocksTopicChannel
		blockETL, err := convertToBlockETLProtoBuf(consumerTopicMsg.Value)
		if err != nil {
			zap.S().Fatal("transformer: Unable to proceed cannot convert kafka msg value to BlockETL, err: ", err.Error())
		}
		zap.S().Info("Transformer: Processing block #", blockETL.Number)

		/////////////
		// Loaders //
		/////////////

		// Block Loader
		go func() {
			block := transformBlockETLToBlock(blockETL)
			blockLoaderChannel <- block
		}()

		// Transaction Loader
		go func() {
			transactions := transformBlockETLToTransactions(blockETL)
			for _, transaction := range transactions {
				transactionLoaderChannel <- transaction
			}
		}()
	}
}

func convertToBlockETLProtoBuf(value []byte) (*models.BlockETL, error) {
	block := &models.BlockETL{}
	err := proto.Unmarshal(value, block)
	return block, err
}

func transformBlockETLToBlock(blockETL *models.BlockETL) *models.Block {

	//////////////////
	// Transactions //
	//////////////////
	transactionCount := int64(len(blockETL.Transactions))
	transactionAmount := "0x0"
	transactionFees := "0x0"

	sumTransactionAmountBig := big.NewInt(0)
	sumTransactionFeesBig := big.NewInt(0)
	for _, transactionETL := range blockETL.Transactions {

		// transactionAmount
		transactionAmountBig := big.NewInt(0)
		if transactionETL.Value != "" {
			transactionAmountBig.SetString(transactionETL.Value[2:], 16)
		}
		sumTransactionAmountBig = sumTransactionAmountBig.Add(sumTransactionAmountBig, transactionAmountBig)

		// transactionFees
		stepPriceBig := big.NewInt(0)
		if transactionETL.StepPrice != "" {
			stepPriceBig.SetString(transactionETL.StepPrice[2:], 16)
		}
		stepUsedBig := big.NewInt(0)
		if transactionETL.StepUsed != "" {
			stepUsedBig.SetString(transactionETL.StepUsed[2:], 16)
		}
		transactionFeeBig := stepUsedBig.Mul(stepUsedBig, stepPriceBig)
		sumTransactionFeesBig = sumTransactionFeesBig.Add(sumTransactionFeesBig, transactionFeeBig)
	}
	transactionAmount = fmt.Sprintf("0x%x", sumTransactionAmountBig)
	transactionFees = fmt.Sprintf("0x%x", sumTransactionFeesBig)

	/////////////////////////
	// Failed Transactions //
	/////////////////////////
	failedTransactionCount := int64(0)
	for _, transactionETL := range blockETL.Transactions {

		if transactionETL.Status == "0x0" {
			// failedTransactionCount
			failedTransactionCount++
		}
	}

	//////////
	// Logs //
	//////////
	logCount := int64(0)
	for _, transactionETL := range blockETL.Transactions {
		logCount += int64(len(transactionETL.Logs))
	}

	///////////////////////////
	// Internal Transactions //
	///////////////////////////
	internalTransactionCount := int64(0)
	internalTransactionAmount := "0x0"

	sumInternalTransactionAmountBig := big.NewInt(0)
	for _, transactionETL := range blockETL.Transactions {
		for _, log := range transactionETL.Logs {
			method := strings.Split(log.Indexed[0], "(")[0]

			if method == "ICXTransfer" {
				// internalTransactionCount
				internalTransactionCount++

				// internalTransactionAmount
				internalTransactionAmountBig := big.NewInt(0)
				internalTransactionAmountBig.SetString(log.Indexed[3][2:], 16)
				sumInternalTransactionAmountBig = sumInternalTransactionAmountBig.Add(sumInternalTransactionAmountBig, internalTransactionAmountBig)
			}
		}
	}
	internalTransactionAmount = fmt.Sprintf("0x%x", sumInternalTransactionAmountBig)

	return &models.Block{
		Number:                    blockETL.Number,
		PeerId:                    blockETL.PeerId,
		Signature:                 blockETL.Signature,
		Version:                   blockETL.Version,
		MerkleRootHash:            blockETL.MerkleRootHash,
		Hash:                      blockETL.Hash,
		ParentHash:                blockETL.ParentHash,
		Timestamp:                 blockETL.Timestamp,
		TransactionCount:          transactionCount,
		LogCount:                  logCount,
		TransactionAmount:         transactionAmount,
		TransactionFees:           transactionFees,
		FailedTransactionCount:    failedTransactionCount,
		InternalTransactionCount:  internalTransactionCount,
		InternalTransactionAmount: internalTransactionAmount,
	}
}

func transformBlockETLToTransactions(blockETL *models.BlockETL) []*models.Transaction {

	transactions := []*models.Transaction{}

	//////////////////
	// Transactions //
	//////////////////
	for _, transactionETL := range blockETL.Transactions {

		// Method
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
				zap.S().Warn("Transaction data field parsing error: ", err.Error(), ",Hash=", transactionETL.Hash)
			}
		}

		// Value Decimal
		valueDecimal := float64(0)
		if transactionETL.Value != "" {
			valueDecimal = stringHexToFloat64(transactionETL.Value, 18)
		}

		// Transaction Fee
		stepPriceBig := big.NewInt(0)
		if transactionETL.StepPrice != "" {
			stepPriceBig.SetString(transactionETL.StepPrice[2:], 16)
		}
		stepUsedBig := big.NewInt(0)
		if transactionETL.StepUsed != "" {
			stepUsedBig.SetString(transactionETL.StepUsed[2:], 16)
		}
		transactionFeeBig := stepUsedBig.Mul(stepUsedBig, stepPriceBig)
		transactionFee := fmt.Sprintf("0x%x", transactionFeeBig)

		transaction := &models.Transaction{
			Hash:               transactionETL.Hash,
			LogIndex:           -1,
			Type:               "transaction",
			Method:             method,
			FromAddress:        transactionETL.FromAddress,
			ToAddress:          transactionETL.ToAddress,
			BlockNumber:        blockETL.Number,
			Version:            transactionETL.Version,
			Value:              transactionETL.Value,
			ValueDecimal:       valueDecimal,
			StepLimit:          transactionETL.StepLimit,
			Timestamp:          transactionETL.Timestamp,
			BlockTimestamp:     blockETL.Timestamp,
			Nid:                transactionETL.Nid,
			Nonce:              transactionETL.Nonce,
			TransactionIndex:   transactionETL.TransactionIndex,
			BlockHash:          blockETL.Hash,
			TransactionFee:     transactionFee,
			Signature:          transactionETL.Signature,
			DataType:           transactionETL.DataType,
			Data:               transactionETL.Data,
			CumulativeStepUsed: transactionETL.CumulativeStepUsed,
			StepUsed:           transactionETL.StepUsed,
			StepPrice:          transactionETL.StepPrice,
			ScoreAddress:       transactionETL.ScoreAddress,
			LogsBloom:          transactionETL.LogsBloom,
			Status:             transactionETL.Status,
		}

		transactions = append(transactions, transaction)
	}

	//////////
	// Logs //
	//////////
	for _, transactionETL := range blockETL.Transactions {
		for iL, logETL := range transactionETL.Logs {

			method := strings.Split(logETL.Indexed[0], "(")[0]
			if method == "ICXTransfer" {
				// Internal Transaction

				// From Address
				fromAddress := logETL.Indexed[1]

				// To Address
				toAddress := logETL.Indexed[2]

				// Value
				value := logETL.Indexed[3]

				// Transaction Decimal Value
				// Hex -> float64
				valueDecimal := stringHexToFloat64(value, 18)

				transaction := &models.Transaction{
					Hash:               transactionETL.Hash,
					LogIndex:           int64(iL), // NOTE logs will always be in the same order from ETL
					Type:               "log",
					Method:             method,
					FromAddress:        fromAddress,
					ToAddress:          toAddress,
					BlockNumber:        blockETL.Number,
					Version:            transactionETL.Version,
					Value:              value,
					ValueDecimal:       valueDecimal,
					StepLimit:          transactionETL.StepLimit,
					Timestamp:          transactionETL.Timestamp,
					BlockTimestamp:     blockETL.Timestamp,
					Nid:                transactionETL.Nid,
					Nonce:              transactionETL.Nonce,
					TransactionIndex:   transactionETL.TransactionIndex,
					BlockHash:          blockETL.Hash,
					TransactionFee:     "0x0",
					Signature:          transactionETL.Signature,
					DataType:           "",
					Data:               "",
					CumulativeStepUsed: "0x0",
					StepUsed:           "0x0",
					StepPrice:          "0x0",
					ScoreAddress:       logETL.Address,
					LogsBloom:          "",
					Status:             "0x1",
				}

				transactions = append(transactions, transaction)
			}
		}
	}

	return transactions
}
