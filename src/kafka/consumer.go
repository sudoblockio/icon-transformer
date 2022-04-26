package kafka

import (
	"time"

	"github.com/Shopify/sarama"
	"github.com/cenkalti/backoff"
	"github.com/sudoblockio/icon-transformer/config"
	"go.uber.org/zap"
)

//main struct
type kafkaTopicConsumer struct {
	brokerURL     string
	topicNames    []string
	TopicChannels map[string]chan *sarama.ConsumerMessage
}

var KafkaTopicConsumer *kafkaTopicConsumer

func StartConsumers() {

	// Init topic names
	// NOTE add new kafka topics here
	topicNames := []string{
		config.Config.KafkaBlocksTopic,
		config.Config.KafkaContractsTopic,
		//config.Config.KafkaDeadMessageTopic,
	}

	// Init topic channels
	topicChannels := make(map[string]chan *sarama.ConsumerMessage)
	for _, topicName := range topicNames {
		topicChannels[topicName] = make(chan *sarama.ConsumerMessage)
	}

	// Init consumer
	KafkaTopicConsumer = &kafkaTopicConsumer{
		brokerURL:     config.Config.KafkaBrokerURL,
		topicNames:    topicNames,
		TopicChannels: topicChannels,
	}

	////////////////////
	// Consumer Modes //
	////////////////////

	// Partition
	if config.Config.ConsumerIsPartitionConsumer == true {
		zap.S().Info(
			"kafkaBroker=", config.Config.KafkaBrokerURL,
			" ConsumerPartitionTopic=", config.Config.ConsumerPartitionTopic,
			" ConsumerPartition=", config.Config.ConsumerPartition,
			" ConsumerPartitionStartOffset=", config.Config.ConsumerPartitionStartOffset,
			" - Starting Consumers")
		go KafkaTopicConsumer.consumePartition(
			config.Config.ConsumerPartitionTopic,
			config.Config.ConsumerPartition,
			config.Config.ConsumerPartitionStartOffset,
		)
		return
	}

	// Tail
	if config.Config.ConsumerIsTail == true {
		zap.S().Info(
			"kafkaBroker=", config.Config.KafkaBrokerURL,
			" consumerTopics=", topicNames,
			" consumerGroup=", config.Config.ConsumerGroup+"-"+config.Config.ConsumerJobID,
			" - Starting Consumers")
		go KafkaTopicConsumer.consumeGroup(config.Config.ConsumerGroup + "-" + config.Config.ConsumerJobID)
		return
	}

	// Head
	// Default
	zap.S().Info(
		"kafkaBroker=", config.Config.KafkaBrokerURL,
		" consumerTopics=", topicNames,
		" consumerGroup=", config.Config.ConsumerGroup+"-head",
		" - Starting Consumers")
	go KafkaTopicConsumer.consumeGroup(config.Config.ConsumerGroup + "-head")
	return
}

func getAdmin(brokerURL string, saramaConfig *sarama.Config) (sarama.ClusterAdmin, error) {
	var admin sarama.ClusterAdmin
	operation := func() error {
		a, err := sarama.NewClusterAdmin([]string{brokerURL}, saramaConfig)
		if err != nil {
			zap.S().Info("KAFKA ADMIN NEWCLUSTERADMIN WARN: ", err.Error())
		} else {
			admin = a
		}
		return err
	}
	neb := backoff.NewConstantBackOff(time.Second)
	err := backoff.Retry(operation, neb)
	return admin, err
}
