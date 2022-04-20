package transformers

import (
	"encoding/json"
	"errors"
	"time"

	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"

	"github.com/jinzhu/copier"
	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/kafka"
	"github.com/sudoblockio/icon-transformer/models"
	"github.com/sudoblockio/icon-transformer/redis"
	"github.com/sudoblockio/icon-transformer/service"
	"github.com/sudoblockio/icon-transformer/utils"
)

func Start() {
	go start()
}

func start() {
	kafkaBlocksTopic := config.Config.KafkaBlocksTopic

	// Input channels
	kafkaBlocksTopicChannel := kafka.KafkaTopicConsumer.TopicChannels[kafkaBlocksTopic]

	zap.S().Debug("Blocks transformer: started working")
	for {

		///////////////////
		// Kafka Message //
		///////////////////

		consumerTopicMsg := <-kafkaBlocksTopicChannel
		blockETL := &models.BlockETL{}
		err := proto.Unmarshal(consumerTopicMsg.Value, blockETL)
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

		// Block loader
		transformToLoadBlock(blockETL)

		// Transaction loader
		transformToLoadTransactions(blockETL)

		// Transaction by address loader
		transformToLoadTransactionByAddresses(blockETL)

		// Transaction internal by address loader
		transformToLoadTransactionInternalByAddresses(blockETL)

		// Log loader
		transformToLoadLogs(blockETL)

		// Token transfer loader
		transformToLoadTokenTransfers(blockETL)

		// Token transfer by address loader
		transformToLoadTokenTransferByAddresses(blockETL)

		// Transaction create score loader
		transformToLoadTransactionCreateScores(blockETL)

		// Address loader
		transformToLoadAddresses(blockETL)

		// Address token loader
		transformToLoadTokenAddresses(blockETL)

		/////////////////////
		// Indexed loaders //
		/////////////////////
		// NOTE indexed loaders index messages by block number
		// NOTE each block number can only pass through once

		_, err = crud.GetBlockIndexCrud().SelectOne(blockETL.Number)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Block not seen yet, proceed

			////////////////////
			// Redis channels //
			////////////////////

			// Blocks channel
			transformToChannelBlocks(blockETL)

			// Transactions channel
			transformToChannelTransactions(blockETL)

			// Logs channel
			transformToChannelLogs(blockETL)

			// Token transfers channel
			transformToChannelTokenTransfers(blockETL)

			//////////////
			// Counters //
			//////////////

			// Block count
			transformToCountBlocks(blockETL)

			// Transactions count
			transformToCountTransactions(blockETL)

			// Logs count
			transformToCountLogs(blockETL)

			// Token transfers count
			transformToCountTokenTransfers(blockETL)

			///////////////////
			// Service Calls //
			///////////////////

			// Address Balance
			transformToServiceAddressBalance(blockETL)

			// Token Address Balance
			transformToServiceTokenAddressBalance(blockETL)

			//////////////////
			// Commit block //
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

// Blocks loader
func transformToLoadBlock(blockETL *models.BlockETL) {
	loaderChannel := crud.GetBlockCrud().LoaderChannel

	block := transformBlockETLToBlock(blockETL)
	loaderChannel <- block
}

// Transactions loader
func transformToLoadTransactions(blockETL *models.BlockETL) {
	loaderChannel := crud.GetTransactionCrud().LoaderChannel

	transactions := transformBlockETLToTransactions(blockETL)
	for _, transaction := range transactions {
		loaderChannel <- transaction
	}
}

// Transaction by addresses loader
func transformToLoadTransactionByAddresses(blockETL *models.BlockETL) {
	loaderChannel := crud.GetTransactionByAddressCrud().LoaderChannel

	transactionByAddresses := transformBlockETLToTransactionByAddresses(blockETL)
	for _, transactionByAddress := range transactionByAddresses {
		loaderChannel <- transactionByAddress
	}
}

// Transaction internal by addresses loader
func transformToLoadTransactionInternalByAddresses(blockETL *models.BlockETL) {
	loaderChannel := crud.GetTransactionInternalByAddressCrud().LoaderChannel

	transactionInternalByAddresses := transformBlockETLToTransactionInternalByAddresses(blockETL)
	for _, transactionInternalByAddress := range transactionInternalByAddresses {
		loaderChannel <- transactionInternalByAddress
	}
}

// Logs loader
func transformToLoadLogs(blockETL *models.BlockETL) {
	loaderChannel := crud.GetLogCrud().LoaderChannel

	logs := transformBlockETLToLogs(blockETL)
	for _, log := range logs {
		loaderChannel <- log
	}
}

// Token transfers loader
func transformToLoadTokenTransfers(blockETL *models.BlockETL) {
	loaderChannel := crud.GetTokenTransferCrud().LoaderChannel

	tokenTransfers := transformBlockETLToTokenTransfers(blockETL)
	for _, tokenTransfer := range tokenTransfers {
		loaderChannel <- tokenTransfer
	}
}

// Token transfer by address loader
func transformToLoadTokenTransferByAddresses(blockETL *models.BlockETL) {
	loaderChannel := crud.GetTokenTransferByAddressCrud().LoaderChannel

	tokenTransferByAddresses := transformBlockETLToTokenTransferByAddresses(blockETL)
	for _, tokenTransferByAddress := range tokenTransferByAddresses {
		loaderChannel <- tokenTransferByAddress
	}
}

// Transaction create scores loader
func transformToLoadTransactionCreateScores(blockETL *models.BlockETL) {
	loaderChannel := crud.GetTransactionCreateScoreCrud().LoaderChannel

	transactionCreateScores := transformBlockETLToTransactionCreateScores(blockETL)
	for _, transactionCreateScore := range transactionCreateScores {
		loaderChannel <- transactionCreateScore
	}
}

// Address loader
func transformToLoadAddresses(blockETL *models.BlockETL) {
	loaderChannel := crud.GetAddressCrud().LoaderChannel

	addresses := transformBlockETLToAddresses(blockETL)
	for _, address := range addresses {
		loaderChannel <- address
	}
}

// Address tokens loader
func transformToLoadTokenAddresses(blockETL *models.BlockETL) {
	loaderChannel := crud.GetTokenAddressCrud().LoaderChannel

	tokenAddresses := transformBlockETLToTokenAddresses(blockETL)
	for _, tokenAddress := range tokenAddresses {
		loaderChannel <- tokenAddress
	}
}

// Blocks channel
func transformToChannelBlocks(blockETL *models.BlockETL) {
	block := transformBlockETLToBlock(blockETL)
	blockJSON, _ := json.Marshal(block)
	redis.GetRedisClient().Publish(config.Config.RedisBlocksChannel, blockJSON)
}

// Transactions channel
func transformToChannelTransactions(blockETL *models.BlockETL) {
	transactions := transformBlockETLToTransactions(blockETL)
	for _, transaction := range transactions {
		transactionJSON, _ := json.Marshal(transaction)
		redis.GetRedisClient().Publish(config.Config.RedisTransactionsChannel, transactionJSON)
	}
}

// Logs channel
func transformToChannelLogs(blockETL *models.BlockETL) {
	logs := transformBlockETLToLogs(blockETL)
	for _, log := range logs {
		logJSON, _ := json.Marshal(log)
		redis.GetRedisClient().Publish(config.Config.RedisLogsChannel, logJSON)
	}
}

// Token Transfers channel
func transformToChannelTokenTransfers(blockETL *models.BlockETL) {
	tokenTransfers := transformBlockETLToTokenTransfers(blockETL)
	for _, tokenTransfer := range tokenTransfers {
		tokenTransferJSON, _ := json.Marshal(tokenTransfer)
		redis.GetRedisClient().Publish(config.Config.RedisTokenTransfersChannel, tokenTransferJSON)
	}
}

// Block count
func transformToCountBlocks(blockETL *models.BlockETL) {
	countKey := config.Config.RedisKeyPrefix + "block_count"
	_, err := redis.GetRedisClient().IncCountBy(countKey, 1)
	if err != nil {
		zap.S().Warn(
			"Routine=Transformer,",
			" BlockNumber=", blockETL.Number,
			" Step=", "Inc count block",
			" Error=", err.Error(),
		)
	}
}

// Transactions count
func transformToCountTransactions(blockETL *models.BlockETL) {
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

	_, err := redis.GetRedisClient().IncCountBy(countKeyRegular, countRegular)
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
}

// Logs count
func transformToCountLogs(blockETL *models.BlockETL) {
	logs := transformBlockETLToLogs(blockETL)

	/////////////////
	// Total count //
	/////////////////

	// Get count
	count := int64(len(logs))

	// Set count
	countKey := config.Config.RedisKeyPrefix + "log_count"

	_, err := redis.GetRedisClient().IncCountBy(countKey, count)
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
}

// Token Transfers count
func transformToCountTokenTransfers(blockETL *models.BlockETL) {
	tokenTransfers := transformBlockETLToTokenTransfers(blockETL)

	/////////////////
	// Total count //
	/////////////////

	// Get count
	count := int64(len(tokenTransfers))

	// Set count
	countKey := config.Config.RedisKeyPrefix + "token_transfer_count"

	_, err := redis.GetRedisClient().IncCountBy(countKey, count)
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
}

// Address Balance
func transformToServiceAddressBalance(blockETL *models.BlockETL) {

	blockTimestamp := time.Unix(int64(blockETL.Timestamp/1000000), 0)

	if time.Since(blockTimestamp) <= config.Config.TransformerServiceCallThreshold {
		// block is recent enough, calculate balances

		addresses := transformBlockETLToAddresses(blockETL)

		for _, address := range addresses {

			// Node call
			balance, err := service.IconNodeServiceGetBalance(address.Address)
			if err != nil {
				// Icon node error
				zap.S().Warn("Routine=Balance, Address=", address.Address, " - Error: ", err.Error())
				continue
			}

			// Hex -> float64
			address.Balance = utils.StringHexToFloat64(balance, 18)

			////////////////////
			// Staked Balance //
			////////////////////
			stakedBalance, err := service.IconNodeServiceGetStakedBalance(address.Address)
			if err != nil {
				// Icon node error
				zap.S().Warn("Routine=Balance, Address=", address.Address, " - Error: ", err.Error())
				continue
			}

			// Hex -> float64
			address.Balance += utils.StringHexToFloat64(stakedBalance, 18)

			// Copy struct for pointer conflicts
			addressCopy := &models.Address{}
			copier.Copy(addressCopy, &address)

			// Insert to database
			crud.GetAddressCrud().LoaderChannel <- addressCopy
		}
	}
}

// Token Address Balance
func transformToServiceTokenAddressBalance(blockETL *models.BlockETL) {

	blockTimestamp := time.Unix(int64(blockETL.Timestamp/1000000), 0)

	if time.Since(blockTimestamp) <= config.Config.TransformerServiceCallThreshold {
		// block is recent enough, calculate balances

		tokenAddresses := transformBlockETLToTokenAddresses(blockETL)

		for _, tokenAddress := range *tokenAddresses {

			/////////////
			// Balance //
			/////////////

			// Node call
			balance, err := service.IconNodeServiceGetTokenBalance(tokenAddress.TokenContractAddress, tokenAddress.Address)
			if err != nil {
				// Icon node error
				zap.S().Warn("Routine=TokenAddressBalanceRoutine, Address=", tokenAddress.Address, " - Error: ", err.Error())
				continue
			}

			// Hex -> float64
			decimalBase, err := service.IconNodeServiceGetTokenDecimalBase(tokenAddress.TokenContractAddress)
			if err != nil {
				// Icon node error
				zap.S().Warn("Routine=TokenAddressBalanceRoutine - Error: ", err.Error())
				continue
			}
			tokenAddress.Balance = utils.StringHexToFloat64(balance, decimalBase)

			// Copy struct for pointer conflicts
			tokenAddressCopy := &models.TokenAddress{}
			copier.Copy(tokenAddressCopy, &tokenAddress)

			// Insert to database
			crud.GetTokenAddressCrud().LoaderChannel <- tokenAddressCopy
		}
	}
}
