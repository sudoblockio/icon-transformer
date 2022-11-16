package routines

import (
	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/models"
	"github.com/sudoblockio/icon-transformer/redis"
	"go.uber.org/zap"
)

func GetAddressTxCounts(address *models.Address) {
	// Regular Tx Count
	countRegular, err := crud.GetTransactionByAddressCrud().CountWhere("address", address.Address)
	if err != nil {
		zap.S().Info("Routine=RegularTxAddressCount - ERROR: ", err.Error())
	} else {
		address.TransactionCount = countRegular
		err = redis.GetRedisClient().SetCount(
			config.Config.RedisKeyPrefix+"transaction_regular_count_by_address_"+address.Address,
			countRegular,
		)
		if err != nil {
			zap.S().Info("Routine=TxCount, Address=", address.Address, " - Error: ", err.Error())
		}
	}

	// Regular Tx Count
	countIcx, err := crud.GetTransactionCrud().CountTransactionIcxByAddress(address.Address)
	if err != nil {
		zap.S().Info("Routine=IcxTxAddressCount - ERROR: ", err.Error())
	} else {
		err = redis.GetRedisClient().SetCount(
			config.Config.RedisKeyPrefix+"transaction_icx_count_by_address_"+address.Address,
			countIcx,
		)
		if err != nil {
			zap.S().Info("Routine=TxCount, Address=", address.Address, " - Error: ", err.Error())
		}
	}

	// Internal Tx Count
	countInternal, err := crud.GetTransactionCrud().CountTransactionsInternalByAddress(address.Address)
	if err != nil {
		zap.S().Info("Routine=InternalTxAddressCount - ERROR: ", err.Error())
	} else {
		address.TransactionInternalCount = countInternal
		err = redis.GetRedisClient().SetCount(
			config.Config.RedisKeyPrefix+"transaction_internal_count_by_address_"+address.Address,
			countInternal,
		)
		if err != nil {
			zap.S().Info("Routine=TxInternalCount, Address=", address.Address, " - Error: ", err.Error())
		}
	}

	// Token Transfer Count
	countTokenTx, err := crud.GetTokenTransferCrud().CountByAddress(address.Address)
	if err != nil {
		zap.S().Info("Routine=TokenTxCount, Address=", address.Address, " - Error: ", err.Error())
	} else {
		address.TokenTransferCount = countTokenTx
		err = redis.GetRedisClient().SetCount(
			config.Config.RedisKeyPrefix+"token_transfer_count_by_address_"+address.Address,
			countTokenTx,
		)
		if err != nil {
			zap.S().Info("Routine=TokenTransfer, Address=", address.Address, " - Error: ", err.Error())
		}
	}
	if address.IsToken {
		countTokenTx, err := crud.GetTokenTransferCrud().CountWhere("token_contract_address", address.Address)
		if err != nil {
			zap.S().Info("Routine=TokenTxCount, Address=", address.Address, " - Error: ", err.Error())
		} else {
			address.TokenTransferCount = countTokenTx
			err = redis.GetRedisClient().SetCount(
				config.Config.RedisKeyPrefix+"token_transfer_count_by_token_contract_"+address.Address,
				countTokenTx,
			)
			if err != nil {
				zap.S().Info("Routine=TokenTransfer, Address=", address.Address, " - Error: ", err.Error())
			}
		}
	}

	// Log Count
	countLog, err := crud.GetLogCrud().CountWhere("address", address.Address)
	if err != nil {
		zap.S().Info("Routine=LogCount, Address=", address.Address, " - Error: ", err.Error())
	} else {
		address.LogCount = countLog
		err = redis.GetRedisClient().SetCount(
			config.Config.RedisKeyPrefix+"log_count_by_address_"+address.Address,
			countLog,
		)
		if err != nil {
			zap.S().Info("Routine=LogCount, Address=", address.Address, " - Error: ", err.Error())
		}
	}
}

func SetAddressTxCounts(address *models.Address) {
	GetAddressTxCounts(address)
	err := crud.GetAddressRoutineCruds()["counts"].UpsertOne(address)
	if err != nil {
		zap.S().Info(err)
	}
}
