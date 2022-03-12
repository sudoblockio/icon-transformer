package transformers

import (
	"encoding/json"

	"github.com/sudoblockio/icon-go-worker/models"
)

func transformBlockETLToLogs(blockETL *models.BlockETL) []*models.Log {

	logs := []*models.Log{}

	//////////
	// Logs //
	//////////
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
			logs = append(logs, log)
		}
	}

	return logs
}
