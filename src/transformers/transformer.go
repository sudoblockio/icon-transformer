package transformers

import (
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
	for _, transaction := range blockETL.Transactions {

		// transactionAmount
		transactionAmountBig := big.NewInt(0)
		if transaction.Value != "" {
			transactionAmountBig.SetString(transaction.Value[2:], 16)
		}
		sumTransactionAmountBig = sumTransactionAmountBig.Add(sumTransactionAmountBig, transactionAmountBig)

		// transactionFees
		stepPriceBig := big.NewInt(0)
		if transaction.StepPrice != "" {
			stepPriceBig.SetString(transaction.StepPrice[2:], 16)
		}
		stepUsedBig := big.NewInt(0)
		if transaction.StepUsed != "" {
			stepUsedBig.SetString(transaction.StepUsed[2:], 16)
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
	for _, transaction := range blockETL.Transactions {

		if transaction.Status == "0x0" {
			// failedTransactionCount
			failedTransactionCount++
		}
	}

	//////////
	// Logs //
	//////////
	logCount := int64(0)
	for _, transaction := range blockETL.Transactions {
		logCount += int64(len(transaction.Logs))
	}

	///////////////////////////
	// Internal Transactions //
	///////////////////////////
	internalTransactionCount := int64(0)
	internalTransactionAmount := "0x0"

	sumInternalTransactionAmountBig := big.NewInt(0)
	for _, transaction := range blockETL.Transactions {
		for _, log := range transaction.Logs {
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
