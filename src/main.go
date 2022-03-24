package main

import (
	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/global"
	"github.com/sudoblockio/icon-transformer/kafka"
	"github.com/sudoblockio/icon-transformer/logging"
	"github.com/sudoblockio/icon-transformer/routines"
	"github.com/sudoblockio/icon-transformer/transformers"
)

func main() {
	config.ReadEnvironment()
	logging.Init()

	// Feature flags
	if config.Config.RoutinesRunOnly == true {
		// Start routines
		routines.StartRoutines()

		global.WaitShutdownSig()
	}

	kafka.StartConsumers()

	transformers.Start()

	global.WaitShutdownSig()
}
