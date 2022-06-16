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
	//config.ReadTestEnvironment()
	logging.Init()

	// Feature flags
	if config.Config.RoutinesRunOnly == true {
		// Start cron
		routines.CronStart()

		global.WaitShutdownSig()
	} else if config.Config.FindMissingRunOnly == true {
		// Start find only
		routines.StartFindMissing()

		global.WaitShutdownSig()
	} else if config.Config.RedisRecoveryRunOnly == true {
		// Start redis recovery
		routines.StartRecovery()

		global.WaitShutdownSig()
	}

	kafka.StartConsumers()

	transformers.Start()

	global.WaitShutdownSig()
}
