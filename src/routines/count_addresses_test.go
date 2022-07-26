package routines

import (
	"github.com/sudoblockio/icon-transformer/config"
	"testing"
)

func TestCountAddresses(t *testing.T) {
	config.ReadEnvironment()
	countAddressesToRedisRoutine()
}
