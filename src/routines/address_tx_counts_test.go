package routines

import (
	"github.com/stretchr/testify/assert"
	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/models"
	"testing"
)

func TestGetAddressTxCounts(t *testing.T) {
	config.ReadEnvironment()
	address := &models.Address{
		Address: "hxc1481b2459afdbbde302ab528665b8603f7014dc",
	}
	GetAddressTxCounts(address)

	assert.NotNil(t, address)
}
