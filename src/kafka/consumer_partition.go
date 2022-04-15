package kafka

import (
	"github.com/Shopify/sarama"
	"go.uber.org/zap"
	"time"
)

////////////////////////
// Partition Consumer //
////////////////////////
func (k *kafkaTopicConsumer) consumePartition(topic string, partition int, startOffset int) {
	saramaConfig := sarama.NewConfig()

	////////////////////
	// Wait for topic //
	////////////////////
	admin, err := getAdmin(k.brokerURL, saramaConfig)
	if err != nil {
		zap.S().Fatal("KAFKA ADMIN ERROR: ", err.Error())
	}
	defer func() { _ = admin.Close() }()

	for {
		allTopicsCreated := true
		topics, _ := admin.ListTopics()

		for _, topicName := range k.topicNames {
			if _, ok := topics[topicName]; ok == false {
				// Topic not created
				allTopicsCreated = false
				break
			}
		}

		if allTopicsCreated {
			break
		}

		zap.S().Info("KAFKA ADMIN WARN: topics not created yet...sleeping")
		time.Sleep(1 * time.Second)
	}

	///////////////////////////
	// Consumer Group Config //
	///////////////////////////

	// Version
	version, err := sarama.ParseKafkaVersion("2.1.1")
	if err != nil {
		zap.S().Panic("CONSUME GROUP ERROR: parsing Kafka version: ", err.Error())
	}
	saramaConfig.Version = version

	// Initial Offset
	saramaConfig.Consumer.Offsets.Initial = sarama.OffsetNewest

	var consumer sarama.Consumer
	for {
		consumer, err = sarama.NewConsumer([]string{k.brokerURL}, saramaConfig)
		if err != nil {
			zap.S().Info("Creating consumer err: ", err.Error())
			zap.S().Info("Retrying in 3 seconds...")
			time.Sleep(3 * time.Second)
			continue
		}
		break
	}

	// Clean up
	defer func() {
		if err := consumer.Close(); err != nil {
			zap.S().Panic("KAFKA CONSUMER CLOSE PANIC: ", err.Error())
		}
	}()

	// Connect to partition
	pc, err := consumer.ConsumePartition(topic, int32(partition), int64(startOffset))

	if err != nil {
		zap.S().Panic("KAFKA CONSUMER PARTITIONS PANIC: ", err.Error())
	}
	if pc == nil {
		zap.S().Panic("KAFKA CONSUMER PARTITIONS PANIC: Failed to create PartitionConsumer")
	}

	// Read partition
	for {
		var topic_msg *sarama.ConsumerMessage
		select {
		case msg := <-pc.Messages():
			topic_msg = msg
		case consumerErr := <-pc.Errors():
			zap.S().Info("KAFKA PARTITION CONSUMER ERROR:", consumerErr.Err)
			continue
		case <-time.After(5 * time.Second):
			zap.S().Debug("Consumer ", topic, ": No new kafka messages, waited 5 secs")
			continue
		}
		zap.S().Debug("Consumer ", topic, ": Consumed message key=", string(topic_msg.Key))

		// Broadcast
		k.TopicChannels[topic] <- topic_msg

		zap.S().Debug("Consumer ", topic, ": Broadcasted message key=", string(topic_msg.Key))
	}
}
