package main

import (
	"github.com/sudoblockio/icon-go-worker/config"
	"github.com/sudoblockio/icon-go-worker/global"
	"github.com/sudoblockio/icon-go-worker/kafka"
	"github.com/sudoblockio/icon-go-worker/logging"
	"github.com/sudoblockio/icon-go-worker/routines"
	"github.com/sudoblockio/icon-go-worker/transformers"
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
