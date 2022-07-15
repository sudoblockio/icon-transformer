package transformers

import (
	"encoding/json"
	"errors"
	"sync"
	"time"

	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/kafka"
	"github.com/sudoblockio/icon-transformer/models"
	"github.com/sudoblockio/icon-transformer/redis"
	"github.com/sudoblockio/icon-transformer/service"
	"github.com/sudoblockio/icon-transformer/utils"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
)

func startBlocks() {
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
		processBlocks(blockETL)
	}
}

var blockProcessors = []func(a *models.BlockETL){
	transformBlocksToLoadBlock,
	transformBlocksToLoadTransactions,
	transformBlocksToLoadTransactionByAddresses,
	transformBlocksToLoadTransactionInternalByAddresses,
	transformBlocksToLoadLogs,
	transformBlocksToLoadTokenTransfers,
	transformBlocksToLoadTokenTransferByAddresses,
	transformBlocksToLoadTransactionByAddressCreateScores,
	transformBlocksToLoadAddresses,
	transformBlocksToLoadTokenAddresses,
}

var blockCounters = []func(etl *models.BlockETL){
	transformBlocksToChannelBlocks,
	transformBlocksToChannelTransactions,
	transformBlocksToChannelLogs,
	transformBlocksToChannelTokenTransfers,
	transformBlocksToCountBlocks,
	transformBlocksToCountTransactions,
	transformBlocksToCountLogs,
	transformBlocksToCountTokenTransfers,
	transformBlocksToServiceAddressBalance,
	transformBlocksToServiceTokenAddressBalance,
}

func runBlockProcessors(blockProcessors []func(blockETL *models.BlockETL), blockETL *models.BlockETL) {
	var wg sync.WaitGroup
	for _, f := range blockProcessors {
		wg.Add(1)

		f := f
		go func() {
			defer wg.Done()
			f(blockETL)
		}()
	}

	wg.Wait()
}

