package transformers

import (
	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/kafka"
	"github.com/sudoblockio/icon-transformer/models"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

func startContracts() {
	kafkaContractsTopic := config.Config.KafkaContractsTopic

	// Input channels
	kafkaContractsTopicChannel := kafka.KafkaTopicConsumer.TopicChannels[kafkaContractsTopic]

	zap.S().Debug("Contracts transformer: started working")
	for {

		///////////////////
		// Kafka Message //
		///////////////////

		consumerTopicMsg := <-kafkaContractsTopicChannel
		contractProcessed := &models.ContractProcessed{}
		err := proto.Unmarshal(consumerTopicMsg.Value, contractProcessed)
		if err != nil {
			zap.S().Warn(
				"Routine=", "Transformer",
				" Partition=", consumerTopicMsg.Partition,
				" Offset=", consumerTopicMsg.Offset,
				" Key=", consumerTopicMsg.Key,
				" Value=", consumerTopicMsg.Value,
				" Step=", "Parse contract processed from kafka proto",
				" Error=", err.Error(),
			)
			continue
		}
		zap.S().Info("Transformer: Processing contract address ", contractProcessed.Address)

		/////////////
		// Loaders //
		/////////////
		// NOTE transform contracts processed to various database views
		// NOTE contracts may be passed multiple times, loaders use upserts

		// Address loader
		transformContractsToLoadAddress(contractProcessed)
	}
}

// Address loader
func transformContractsToLoadAddress(contract *models.ContractProcessed) {
	address := transformContractToAddress(contract)
	db_address, err := crud.GetAddressCrud().SelectOneAddress(address.Address)
	if err != nil {
		zap.S().Info(err.Error())
	}

	if db_address.ContractUpdatedBlock <= address.ContractUpdatedBlock {
		crud.GetAddressCrud().UpsertOneCols(address, []string{"address", "name", "created_timestamp", "status", "is_token", "is_contract", "contract_updated_block"})
	}
}
