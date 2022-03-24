package transformers

import (
	"fmt"
	"math/big"

	"github.com/sudoblockio/icon-transformer/models"
	"github.com/sudoblockio/icon-transformer/service"
	"github.com/sudoblockio/icon-transformer/utils"
	"go.uber.org/zap"
)

func transformBlockETLToTokenTransfers(blockETL *models.BlockETL) []*models.TokenTransfer {

	tokenTransfers := []*models.TokenTransfer{}

	//////////
	// Logs //
	//////////
	for _, transactionETL := range blockETL.Transactions {
		for iL, logETL := range transactionETL.Logs {

			// NOTE check for specific log definition
			// NOTE 'Transfer' is not a protected name in Icon
			if logETL.Indexed[0] == "Transfer(Address,Address,int,bytes)" && len(logETL.Indexed) == 4 {
				// Token Transfers

				// Token Contract Address
				tokenContractAddress := logETL.Address

				// From Address
				fromAddress := logETL.Indexed[1]

				// To Address
				toAddress := logETL.Indexed[2]

				// Value
				value := logETL.Indexed[3]

				// Transaction Decimal Value
				// NOTE every token has a different decimal base
				tokenDecimalBase, err := service.IconNodeServiceGetTokenDecimalBase(tokenContractAddress)
				if err != nil {
					zap.S().Fatal(err)
				}

				valueDecimal := utils.StringHexToFloat64(value, tokenDecimalBase)

				// Block Timestamp
				blockTimestamp := blockETL.Timestamp

				// Token Contract Name
				tokenContractName, err := service.IconNodeServiceGetTokenContractName(tokenContractAddress)
				if err != nil {
					zap.S().Fatal(err)
				}

				// Token Contract Symbol
				tokenContractSymbol, err := service.IconNodeServiceGetTokenContractSymbol(tokenContractAddress)
				if err != nil {
					zap.S().Fatal(err)
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

				tokenTransfer := &models.TokenTransfer{
					TransactionHash:      transactionETL.Hash,
					LogIndex:             int64(iL),
					TokenContractAddress: tokenContractAddress,
					FromAddress:          fromAddress,
					ToAddress:            toAddress,
					BlockNumber:          blockETL.Number,
					Value:                value,
					ValueDecimal:         valueDecimal,
					BlockTimestamp:       blockTimestamp,
					TokenContractName:    tokenContractName,
					TokenContractSymbol:  tokenContractSymbol,
					TransactionFee:       transactionFee,
				}

				tokenTransfers = append(tokenTransfers, tokenTransfer)
			}
		}
	}

	return tokenTransfers
}
