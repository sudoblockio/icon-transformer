package routines

import (
	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/models"
	"github.com/sudoblockio/icon-transformer/redis"
	"go.uber.org/zap"
	"time"
)

func setAddressTxCounts(address *models.Address) {
	// Regular Tx Count
	countRegular, err := crud.GetTransactionCrud().CountRegularByAddress(address.Address)
	if err != nil {
		// Try again
		zap.S().Warn("Routine=AddressCount - ERROR: ", err.Error())
		time.Sleep(1 * time.Second)
		return
	}
	address.TransactionCount = countRegular
	err = redis.GetRedisClient().SetCount(
		config.Config.RedisKeyPrefix+"transaction_regular_count_by_address_"+address.Address,
		countRegular,
	)
	if err != nil {
		zap.S().Warn("Routine=TxCount, Address=", address.Address, " - Error: ", err.Error())
		time.Sleep(1 * time.Second)
		return
	}

	// Internal Tx Count
	countInternal, err := crud.GetTransactionCrud().CountInternalByAddress(address.Address)
	if err != nil {
		zap.S().Warn("Routine=InternalTxCount, Address=", address.Address, " - Error: ", err.Error())
		return
	}
	address.TransactionInternalCount = countInternal
	err = redis.GetRedisClient().SetCount(
		config.Config.RedisKeyPrefix+"transaction_internal_count_by_address_"+address.Address,
		countInternal,
	)
	if err != nil {
		zap.S().Warn("Routine=TxInternalCount, Address=", address.Address, " - Error: ", err.Error())
		return
	}

	// Token Transfer Count
	countTokenTx, err := crud.GetTokenTransferCrud().CountByAddress(address.Address)
	if err != nil {
		zap.S().Warn("Routine=InternalTxCount, Address=", address.Address, " - Error: ", err.Error())
		return
	}
	address.TokenTransferCount = countTokenTx
	err = redis.GetRedisClient().SetCount(
		config.Config.RedisKeyPrefix+"token_transfer_count_by_address_"+address.Address,
		countTokenTx,
	)
	if err != nil {
		zap.S().Warn("Routine=TokenTransfer, Address=", address.Address, " - Error: ", err.Error())
		return
	}

	// Log Count
	countLog, err := crud.GetLogCrud().CountLogsByAddress(address.Address)
	if err != nil {
		zap.S().Warn("Routine=InternalTxCount, Address=", address.Address, " - Error: ", err.Error())
		return
	}
	address.TransactionInternalCount = countLog
	err = redis.GetRedisClient().SetCount(
		config.Config.RedisKeyPrefix+"transaction_regular_count_by_address_"+address.Address,
		countLog,
	)
	if err != nil {
		zap.S().Warn("Routine=LogCount, Address=", address.Address, " - Error: ", err.Error())
		return
	}

	crud.GetAddressCrud().UpsertOneCols(address, []string{
		"address",
		"transaction_count",
		"transaction_internal_count",
		"token_transfer_count",
		"log_count",
	})
}