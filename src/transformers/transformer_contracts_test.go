package transformers

import (
	"github.com/sudoblockio/icon-transformer/models"
	"testing"

	"github.com/sudoblockio/icon-transformer/config"
)

func TestProcessContracts(t *testing.T) {
	config.ReadTestEnvironment()

	contract := models.ContractProcessed{
		Address:          "testing",
		Name:             "testing",
		CreatedTimestamp: 1,
		Status:           "testing",
		IsToken:          true,
	}

	transformContractsToLoadAddress(&contract)
}
