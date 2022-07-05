package transformers

import "github.com/sudoblockio/icon-transformer/config"

func Start() {
	// Blocks Topic
	if config.Config.KafkaBlocksTopic != "" {
		go startBlocks()
	}

	// Contracts Topic
	if config.Config.KafkaContractsTopic != "" {
		go startContracts()
	}

	// Dead Message Topic
	if config.Config.KafkaDeadMessageTopic != "" {
		go startDeadMessages()
	}
}
