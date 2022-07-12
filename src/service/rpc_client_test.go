package service

import (
	"github.com/stretchr/testify/require"
	"github.com/sudoblockio/icon-transformer/config"
	"testing"
)

var getLastBlockPayload string = `{
    "jsonrpc": "2.0",
    "method": "icx_getLastBlock",
    "id": 1234
	}`

var getBalance string = `{
	"jsonrpc": "2.0",
    "method": "icx_getBalance",
    "id": 1234,
    "params": {
         "address": "hx562dc1e2c7897432c298115bc7fbcc3b9d5df294"}
    }`

func updateConfig() {
	config.Config.IconNodeServiceURL = []string{
		"http://funky.wrong",
		"https://api.icon.community/api/v3",
	}
}

func TestJsonRpcRequestWithRetry(t *testing.T) {
	config.ReadEnvironment()
	updateConfig()
	body, err := JsonRpcRequestWithRetry(getBalance)
	require.Nil(t, err)
	require.NotEmpty(t, body)
}

func TestJsonRpcRequestWithBackup(t *testing.T) {
	config.ReadEnvironment()
	updateConfig()
	body, err := JsonRpcRequestWithBackup(getLastBlockPayload)
	require.Nil(t, err)
	require.NotEmpty(t, body)
}

func TestJsonRpcRequestError(t *testing.T) {
	config.ReadEnvironment()
	updateConfig()
	_, err := JsonRpcRequest(getLastBlockPayload, "http://funky.wrong")
	require.NotEmpty(t, err)
}
