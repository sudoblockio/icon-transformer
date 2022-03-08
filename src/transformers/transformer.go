package transformers

import (
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

	// Transaction Count
	transactionCount := int64(len(blockETL.Transactions))

	// Log Count
	logCount := int64(0)
	for _, transaction := range blockETL.Transactions {
		logCount += int64(len(transaction.Logs))
	}

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
		FailedTransactionCount:    0,  // TODO
		InternalTransactionCount:  0,  // TODO
		TransactionAmount:         "", // TODO
		InternalTransactionAmount: "", // TODO
		TransactionFees:           "", // TODO
	}
}
