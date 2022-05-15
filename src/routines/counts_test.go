package routines

import (
	"github.com/stretchr/testify/require"
	"github.com/sudoblockio/icon-transformer/config"
	"testing"
)

func TestBlockCount(t *testing.T) {
	config.ReadTestEnvironment()
	err := blockCountExec()
	require.Nil(t, err)
}

func TestTransactionRegularCount(t *testing.T) {
	config.ReadTestEnvironment()
	err := transactionRegularCountExec()
	require.Nil(t, err)
}

// TODO: Currently not used
//func TestTransactionInternalCount(t *testing.T) {
//	config.ReadTestEnvironment()
//	err := transactionRegularCountExec()
//	require.Nil(t, err)
//}

func TestTokenTransferCount(t *testing.T) {
	config.ReadTestEnvironment()
	err := tokenTransferCountExec()
	require.Nil(t, err)
}

//func TestCountByAddress(t *testing.T) {
//	config.ReadTestEnvironment()
//	err := byAddressCountExec()
//	require.Nil(t, err)
//}

//func TestCountByTokenContract(t *testing.T) {
//	config.ReadTestEnvironment()
//	err := byTokenContractCountExec()
//	require.Nil(t, err)
//}

//func TestStartRedisRecovery(t *testing.T) {
//	config.ReadTestEnvironment()
//	StartRedisRecovery()
//}
