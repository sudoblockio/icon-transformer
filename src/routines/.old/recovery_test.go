package _old

import (
	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/routines"
	"testing"
)

func TestRedisRecovery(t *testing.T) {
	config.ReadTestEnvironment()
	routines.StartRecovery()
}
