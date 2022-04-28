package transformers

import (
	"github.com/Shopify/sarama"
	"github.com/sudoblockio/icon-transformer/models"
)

func transformDeadMessageToDeadBlock(deadMessage *sarama.ConsumerMessage) *models.DeadBlock {

	return &models.DeadBlock{
		Topic:     deadMessage.Topic,
		Partition: int64(deadMessage.Partition),
		Offset:    deadMessage.Offset,
		Key:       string(deadMessage.Key),
		Value:     string(deadMessage.Value),
	}
}
