package routines

import (
	"github.com/sudoblockio/icon-transformer/transformers"
	"go.uber.org/zap"

	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/models"
)

func setAddressContractMeta(address *models.Address) {
	if !address.IsContract {
		return
	}

	transformers.EnrichContractsMeta(address)

	if address != nil {
		err := crud.GetAddressRoutineCruds()["contract_meta"].UpsertOne(address)
		if err != nil {
			zap.S().Fatal(err.Error())
		}
	}
}
