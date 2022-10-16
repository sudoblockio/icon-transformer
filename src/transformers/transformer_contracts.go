package transformers

import (
	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/kafka"
	"github.com/sudoblockio/icon-transformer/models"
	"go.uber.org/zap"
	// Depending on the version of protoc, we need to switch these out
	// https://stackoverflow.com/a/65962599/12642712
	//"google.golang.org/protobuf/proto"
	"github.com/golang/protobuf/proto"
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

		// Upsert Contract to Address
		transformContractsToLoadAddress(contractProcessed)
	}
}

func transformContractToAddress(contract *models.ContractProcessed) *models.Address {

	address := &models.Address{
		Address:              contract.Address,
		Name:                 contract.Name,
		CreatedTimestamp:     contract.CreatedTimestamp,
		Status:               contract.Status,
		IsToken:              contract.IsToken,
		IsContract:           true,
		ContractUpdatedBlock: contract.ContractUpdatedBlock,
		ContractType:         contract.ContractType,
		TokenStandard:        contract.TokenStandard,
		//Symbol:               contract.Symbol,
	}

	return address
}

//func transformContractToTransactionByAddress(contract *models.ContractProcessed) *models.TransactionByAddress {
//	// Facing issues getting the contract creation events into the transactions so this hack is here to make
//	// sure we have a hash to lookup the Tx by.
//	transaction_by_address := &models.TransactionByAddress{
//		Address:         contract.Address,
//		TransactionHash: contract.TransactionHash,
//		BlockNumber:     contract.ContractUpdatedBlock,
//	}
//
//	return transaction_by_address
//}
//
//func transformContractToTransaction(contract *models.ContractProcessed) *models.Transaction {
//	// Facing issues getting the contract creation events into the transactions so this hack is here to make
//	// sure we get the right `transaction_type` classification into the transactions page.
//	// Need to first lookup prior
//
//	transaction := &models.Transaction{
//		Hash: contract.TransactionHash,
//	}
//
//	return transaction
//}

// Address loader
func transformContractsToLoadAddress(contract *models.ContractProcessed) {
	address := transformContractToAddress(contract)

	var dbAddresses *models.Address
	dbAddresses, err := crud.GetAddressCrud().SelectOneWhere("address", address.Address)
	if err != nil {
		zap.S().Info(err.Error())
	}

	if dbAddresses.ContractUpdatedBlock <= address.ContractUpdatedBlock {
		err = crud.GetAddressCrud().UpsertOneColumns(address, []string{
			"address",
			"name",
			"created_timestamp",
			"status",
			"is_token",
			"is_contract",
			"contract_updated_block",
			"contract_type",
			"token_standard",
		})
		if err != nil {
			zap.S().Fatal(err.Error())
		}
	}

	//transaction_by_address := transformContractToTransactionByAddress(contract)
	//err = crud.GetTransactionByAddressCrud().UpsertOne(transaction_by_address)
	//if err != nil {
	//	zap.S().Info(err.Error())
	//}
	//
	//transaction := transformContractToTransactionByAddress(contract)
	//err = crud.GetTransactionCrud().UpsertOne(transaction)
	//if err != nil {
	//	zap.S().Info(err.Error())
	//}
}
