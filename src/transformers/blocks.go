package transformers

import (
	"fmt"
	"math/big"

	"github.com/sudoblockio/icon-transformer/models"
)

func transformBlockETLToBlock(blockETL *models.BlockETL) *models.Block {

	//////////////////
	// Transactions //
	//////////////////
	transactionCount := int64(len(blockETL.Transactions))
	transactionAmount := "0x0"
	transactionFees := "0x0"

	sumTransactionAmountBig := big.NewInt(0)
	sumTransactionFeesBig := big.NewInt(0)
	for _, transactionETL := range blockETL.Transactions {

		// transactionAmount
		transactionAmountBig := big.NewInt(0)
		if transactionETL.Value != "" {
			transactionAmountBig.SetString(transactionETL.Value[2:], 16)
		}
		sumTransactionAmountBig = sumTransactionAmountBig.Add(sumTransactionAmountBig, transactionAmountBig)

		// transactionFees
		stepPriceBig := big.NewInt(0)
		if transactionETL.StepPrice != "" {
			stepPriceBig.SetString(transactionETL.StepPrice[2:], 16)
		}
		stepUsedBig := big.NewInt(0)
		if transactionETL.StepUsed != "" {
			stepUsedBig.SetString(transactionETL.StepUsed[2:], 16)
		}
		transactionFeeBig := stepUsedBig.Mul(stepUsedBig, stepPriceBig)
		sumTransactionFeesBig = sumTransactionFeesBig.Add(sumTransactionFeesBig, transactionFeeBig)
	}
	transactionAmount = fmt.Sprintf("0x%x", sumTransactionAmountBig)
	transactionFees = fmt.Sprintf("0x%x", sumTransactionFeesBig)

	/////////////////////////
	// Failed Transactions //
	/////////////////////////
	failedTransactionCount := int64(0)
	for _, transactionETL := range blockETL.Transactions {

		if transactionETL.Status == "0x0" {
			// failedTransactionCount
			failedTransactionCount++
		}
	}

	//////////
	// Logs //
	//////////
	logCount := int64(0)
	for _, transactionETL := range blockETL.Transactions {
		logCount += int64(len(transactionETL.Logs))
	}

	///////////////////////////
	// Internal Transactions //
	///////////////////////////
	internalTransactionCount := int64(0)
	internalTransactionAmount := "0x0"

	sumInternalTransactionAmountBig := big.NewInt(0)
	for _, transactionETL := range blockETL.Transactions {
		for _, logETL := range transactionETL.Logs {
			method := extractMethodFromLogETL(logETL)

			if method == "ICXTransfer" {
				// internalTransactionCount
				internalTransactionCount++

				// internalTransactionAmount
				internalTransactionAmountBig := big.NewInt(0)
				internalTransactionAmountBig.SetString(logETL.Indexed[3][2:], 16)
				sumInternalTransactionAmountBig = sumInternalTransactionAmountBig.Add(sumInternalTransactionAmountBig, internalTransactionAmountBig)
			}
		}
	}
	internalTransactionAmount = fmt.Sprintf("0x%x", sumInternalTransactionAmountBig)

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
		TransactionAmount:         transactionAmount,
		TransactionFees:           transactionFees,
		FailedTransactionCount:    failedTransactionCount,
		InternalTransactionCount:  internalTransactionCount,
		InternalTransactionAmount: internalTransactionAmount,
	}
}
