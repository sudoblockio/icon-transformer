package transformers

import (
	"github.com/Shopify/sarama"
	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/kafka"
	"github.com/sudoblockio/icon-transformer/models"
	"go.uber.org/zap"
)

func startDeadMessages() {
	kafkaDeadMessageTopic := config.Config.KafkaDeadMessageTopic
	kafkaDeadMessageTopicChannel := kafka.KafkaTopicConsumer.TopicChannels[kafkaDeadMessageTopic]

	loaderChannel := crud.GetDeadBlockCrud().LoaderChannel
	zap.S().Debug("DeadMessage transformer: started working")
	for {
		deadMessage := <-kafkaDeadMessageTopicChannel

		//deadBlock := transformDeadMessageToDeadBlock(deadMessage)

		deadBlock := &models.DeadBlock{
			Topic:     deadMessage.Topic,
			Partition: int64(deadMessage.Partition),
			Offset:    deadMessage.Offset,
			Key:       string(deadMessage.Key),
			Value:     string(deadMessage.Value),
		}

		loaderChannel <- deadBlock
	}
}

// Dead block loader
func transformDeadMessageToLoadDeadBlock(deadMessage *sarama.ConsumerMessage) {
	loaderChannel := crud.GetDeadBlockCrud().LoaderChannel

	deadBlock := transformDeadMessageToDeadBlock(deadMessage)
	loaderChannel <- deadBlock
}
