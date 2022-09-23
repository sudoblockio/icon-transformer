package transformers

import (
	"fmt"
	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/redis"
	"go.uber.org/zap"
	"math/big"

	"github.com/sudoblockio/icon-transformer/models"
	"github.com/sudoblockio/icon-transformer/utils"
)

func transactions(blockETL *models.BlockETL) {

	var transactions []*models.Transaction

	//loaderChannel := crud.GetTransactionCrud().LoaderChannel
	for _, transactionETL := range blockETL.Transactions {

		// Method
		method := extractMethodFromTransactionETL(transactionETL)

		// Log Count
		logCount := int64(len(transactionETL.Logs))

		// Value Decimal
		valueDecimal := float64(0)
		if transactionETL.Value != "" {
			valueDecimal = utils.StringHexToFloat64(transactionETL.Value, 18)
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
			LogCount:           logCount,
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

		if config.Config.ProcessCounts {
			transactions = append(transactions, transaction)
		}

		crud.TransactionCrud.LoaderChannel <- transaction
		broadcastToWebsocketRedisChannel(blockETL, transaction, config.Config.RedisTransactionsChannel)
	}

	//////////
	// Logs //
	//////////
	for _, transactionETL := range blockETL.Transactions {
		for iL, logETL := range transactionETL.Logs {

			method := extractMethodFromLogETL(logETL)

			// NOTE 'ICXTransfer' is a protected name in Icon
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
				valueDecimal := utils.StringHexToFloat64(value, 18)

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

				if config.Config.ProcessCounts {
					transactions = append(transactions, transaction)
				}
				crud.TransactionCrud.LoaderChannel <- transaction
				broadcastToWebsocketRedisChannel(blockETL, transaction, config.Config.RedisTransactionsChannel)
			}
		}
	}
	if config.Config.ProcessCounts {
		transformBlocksToCountTransactions(blockETL, transactions)
	}
}

// Transactions count
func transformBlocksToCountTransactions(blockETL *models.BlockETL, transactions []*models.Transaction) {
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
