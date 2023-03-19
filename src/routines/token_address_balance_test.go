package routines

import (
	"github.com/stretchr/testify/assert"
	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/models"
	"testing"
)

func TestSetTokenAddressBalances(t *testing.T) {
	config.ReadEnvironment()
	tokenAddress := &models.TokenAddress{
		TokenContractAddress: "cxcfe9d1f83fa871e903008471cca786662437e58d",
		Address:              "hxb41775a05572c421917b6d5d80fd5f31c495b7f8",
	}
	setTokenAddressBalances(tokenAddress)

	assert.NotNil(t, tokenAddress.Balance)
}
