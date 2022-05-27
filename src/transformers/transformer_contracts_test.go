package transformers

import (
	"github.com/sudoblockio/icon-transformer/models"
	"testing"

	"github.com/sudoblockio/icon-transformer/config"
)

func TestProcessContracts(t *testing.T) {
	//config.ReadTestEnvironment()
	config.ReadEnvironment()

	contract := models.ContractProcessed{
		Address:              "testing",
		Name:                 "testing",
		Status:               "testing",
		CreatedTimestamp:     1,
		IsToken:              true,
		ContractUpdatedBlock: 50,
	}

	transformContractsToLoadAddress(&contract)
}
