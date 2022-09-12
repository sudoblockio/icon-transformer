package transformers

import (
	"encoding/json"
	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/redis"
	"go.uber.org/zap"

	"github.com/sudoblockio/icon-transformer/models"
)

func logs(blockETL *models.BlockETL) {

	logs := []*models.Log{}
	for _, transactionETL := range blockETL.Transactions {
		for iL, logETL := range transactionETL.Logs {

			// Method
			method := extractMethodFromLogETL(logETL)

			// Data
			data, _ := json.Marshal(logETL.Data)

			// Indexed
			indexed, _ := json.Marshal(logETL.Indexed)

			log := &models.Log{
				TransactionHash: transactionETL.Hash,
				LogIndex:        int64(iL),
				Address:         logETL.Address,
				BlockNumber:     blockETL.Number,
				Method:          method,
				Data:            string(data),
				Indexed:         string(indexed),
				BlockTimestamp:  blockETL.Timestamp,
			}

			if config.Config.ProcessCounts {
				logs = append(logs, log)
			}
			crud.LogCrud.LoaderChannel <- log
			broadcastToWebsocketRedisChannel(blockETL, log, config.Config.RedisLogsChannel)
		}
	}
	if config.Config.ProcessCounts {
		transformBlocksToCountLogs(blockETL, logs)
	}
}

func transformBlocksToCountLogs(blockETL *models.BlockETL, logs []*models.Log) {
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
