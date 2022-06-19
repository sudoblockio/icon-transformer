package routines

import (
	"github.com/stretchr/testify/assert"
	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/models"
	"testing"
)

func TestGetAddressTxCounts(t *testing.T) {
	config.ReadTestEnvironment()
	address := &models.Address{
		Address: "hx68646780e14ee9097085f7280ab137c3633b4b5f",
	}
	getAddressTxCounts(address)

	assert.NotNil(t, address)
}
