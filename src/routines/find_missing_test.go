package routines

import (
	"github.com/sudoblockio/icon-transformer/config"
	"testing"
)

func TestFindMissing(t *testing.T) {
	config.ReadTestEnvironment()
	findMissingBlocks()
}
