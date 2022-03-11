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
	transactionLoaderChannel := crud.GetTransactionCrud().LoaderChannel
	tokenTransferLoaderChannel := crud.GetTokenTransferCrud().LoaderChannel

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

		// Token Transfer Loader
		go func() {
			tokenTransfers := transformBlockETLToTokenTransfers(blockETL)
			for _, tokenTransfer := range tokenTransfers {
				tokenTransferLoaderChannel <- tokenTransfer
			}
		}()

		// Token Holder Loader
		// TODO

		// Transaction Create Score Loader
		// TODO
	}
}

func convertToBlockETLProtoBuf(value []byte) (*models.BlockETL, error) {
	block := &models.BlockETL{}
	err := proto.Unmarshal(value, block)
	return block, err
}