func processBlocks(blockETL *models.BlockETL) {
	// NOTE transform blockETL to various database views
	// NOTE blocks may be passed multiple times, loaders use upserts
	runBlockProcessors(blockProcessors, blockETL)

	// NOTE indexed loaders index messages by block number
	// NOTE each block number can only pass through once
	// Mostly these are items for keeping redis cache up to date
	_, err := crud.GetBlockIndexCrud().SelectOne(blockETL.Number)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		runBlockProcessors(blockCounters, blockETL)

		// Commit block
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

// Blocks loader
func transformBlocksToLoadBlock(blockETL *models.BlockETL) {
	block := transformBlockETLToBlock(blockETL)
	err := crud.GetBlockCrud().UpsertOne(block)
	if err != nil {
		zap.S().Fatal(err.Error())
	}
}

// Transactions loader
func transformBlocksToLoadTransactions(blockETL *models.BlockETL) {
	transactions := transformBlockETLToTransactions(blockETL)
	for _, transaction := range transactions {
		err := crud.GetTransactionCrud().UpsertOne(transaction)
		if err != nil {
			zap.S().Fatal(err.Error())
		}
	}
}

// Transaction by addresses loader
func transformBlocksToLoadTransactionByAddresses(blockETL *models.BlockETL) {
	transactionByAddresses := transformBlockETLToTransactionByAddresses(blockETL)
	for _, transactionByAddress := range transactionByAddresses {
		err := crud.GetTransactionByAddressCrud().UpsertOne(transactionByAddress)
		if err != nil {
			zap.S().Fatal(err.Error())
		}
	}
}

// Transaction internal by addresses loader
func transformBlocksToLoadTransactionInternalByAddresses(blockETL *models.BlockETL) {
	transactionInternalByAddresses := transformBlockETLToTransactionInternalByAddresses(blockETL)
	for _, transactionInternalByAddress := range transactionInternalByAddresses {
		err := crud.GetTransactionInternalByAddressCrud().UpsertOne(transactionInternalByAddress)
		if err != nil {
			zap.S().Fatal(err.Error())
		}
	}
}

// Logs loader
func transformBlocksToLoadLogs(blockETL *models.BlockETL) {
	//loaderChannel := crud.GetLogCrud().LoaderChannel

	logs := transformBlockETLToLogs(blockETL)
	for _, log := range logs {
		err := crud.GetLogCrud().UpsertOne(log)
		if err != nil {
			zap.S().Fatal(err.Error())
		}
		//loaderChannel <- log
	}
}

// Token transfers loader
func transformBlocksToLoadTokenTransfers(blockETL *models.BlockETL) {
	tokenTransfers := transformBlockETLToTokenTransfers(blockETL)
	for _, tokenTransfer := range tokenTransfers {
		err := crud.GetTokenTransferCrud().UpsertOne(tokenTransfer)
		if err != nil {
			zap.S().Fatal(err.Error())
		}
	}
}

// Token transfer by address loader
func transformBlocksToLoadTokenTransferByAddresses(blockETL *models.BlockETL) {
	tokenTransferByAddresses := transformBlockETLToTokenTransferByAddresses(blockETL)
	for _, tokenTransferByAddress := range tokenTransferByAddresses {
		err := crud.GetTokenTransferByAddressCrud().UpsertOne(tokenTransferByAddress)
		if err != nil {
			zap.S().Fatal(err.Error())
		}
	}
}

// Deprecated because this additionally needs a routine to be able to get the Tx to show up
//// Transaction create scores loader
//func transformBlocksToLoadTransactionCreateScores(blockETL *models.BlockETL) {
//	loaderChannel := crud.GetTransactionCreateScoreCrud().LoaderChannel
//
//	transactionCreateScores := transformBlockETLToTransactionCreateScores(blockETL)
//	for _, transactionCreateScore := range transactionCreateScores {
//		loaderChannel <- transactionCreateScore
//	}
//}

// Transaction by address for creating scores loader. Finds approval Txs for core submission and inserts into
//  transaction_by_address table so that the Txs are brought up in the list view when viewing a contract's Txs.
func transformBlocksToLoadTransactionByAddressCreateScores(blockETL *models.BlockETL) {
	transactionCreateScores := transformBlockETLToTransactionByAddressCreateScores(blockETL)
	for _, transactionCreateScore := range transactionCreateScores {
		err := crud.GetTransactionByAddressCrud().UpsertOne(transactionCreateScore)
		if err != nil {
			zap.S().Fatal(err.Error())
		}
	}
}

//// Transaction by address for creating scores loader. Finds approval Txs for core submission and inserts into
////  transaction_by_address table so that the Txs are brought up in the list view when viewing a contract's Txs.
//func transformBlocksToLoadTransactionCreateScores(blockETL *models.BlockETL) {
//	loaderChannel := crud.GetTransactionCrud().LoaderChannel
//
//	transactionCreateScores := transformBlockETLToTransactionByAddressCreateScores(blockETL)
//	for _, transactionCreateScore := range transactionCreateScores {
//		loaderChannel <- transactionCreateScore
//	}
//}

// Address loader
func transformBlocksToLoadAddresses(blockETL *models.BlockETL) {
	addresses := transformBlockETLToAddresses(blockETL)
	for _, address := range addresses {
		err := crud.GetAddressCrud().UpsertOneColsE(address, []string{"address", "is_contract"})
		if err != nil {
			zap.S().Fatal(err.Error())
		}
	}
}

// Address tokens loader
func transformBlocksToLoadTokenAddresses(blockETL *models.BlockETL) {
	tokenAddresses := transformBlockETLToTokenAddresses(blockETL)
	for _, tokenAddress := range tokenAddresses {
		err := crud.GetTokenAddressCrud().UpsertOne(tokenAddress)
		if err != nil {
			zap.S().Fatal(err.Error())
		}
	}
}

// Blocks channel
func transformBlocksToChannelBlocks(blockETL *models.BlockETL) {

	blockTimestamp := time.Unix(blockETL.Timestamp/1000000, 0)

	if time.Since(blockTimestamp) <= config.Config.TransformerRedisChannelThreshold ||
		int64(config.Config.TransformerRedisChannelThreshold) == 0 {
		// block is recent enough, calculate balances
		block := transformBlockETLToBlock(blockETL)
		blockJSON, _ := json.Marshal(block)
		redis.GetRedisClient().Publish(config.Config.RedisBlocksChannel, blockJSON)
	}
}

// Transactions channel
func transformBlocksToChannelTransactions(blockETL *models.BlockETL) {

	blockTimestamp := time.Unix(blockETL.Timestamp/1000000, 0)

	if time.Since(blockTimestamp) <= config.Config.TransformerRedisChannelThreshold ||
		int64(config.Config.TransformerRedisChannelThreshold) == 0 {
		transactions := transformBlockETLToTransactions(blockETL)
		for _, transaction := range transactions {
			transactionJSON, _ := json.Marshal(transaction)
			redis.GetRedisClient().Publish(config.Config.RedisTransactionsChannel, transactionJSON)
		}
	}
}

// Logs channel
func transformBlocksToChannelLogs(blockETL *models.BlockETL) {

	blockTimestamp := time.Unix(blockETL.Timestamp/1000000, 0)

	if time.Since(blockTimestamp) <= config.Config.TransformerRedisChannelThreshold ||
		int64(config.Config.TransformerRedisChannelThreshold) == 0 {
		logs := transformBlockETLToLogs(blockETL)
		for _, log := range logs {
			logJSON, _ := json.Marshal(log)
			redis.GetRedisClient().Publish(config.Config.RedisLogsChannel, logJSON)
		}
	}
}

// Token Transfers channel
func transformBlocksToChannelTokenTransfers(blockETL *models.BlockETL) {

	blockTimestamp := time.Unix(blockETL.Timestamp/1000000, 0)

	if time.Since(blockTimestamp) <= config.Config.TransformerRedisChannelThreshold ||
		int64(config.Config.TransformerRedisChannelThreshold) == 0 {
		tokenTransfers := transformBlockETLToTokenTransfers(blockETL)
		for _, tokenTransfer := range tokenTransfers {
			tokenTransferJSON, _ := json.Marshal(tokenTransfer)
			redis.GetRedisClient().Publish(config.Config.RedisTokenTransfersChannel, tokenTransferJSON)
		}
	}
}

// Block count
func transformBlocksToCountBlocks(blockETL *models.BlockETL) {
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
func transformBlocksToCountTransactions(blockETL *models.BlockETL) {
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
func transformBlocksToCountLogs(blockETL *models.BlockETL) {
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
func transformBlocksToCountTokenTransfers(blockETL *models.BlockETL) {
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
func transformBlocksToServiceAddressBalance(blockETL *models.BlockETL) {

	blockTimestamp := time.Unix(blockETL.Timestamp/1000000, 0)

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
			//addressCopy := &models.Address{}
			//copier.Copy(addressCopy, &address)

			// Insert to database
			//crud.GetAddressCrud().LoaderChannel <- addressCopy
			crud.GetAddressCrud().UpsertOneCols(address, []string{"address", "balance"})
		}
	}
}

// Token Address Balance
func transformBlocksToServiceTokenAddressBalance(blockETL *models.BlockETL) {

	blockTimestamp := time.Unix(blockETL.Timestamp/1000000, 0)

	if time.Since(blockTimestamp) <= config.Config.TransformerServiceCallThreshold {
		// block is recent enough, calculate balances

		tokenAddresses := transformBlockETLToTokenAddresses(blockETL)

		for _, tokenAddress := range tokenAddresses {

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

			//// Copy struct for pointer conflicts
			//tokenAddressCopy := &models.TokenAddress{}
			//copier.Copy(tokenAddressCopy, &tokenAddress)

			// Insert to database
			crud.GetTokenAddressCrud().UpsertOneCols(tokenAddress, []string{"address", "balance", "token_contract_address"})
		}
	}
}
