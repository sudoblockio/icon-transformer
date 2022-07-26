package routines

import (
	"github.com/stretchr/testify/assert"
	"github.com/sudoblockio/icon-transformer/models"
	"testing"

	"github.com/sudoblockio/icon-transformer/config"
)

func TestAddressBalance(t *testing.T) {
	config.ReadEnvironment()
	address := &models.Address{Address: "hx562dc1e2c7897432c298115bc7fbcc3b9d5df294"}
	getAddressBalances(address)

	assert.NotNil(t, address.Balance)
}
