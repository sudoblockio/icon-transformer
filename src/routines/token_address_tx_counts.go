package routines

import (
	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/models"
	"github.com/sudoblockio/icon-transformer/redis"
	"go.uber.org/zap"

	"github.com/sudoblockio/icon-transformer/crud"
)

func GetTokenAddressTxCounts(address *models.Address) *models.Address {

	// Token Transfer Count
	countTokenTx, err := crud.GetTokenTransferCrud().CountByAddress(address.Address)
	if err != nil {
		zap.S().Warn("Routine=InternalTxCount, Address=", address.Address, " - Error: ", err.Error())
		return nil
	}
	address.TokenTransferCount = countTokenTx
	err = redis.GetRedisClient().SetCount(
		config.Config.RedisKeyPrefix+"token_transfer_count_by_token_contract_"+address.Address,
		countTokenTx,
	)
	if err != nil {
		zap.S().Warn("Routine=TokenTransfer, Address=", address.Address, " - Error: ", err.Error())
		return nil
	}

	return address
}

func setTokenAddressTxCounts(address *models.Address) {
	addressNew := GetTokenAddressTxCounts(address)
	if address != nil {
		//crud.GetAddressRoutineCruds()["counts"].LoaderChannel <- addressNew
		err := crud.GetAddressRoutineCruds()["counts"].UpsertOne(addressNew)
		if err != nil {
			zap.S().Fatal(err.Error())
		}
	}
}
