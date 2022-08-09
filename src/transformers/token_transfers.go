package transformers

import (
	"fmt"
	"math/big"

	"github.com/sudoblockio/icon-transformer/models"
	"github.com/sudoblockio/icon-transformer/service"
	"github.com/sudoblockio/icon-transformer/utils"
	"go.uber.org/zap"
)

// Find the highest value for the ICX amount for the transfer
func findHighestIcxTransfer(logs []*models.LogETL) (string, float64) {

	var highestIcxTransfer float64 = 0
	var highestIcxTransferValue = "0x0"
	for _, v := range logs {
		if v.Indexed[0] == "ICXTransfer(Address,Address,int)" {
			icxValue := v.Indexed[3]
			icxAmount := utils.StringHexToFloat64(icxValue, 18)

			if icxAmount > highestIcxTransfer {
				highestIcxTransfer = icxAmount
				highestIcxTransferValue = icxValue
			}
		}
	}
	return highestIcxTransferValue, highestIcxTransfer
}

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
			} else if logETL.Indexed[0] == "Transfer(Address,Address,int)" && len(logETL.Indexed) == 4 {
				// Transfer is not a protected method
				// Handle IRC3 transfers

				// Token Contract Address
				tokenContractAddress := logETL.Address

				// From Address
				fromAddress := logETL.Indexed[1]

				// To Address
				toAddress := logETL.Indexed[2]

				// Value
				value, valueDecimal := findHighestIcxTransfer(transactionETL.Logs)

				// NFT ID
				nftId := utils.StringHexToInt64(logETL.Indexed[3])

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
					// IRC3 has symbol in spec. IRC31 does not
					tokenContractSymbol = ""
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
					NftId:                nftId,
				}

				tokenTransfers = append(tokenTransfers, tokenTransfer)
			} else if logETL.Indexed[0] == "TransferSingle(Address,Address,Address,int,int)" && len(logETL.Indexed) == 4 {
				// TransferSingle is not a protected method
				// Handle IRC31 transfers single

				// Token Contract Address
				tokenContractAddress := logETL.Address

				// From Address
				fromAddress := logETL.Indexed[1]

				// To Address
				toAddress := logETL.Indexed[2]

				// Value
				value, valueDecimal := findHighestIcxTransfer(transactionETL.Logs)

				// NFT ID
				var nftId int64
				if len(logETL.Data) == 2 {
					nftId = utils.StringHexToInt64(logETL.Data[0])
				} else {
					continue
				}

				// Block Timestamp
				blockTimestamp := blockETL.Timestamp

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
					TransactionFee:       transactionFee,
					NftId:                nftId,
				}

				tokenTransfers = append(tokenTransfers, tokenTransfer)
			}
		}
	}

	return tokenTransfers
}
