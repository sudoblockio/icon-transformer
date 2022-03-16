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

		// TODO export all go func to their own functions

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

			// Block count
			go func() {
				countKey := config.Config.RedisKeyPrefix + "block_count"
				_, err = redis.GetRedisClient().IncCountBy(countKey, 1)
				if err != nil {
					zap.S().Warn(
						"Routine=Transformer,",
						" BlockNumber=", blockETL.Number,
						" Step=", "Inc count block",
						" Error=", err.Error(),
					)
				}
			}()

			// Transactions count
			go func() {
				transactions := transformBlockETLToTransactions(blockETL)

				/////////////////
				// Total count //
				/////////////////

				// Get count
				countRegular := int64(0)
				countInternal := int64(0)
				for _, transaction := range transactions {
					if transaction.Type == "transaction" {
						// Regular

						countRegular++
					} else if transaction.Type == "log" {
						// Internal

						countInternal++
					}
				}

				// Set count
				countKeyRegular := config.Config.RedisKeyPrefix + "transaction_regular_count"
				countKeyInternal := config.Config.RedisKeyPrefix + "transaction_internal_count"

				_, err = redis.GetRedisClient().IncCountBy(countKeyRegular, countRegular)
				if err != nil {
					zap.S().Warn(
						"Routine=Transformer,",
						" BlockNumber=", blockETL.Number,
						" Step=", "Inc count transactions regular",
						" Error=", err.Error(),
					)
				}
				_, err = redis.GetRedisClient().IncCountBy(countKeyInternal, countInternal)
				if err != nil {
					zap.S().Warn(
						"Routine=Transformer,",
						" BlockNumber=", blockETL.Number,
						" Step=", "Inc count transactions internal",
						" Error=", err.Error(),
					)
				}

				//////////////////////
				// Count by address //
				//////////////////////

				// Get count
				countByAddressRegular := map[string]int64{}
				countByAddressInternal := map[string]int64{}
				for _, transaction := range transactions {
					fromAddress := transaction.FromAddress
					toAddress := transaction.ToAddress

					if transaction.Type == "transaction" {
						// Regular

						// From address
						if _, ok := countByAddressRegular[fromAddress]; ok == true {
							countByAddressRegular[fromAddress]++
						} else {
							countByAddressRegular[fromAddress] = 1
						}

						// To address
						if _, ok := countByAddressRegular[toAddress]; ok == true {
							countByAddressRegular[toAddress]++
						} else {
							countByAddressRegular[toAddress] = 1
						}

					} else if transaction.Type == "log" {
						// Internal

						// From address
						if _, ok := countByAddressInternal[fromAddress]; ok == true {
							countByAddressInternal[fromAddress]++
						} else {
							countByAddressInternal[fromAddress] = 1
						}

						// To address
						if _, ok := countByAddressInternal[toAddress]; ok == true {
							countByAddressInternal[toAddress]++
						} else {
							countByAddressInternal[toAddress] = 1
						}
					}
				}

				// Set count
				for address, count := range countByAddressRegular {
					// Regular transactions

					countByAddressKey := config.Config.RedisKeyPrefix + "transaction_regular_count_by_address_" + address
					_, err = redis.GetRedisClient().IncCountBy(countByAddressKey, count)
					if err != nil {
						zap.S().Warn(
							"Routine=Transformer,",
							" BlockNumber=", blockETL.Number,
							" Address=", address,
							" Step=", "Inc count transaction regular by address",
							" Error=", err.Error(),
						)
					}
				}
				for address, count := range countByAddressInternal {
					// Internal transactions

					countByAddressKey := config.Config.RedisKeyPrefix + "transaction_internal_count_by_address_" + address
					_, err = redis.GetRedisClient().IncCountBy(countByAddressKey, count)
					if err != nil {
						zap.S().Warn(
							"Routine=Transformer,",
							" BlockNumber=", blockETL.Number,
							" Address=", address,
							" Step=", "Inc count transaction internal by address",
							" Error=", err.Error(),
						)
					}
				}
			}()

			// Logs count
			go func() {
				logs := transformBlockETLToLogs(blockETL)

				/////////////////
				// Total count //
				/////////////////

				// Get count
				count := int64(len(logs))

				// Set count
				countKey := config.Config.RedisKeyPrefix + "log_count"

				_, err = redis.GetRedisClient().IncCountBy(countKey, count)
				if err != nil {
					zap.S().Warn(
						"Routine=Transformer,",
						" BlockNumber=", blockETL.Number,
						" Step=", "Inc count log",
						" Error=", err.Error(),
					)
				}

				//////////////////////
				// Count by address //
				//////////////////////
				countByAddress := map[string]int64{}

				// Get count
				for _, log := range logs {
					address := log.Address

					if _, ok := countByAddress[address]; ok == true {
						countByAddress[address]++
					} else {
						countByAddress[address] = 1
					}
				}

				// Set count
				for address, count := range countByAddress {

					countByAddressKey := config.Config.RedisKeyPrefix + "log_count_by_address_" + address
					_, err = redis.GetRedisClient().IncCountBy(countByAddressKey, count)
					if err != nil {
						zap.S().Warn(
							"Routine=Transformer,",
							" BlockNumber=", blockETL.Number,
							" Address=", address,
							" Step=", "Inc count log by address",
							" Error=", err.Error(),
						)
					}
				}
			}()

			// Token Transfers count
			go func() {
				tokenTransfers := transformBlockETLToTokenTransfers(blockETL)

				/////////////////
				// Total count //
				/////////////////

				// Get count
				count := int64(len(tokenTransfers))

				// Set count
				countKey := config.Config.RedisKeyPrefix + "token_transfer_count"

				_, err = redis.GetRedisClient().IncCountBy(countKey, count)
				if err != nil {
					zap.S().Warn(
						"Routine=Transformer,",
						" BlockNumber=", blockETL.Number,
						" Step=", "Inc count token transfers",
						" Error=", err.Error(),
					)
				}

				//////////////////////
				// Count by address //
				//////////////////////

				// Get count
				countByAddress := map[string]int64{}
				countByTokenContract := map[string]int64{}
				for _, tokenTransfer := range tokenTransfers {
					fromAddress := tokenTransfer.FromAddress
					toAddress := tokenTransfer.ToAddress
					tokenContractAddress := tokenTransfer.TokenContractAddress

					// From address
					if _, ok := countByAddress[fromAddress]; ok == true {
						countByAddress[fromAddress]++
					} else {
						countByAddress[fromAddress] = 1
					}

					// To address
					if _, ok := countByAddress[toAddress]; ok == true {
						countByAddress[toAddress]++
					} else {
						countByAddress[toAddress] = 1
					}

					// Token Contract Address
					if _, ok := countByTokenContract[tokenContractAddress]; ok == true {
						countByTokenContract[tokenContractAddress]++
					} else {
						countByTokenContract[tokenContractAddress] = 1
					}
				}

				// Set count
				for address, count := range countByAddress {
					// Count by address

					countByAddressKey := config.Config.RedisKeyPrefix + "token_transfer_count_by_address_" + address
					_, err = redis.GetRedisClient().IncCountBy(countByAddressKey, count)
					if err != nil {
						zap.S().Warn(
							"Routine=Transformer,",
							" BlockNumber=", blockETL.Number,
							" Address=", address,
							" Step=", "Inc count token transfer by address",
							" Error=", err.Error(),
						)
					}
				}
				for tokenContract, count := range countByTokenContract {
					// Count by token contract

					countByTokenContractKey := config.Config.RedisKeyPrefix + "token_transfer_count_by_token_contract_" + tokenContract
					_, err = redis.GetRedisClient().IncCountBy(countByTokenContractKey, count)
					if err != nil {
						zap.S().Warn(
							"Routine=Transformer,",
							" BlockNumber=", blockETL.Number,
							" TokenContract=", tokenContract,
							" Step=", "Inc count token transfer by token contract",
							" Error=", err.Error(),
						)
					}
				}
			}()

			//////////////////
			// Commit Block //
			//////////////////
			blockIndex := &models.BlockIndex{Number: blockETL.Number}
			err = crud.GetBlockIndexCrud().InsertOne(blockIndex)

		} else if err != nil {
			// ERROR inserting block index
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

// Testing generic function
// func sendToLoader(
// 	loader chan *interface{},
// 	transformerFunc func(*models.BlockETL) *interface{},
// 	block *models.BlockETL,
// ) {
//
// 	transformedObject := transformerFunc(blockETL)
//
// 	transformedObjects, ok := transformedObject.([]interface{})
// 	if ok == true {
// 		for _, object := range transformedObjects {
// 			loader <- object
// 		}
// 	} else {
// 		loader <- transformedObject
// 	}
// }
