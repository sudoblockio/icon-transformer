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

func TestIconNodeServiceGetStakedBalance(t *testing.T) {
	config.ReadEnvironment()
	body, err := IconNodeServiceGetStakedBalance("cx993810b4523ab6b1658925d8d6c234f286adbdba")

	require.Nil(t, err)
	require.NotEmpty(t, body)
}

func TestIconNodeServiceGetTokenBalance(t *testing.T) {
	config.ReadEnvironment()
	body, err := IconNodeServiceGetTokenBalance("cx993810b4523ab6b1658925d8d6c234f286adbdba", "hx79dec003161ca695637d7d02143c07ac72cd3018")

	require.Nil(t, err)
	require.NotEmpty(t, body)
}

func TestIconNodeServiceGetPreps(t *testing.T) {
	config.ReadEnvironment()
	body, err := IconNodeServiceGetPreps()

	require.Nil(t, err)
	require.NotEmpty(t, body)
}
