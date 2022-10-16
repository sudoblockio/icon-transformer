package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sudoblockio/icon-transformer/config"
	"testing"
)

func TestIconNodeServiceGetBlockTransactionHashes(t *testing.T) {
	config.ReadEnvironment()
	body, err := IconNodeServiceGetBlockTransactionHashes(1)

	bodyVal := *body
	expected := []string{"0x375540830d475a73b704cf8dee9fa9eba2798f9d2af1fa55a85482e48daefd3b"}
	assert.Equal(t, bodyVal, expected)

	require.Nil(t, err)
	require.NotEmpty(t, body)
}

func TestIconNodeServiceGetTokenDecimalBase(t *testing.T) {
	config.ReadEnvironment()
	body, err := IconNodeServiceGetTokenDecimalBase("cx993810b4523ab6b1658925d8d6c234f286adbdba")

	require.Nil(t, err)
	require.NotEmpty(t, body)
}

func TestIconNodeServiceGetTokenContractName(t *testing.T) {
	config.ReadEnvironment()
	body, err := IconNodeServiceGetTokenContractName("cx993810b4523ab6b1658925d8d6c234f286adbdba")

	require.Nil(t, err)
	require.NotEmpty(t, body)
}

func TestIconNodeServiceGetTokenContractSymbol(t *testing.T) {
	config.ReadEnvironment()
	body, err := IconNodeServiceGetTokenContractSymbol("cx993810b4523ab6b1658925d8d6c234f286adbdba")

	require.Nil(t, err)
	require.NotEmpty(t, body)
}

func TestIconNodeServiceGetBalance(t *testing.T) {
	config.ReadEnvironment()
	body, err := IconNodeServiceGetBalance("cx993810b4523ab6b1658925d8d6c234f286adbdba")

	require.Nil(t, err)
	require.NotEmpty(t, body)
}

func TestIconNodeServiceGetBalanceGovAddress(t *testing.T) {
	config.ReadEnvironment()
	body, err := IconNodeServiceGetBalance("hx0000000000000000000000000000000000000000")

	require.Nil(t, err)
	require.NotEmpty(t, body)
}

func TestIconNodeServiceGetStakedBalance(t *testing.T) {
	config.ReadEnvironment()
	body, err := IconNodeServiceGetStakedBalance("cx993810b4523ab6b1658925d8d6c234f286adbdba")

	require.Nil(t, err)
	require.NotEmpty(t, body)
}

func TestIconNodeServiceGetTokenBalance(t *testing.T) {
	config.ReadEnvironment()
	body, err := IconNodeServiceGetTokenBalance(
		"cx993810b4523ab6b1658925d8d6c234f286adbdba",
		"hx79dec003161ca695637d7d02143c07ac72cd3018",
	)

	require.Nil(t, err)
	require.NotEmpty(t, body)
}

func TestIconNodeServiceGetPreps(t *testing.T) {
	config.ReadEnvironment()
	body, err := IconNodeServiceGetPreps()

	require.Nil(t, err)
	require.NotEmpty(t, body)
}

func TestIconNodeServiceGetTransactionResult(t *testing.T) {
	config.ReadEnvironment()
	result, err := IconNodeServiceGetTransactionResult("0x362b4a1f81d3b3f505ba93a4dbd65527d8ba21938d864fa377a07d3de8460401")

	scoreAddress, ok := result["scoreAddress"].(string)

	require.Equal(t, ok, true)
	assert.Equal(t, scoreAddress, "cxf81989e82ebfc3b69c758f8f3017822d5dc6ab46")
	require.Nil(t, err)
	require.NotEmpty(t, result)
}

func TestIconNodeServiceGetScoreAddressFromTransactionResult(t *testing.T) {
	config.ReadEnvironment()
	result, err := IconNodeServiceGetTransactionResult(
		"0x85fff93f669f778254d8a4e484683ccb457b0e0c0d6ec61e410401c13bb0c162",
	)
	require.Nil(t, err)
	require.NotEmpty(t, result)
}

func TestIconNodeServiceGetScoreStatus(t *testing.T) {
	config.ReadEnvironment()
	config.Config.IconNodeServiceURL = []string{"https://api.icon.community/api/v3"}
	result, err := IconNodeServiceGetScoreStatus(
		"cx203d9cd2a669be67177e997b8948ce2c35caffae",
	)
	require.Nil(t, err)
	require.NotEmpty(t, result)
}

func TestIconNodeServiceCallContractMethod(t *testing.T) {
	config.ReadEnvironment()
	config.Config.IconNodeServiceURL = []string{"https://api.icon.community/api/v3"}
	result, err := IconNodeServiceCallContractMethod(
		"cx077807f2322aeb42ea19a1fcc0c9f3d3f35e1461", "symbol",
	)
	require.Nil(t, err)
	assert.Equal(t, result, "BNB")
}
