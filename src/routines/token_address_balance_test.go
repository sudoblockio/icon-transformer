package routines

import (
	"github.com/stretchr/testify/assert"
	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/models"
	"testing"
	"time"
)

func TestSetTokenAddressBalances(t *testing.T) {
	config.ReadEnvironment()
	tokenAddress := &models.TokenAddress{
		TokenContractAddress: "cxcfe9d1f83fa871e903008471cca786662437e58d",
		Address:              "hx42c7d33652beabb87eae2a09f67c53b1927b2a96",
	}
	setTokenAddressBalances(tokenAddress)

	assert.NotNil(t, tokenAddress.Balance)
	time.Sleep(5 * time.Second)
}
