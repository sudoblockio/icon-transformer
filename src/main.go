package main

import (
	"github.com/sudoblockio/icon-go-worker/config"
	"github.com/sudoblockio/icon-go-worker/global"
	"github.com/sudoblockio/icon-go-worker/kafka"
	"github.com/sudoblockio/icon-go-worker/logging"
	"github.com/sudoblockio/icon-go-worker/transformers"
)

func main() {
	config.ReadEnvironment()
	logging.Init()

	kafka.StartConsumers()

	transformers.Start()

	global.WaitShutdownSig()
}
