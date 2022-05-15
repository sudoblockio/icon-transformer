//
package routines

import (
	"github.com/stretchr/testify/require"
	"github.com/sudoblockio/icon-transformer/config"
	"testing"
)

func TestAddressTransactionCount(t *testing.T) {
	config.ReadTestEnvironment()
	err := addressTransactionCountExec()
	require.Nil(t, err)
}
