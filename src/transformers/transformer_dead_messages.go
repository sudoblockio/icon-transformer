package transformers

import (
	"github.com/Shopify/sarama"
	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/kafka"
	"go.uber.org/zap"
)

func startDeadMessages() {
	kafkaDeadMessageTopic := config.Config.KafkaDeadMessageTopic

	// Input channels
	kafkaDeadMessageTopicChannel := kafka.KafkaTopicConsumer.TopicChannels[kafkaDeadMessageTopic]

	zap.S().Debug("DeadMessage transformer: started working")
	for {

		///////////////////
		// Kafka Message //
		///////////////////

		deadMessage := <-kafkaDeadMessageTopicChannel

		/////////////
		// Loaders //
		/////////////

		// Dead block loader
		transformDeadMessageToLoadDeadBlock(deadMessage)
	}
}

// Dead block loader
func transformDeadMessageToLoadDeadBlock(deadMessage *sarama.ConsumerMessage) {
	loaderChannel := crud.GetDeadBlockCrud().LoaderChannel

	deadBlock := transformDeadMessageToDeadBlock(deadMessage)
	loaderChannel <- deadBlock
}
