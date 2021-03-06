package routines

import (
	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/models"
	"github.com/sudoblockio/icon-transformer/redis"
	"go.uber.org/zap"
	"time"
)

func getAddressTxCounts(address *models.Address) *models.Address {
	// Regular Tx Count
	countRegular, err := crud.GetTransactionByAddressCrud().CountByAddress(address.Address)
	if err != nil {
		// Try again
		zap.S().Warn("Routine=AddressCount - ERROR: ", err.Error())
		time.Sleep(1 * time.Second)
		return nil
	}
	address.TransactionCount = countRegular
	err = redis.GetRedisClient().SetCount(
		config.Config.RedisKeyPrefix+"transaction_regular_count_by_address_"+address.Address,
		countRegular,
	)
	if err != nil {
		zap.S().Warn("Routine=TxCount, Address=", address.Address, " - Error: ", err.Error())
		time.Sleep(1 * time.Second)
		return nil
	}

	// Internal Tx Count
	countInternal, err := crud.GetTransactionCrud().CountInternalByAddress(address.Address)
	if err != nil {
		zap.S().Warn("Routine=InternalTxCount, Address=", address.Address, " - Error: ", err.Error())
		return nil
	}
	address.TransactionInternalCount = countInternal
	err = redis.GetRedisClient().SetCount(
		config.Config.RedisKeyPrefix+"transaction_internal_count_by_address_"+address.Address,
		countInternal,
	)
	if err != nil {
		zap.S().Warn("Routine=TxInternalCount, Address=", address.Address, " - Error: ", err.Error())
		return nil
	}

	// Token Transfer Count
	countTokenTx, err := crud.GetTokenTransferCrud().CountByAddress(address.Address)
	if err != nil {
		zap.S().Warn("Routine=InternalTxCount, Address=", address.Address, " - Error: ", err.Error())
		return nil
	}
	address.TokenTransferCount = countTokenTx
	err = redis.GetRedisClient().SetCount(
		config.Config.RedisKeyPrefix+"token_transfer_count_by_address_"+address.Address,
		countTokenTx,
	)
	if err != nil {
		zap.S().Warn("Routine=TokenTransfer, Address=", address.Address, " - Error: ", err.Error())
		return nil
	}

	// Log Count
	countLog, err := crud.GetLogCrud().CountLogsByAddress(address.Address)
	if err != nil {
		zap.S().Warn("Routine=InternalTxCount, Address=", address.Address, " - Error: ", err.Error())
		return nil
	}
	address.TransactionInternalCount = countLog
	err = redis.GetRedisClient().SetCount(
		config.Config.RedisKeyPrefix+"log_count_by_address_"+address.Address,
		countLog,
	)
	if err != nil {
		zap.S().Warn("Routine=LogCount, Address=", address.Address, " - Error: ", err.Error())
		return nil
	}

	return address
}

func setAddressTxCounts(address *models.Address) {
	address = getAddressTxCounts(address)
	if address != nil {
		crud.GetAddressCrud().UpsertOneCols(address, []string{
			"address",
			"transaction_count",
			"transaction_internal_count",
			"token_transfer_count",
			"log_count",
		})
	}
}
