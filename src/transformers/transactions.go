package transformers

import (
	"fmt"
	"math/big"

	"github.com/sudoblockio/icon-go-worker/utils"

	"github.com/sudoblockio/icon-go-worker/models"
)

func transformBlockETLToTransactions(blockETL *models.BlockETL) []*models.Transaction {

	transactions := []*models.Transaction{}

	//////////////////
	// Transactions //
	//////////////////
	for _, transactionETL := range blockETL.Transactions {

		// Method
		method := extractMethodFromTransactionETL(transactionETL)

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

		transactions = append(transactions, transaction)
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

				transactions = append(transactions, transaction)
			}
		}
	}

	return transactions
}
