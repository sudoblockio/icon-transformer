package transformers

import (
	"encoding/json"
	"errors"

	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"

	"github.com/sudoblockio/icon-go-worker/config"
	"github.com/sudoblockio/icon-go-worker/crud"
	"github.com/sudoblockio/icon-go-worker/kafka"
	"github.com/sudoblockio/icon-go-worker/models"
	"github.com/sudoblockio/icon-go-worker/redis"
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
	logLoaderChannel := crud.GetLogCrud().LoaderChannel
	tokenTransferLoaderChannel := crud.GetTokenTransferCrud().LoaderChannel
	transactionCreateScoreLoaderChannel := crud.GetTransactionCreateScoreCrud().LoaderChannel

	zap.S().Debug("Blocks transformer: started working")
	for {

		///////////////////
		// Kafka Message //
		///////////////////

		consumerTopicMsg := <-kafkaBlocksTopicChannel
		blockETL, err := convertToBlockETLProtoBuf(consumerTopicMsg.Value)
		if err != nil {
			zap.S().Warn(
				"Routine=", "Transformer",
				" Partition=", consumerTopicMsg.Partition,
				" Offset=", consumerTopicMsg.Offset,
				" Key=", consumerTopicMsg.Key,
				" Value=", consumerTopicMsg.Value,
				" Step=", "Parse block ETL from kafka proto",
				" Error=", err.Error(),
			)
			continue
		}
		zap.S().Info("Transformer: Processing block #", blockETL.Number)

		/////////////
		// Loaders //
		/////////////
		// NOTE transform blockETL to various database views
		// NOTE blocks may be passed multiple times, loaders use upserts

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

		// Log Loader
		go func() {
			logs := transformBlockETLToLogs(blockETL)
			for _, log := range logs {
				logLoaderChannel <- log
			}
		}()

		// Token Transfer Loader
		go func() {
			tokenTransfers := transformBlockETLToTokenTransfers(blockETL)
			for _, tokenTransfer := range tokenTransfers {
				tokenTransferLoaderChannel <- tokenTransfer
			}
		}()

		// Transaction Create Score Loader
		go func() {
			transactionCreateScores := transformBlockETLToTransactionCreateScores(blockETL)
			for _, transactionCreateScore := range transactionCreateScores {
				transactionCreateScoreLoaderChannel <- transactionCreateScore
			}
		}()

		/////////////////////
		// Indexed loaders //
		/////////////////////
		// NOTE indexed loaders index messages by block number
		// NOTE each block number can only pass through once
		_, err = crud.GetBlockIndexCrud().SelectOne(blockETL.Number)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Block not seen yet, proceed

			////////////////////
			// Redis Channels //
			////////////////////

			// Blocks channel
			go func() {
				block := transformBlockETLToBlock(blockETL)
				blockJSON, _ := json.Marshal(block)
				redis.GetRedisClient().Publish(config.Config.RedisBlocksChannel, blockJSON)
			}()

			// Transactions channel
			go func() {
				transactions := transformBlockETLToTransactions(blockETL)
				for _, transaction := range transactions {
					transactionJSON, _ := json.Marshal(transaction)
					redis.GetRedisClient().Publish(config.Config.RedisTransactionsChannel, transactionJSON)
				}
			}()

			// Logs channel
			go func() {
				logs := transformBlockETLToLogs(blockETL)
				for _, log := range logs {
					logJSON, _ := json.Marshal(log)
					redis.GetRedisClient().Publish(config.Config.RedisLogsChannel, logJSON)
				}
			}()

			// Token Transfers channel
			go func() {
				tokenTransfers := transformBlockETLToTokenTransfers(blockETL)
				for _, tokenTransfer := range tokenTransfers {
					tokenTransferJSON, _ := json.Marshal(tokenTransfer)
					redis.GetRedisClient().Publish(config.Config.RedisTokenTransfersChannel, tokenTransferJSON)
				}
			}()

			//////////////
			// Counters //
			//////////////
			// TODO

			//////////////////
			// Commit Block //
			//////////////////
			blockIndex := &models.BlockIndex{Number: blockETL.Number}
			err = crud.GetBlockIndexCrud().InsertOne(blockIndex)

		} else if err != nil {
			// ERROR
			zap.S().Warn(
				"Routine=Transformer,",
				" BlockNumber=", blockETL.Number,
				" Step=", "Insert block index into postgres",
				" Error=", err.Error(),
			)
		}
	}
}

func convertToBlockETLProtoBuf(value []byte) (*models.BlockETL, error) {
	block := &models.BlockETL{}
	err := proto.Unmarshal(value, block)
	return block, err
}
