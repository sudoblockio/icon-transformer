package kafka

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/Shopify/sarama"
	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

////////////////////
// Group Consumer //
////////////////////
func (k *kafkaTopicConsumer) consumeGroup(group string) {
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
		missingTopics := []string{}

		for _, topicName := range k.topicNames {

			// NOTE remove when contracts service is fixed
			if topicName == config.Config.KafkaContractsTopic {
				continue
			}

			if _, ok := topics[topicName]; ok == false {
				// Topic not created
				allTopicsCreated = false
				missingTopics = append(missingTopics, topicName)
				break
			}
		}

		if allTopicsCreated {
			break
		}

		zap.S().Info("KAFKA ADMIN WARN: topics not created yet, topics=", missingTopics)
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
	saramaConfig.Consumer.Offsets.Initial = sarama.OffsetOldest

	// Config values
	saramaConfig.Consumer.MaxProcessingTime = 2 * time.Second
	saramaConfig.Consumer.Group.Session.Timeout = 20 * time.Second
	saramaConfig.Net.KeepAlive = 1 * time.Minute
	saramaConfig.Metadata.Retry.Max = 10
	saramaConfig.Metadata.Retry.Backoff = 2 * time.Second

	// Balance Strategy
	switch config.Config.ConsumerGroupBalanceStrategy {
	case "BalanceStrategyRange":
		saramaConfig.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	case "BalanceStrategySticky":
		saramaConfig.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
	case "BalanceStrategyRoundRobin":
		saramaConfig.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	default:
		saramaConfig.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	}

	var consumerGroup sarama.ConsumerGroup
	for {
		consumerGroup, err = sarama.NewConsumerGroup([]string{k.brokerURL}, group, saramaConfig)
		if err != nil {
			zap.S().Info("Creating consumer group consumerGroup err: ", err.Error())
			zap.S().Info("Retrying in 3 seconds...")
			time.Sleep(3 * time.Second)
			continue
		}
		break
	}

	// Get Kafka Jobs from database
	jobID := config.Config.ConsumerJobID
	kafkaJobs := &[]models.KafkaJob{}
	if jobID != "" {
		for {
			// Wait until jobs are present

			kafkaJobs, err = crud.GetKafkaJobCrud().SelectMany(
				jobID,
				group,
			)
			if errors.Is(err, gorm.ErrRecordNotFound) {
				zap.S().Info(
					"JobID=", jobID,
					",ConsumerGroup=", group,
					" - Waiting for Kafka Job in database...")
				time.Sleep(1)
				continue
			} else if err != nil {
				// Postgres error
				zap.S().Fatal(err.Error())
			}

			break
		}
	}

	// From example: /sarama/blob/master/examples/consumergroup/main.go
	ctx, cancel := context.WithCancel(context.Background())
	claimConsumer := &ClaimConsumer{
		topicNames: k.topicNames,
		topicChans: k.TopicChannels,
		group:      group,
		kafkaJobs:  *kafkaJobs,
	}

	for {
		// `Consume` should be called inside an infinite loop, when a
		// server-side rebalance happens, the consumer session will need to be
		// recreated to get the new claims
		err := consumerGroup.Consume(ctx, k.topicNames, claimConsumer)
		if err != nil {
			zap.S().Info("CONSUME GROUP ERROR: from consumer: ", err.Error())
		}
		// check if context was cancelled, signaling that the consumer should stop
		if ctx.Err() != nil {
			zap.S().Warn("CONSUME GROUP WARN: from context: ", ctx.Err().Error())
			return
		}
	}

	// Waiting, so that consumerGroup remains alive
	ch := make(chan int, 1)
	<-ch
	cancel()
}

type ClaimConsumer struct {
	topicNames []string
	topicChans map[string]chan *sarama.ConsumerMessage
	group      string
	kafkaJobs  []models.KafkaJob
}

func (c *ClaimConsumer) Setup(_ sarama.ConsumerGroupSession) error { return nil }
func (*ClaimConsumer) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (c *ClaimConsumer) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

	topicName := claim.Topic()
	partition := uint64(claim.Partition())

	// find kafka job
	var kafkaJob *models.KafkaJob = nil
	for i, k := range c.kafkaJobs {
		if k.Partition == partition && k.Topic == topicName {
			kafkaJob = &(c.kafkaJobs[i])
			break
		}
	}

	for {
		var topicMsg *sarama.ConsumerMessage
		select {
		case msg := <-claim.Messages():
			if msg == nil {
				zap.S().Info("GROUP=", c.group, ",TOPIC=", topicName, " - Kafka message is nil, exiting ConsumeClaim loop...")
				return nil
			}

			topicMsg = msg
		case <-time.After(5 * time.Second):
			zap.S().Info("GROUP=", c.group, ",TOPIC=", topicName, " - No new kafka messages, waited 5 secs...")
			continue
		case <-sess.Context().Done():
			zap.S().Info("GROUP=", c.group, ",TOPIC=", topicName, " - Session is done, exiting ConsumeClaim loop...")
			return nil
		}

		zap.S().Info("GROUP=", c.group, ",TOPIC=", topicName, ",PARTITION=", partition, ",OFFSET=", topicMsg.Offset, " - New message")
		sess.MarkMessage(topicMsg, "")

		// Broadcast
		c.topicChans[topicName] <- topicMsg

		// Check if kafka job is done
		// NOTE only applicable if ConsumerKafkaJobID is given
		if kafkaJob != nil &&
			uint64(topicMsg.Offset) >= kafkaJob.StopOffset+1000 {
			// Job done
			zap.S().Info(
				"JOBID=", config.Config.ConsumerJobID,
				"GROUP=", c.group,
				",TOPIC=", topicName,
				",PARTITION=", partition,
				" - Kafka Job done...exiting",
			)
			os.Exit(0)
		}
	}
	return nil
}
