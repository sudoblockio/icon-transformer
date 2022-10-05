package transformers

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sudoblockio/icon-transformer/config"
	"github.com/sudoblockio/icon-transformer/crud"
	"github.com/sudoblockio/icon-transformer/kafka"
	"github.com/sudoblockio/icon-transformer/metrics"
	"github.com/sudoblockio/icon-transformer/models"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"
	"google.golang.org/protobuf/proto"
	"sync"
)

func startBlocks() {
	kafkaBlocksTopic := config.Config.KafkaBlocksTopic

	// Input channels
	kafkaBlocksTopicChannel := kafka.KafkaTopicConsumer.TopicChannels[kafkaBlocksTopic]

	// Build the transformer list
	filterTransformers()

	// Call init functions
	callProcessorInits()

	// Setup metrics
	setupBlockTransformerMetrics()

	// Counter for displaying logs
	var blockLogCounter int = 0

	zap.S().Debug("Blocks transformer: started working")
	for {
		// Consume from kafka
		consumerTopicMsg := <-kafkaBlocksTopicChannel
		blockETL := &models.BlockETL{}
		err := proto.Unmarshal(consumerTopicMsg.Value, blockETL)
		if err != nil {
			zap.S().Warn(
				"Routine=", "Transformer",
				" Partition=", consumerTopicMsg.Partition,
				" Offset=", consumerTopicMsg.Offset,
				" Key=", consumerTopicMsg.Key,
				" Value=", consumerTopicMsg.Value,
				" Step=", "Parse block ETL from kafka proto",
				" Error=", err.Error(),
			)
			continue
		}

		if blockLogCounter > config.Config.LogMsgCount-1 {
			zap.S().Info("Transformer: Processing block #", blockETL.Number)
			blockLogCounter = 0
		}
		blockLogCounter++
		runBlockProcessors(blockETL)
	}
}

type Processor struct {
	transformerFunction func(a *models.BlockETL)
	initFunctions       []func()
}

var Processors = []Processor{
	{
		transformerFunction: blocks,
		initFunctions:       []func(){crud.InitBlockCrud},
	},
	{
		transformerFunction: transactions,
		initFunctions:       []func(){crud.InitTransactionCrud},
	},
	{
		transformerFunction: logs,
		initFunctions:       []func(){crud.InitLogCrud},
	},
	{
		transformerFunction: tokenTransfers,
		initFunctions:       []func(){crud.InitTokenTransferCrud},
	},
	{
		transformerFunction: transactionByAddresses,
		initFunctions:       []func(){crud.InitTransactionByAddressCrud},
	},
	{
		transformerFunction: transactionInternalByAddresses,
		initFunctions:       []func(){crud.InitTransactionInternalByAddressCrud},
	},
	{
		transformerFunction: transactionByAddressCreateScores,
		initFunctions:       []func(){crud.InitTransactionCrud, crud.InitTransactionByAddressCrud},
	},
	{
		transformerFunction: tokenTransfers,
		initFunctions:       []func(){crud.InitTokenTransferCrud},
	},
	{
		transformerFunction: addresses,
		initFunctions:       []func(){crud.InitAddressCrud},
	},
	{
		transformerFunction: tokenAddresses,
		initFunctions:       []func(){crud.InitTokenAddressCrud},
	},
	{
		transformerFunction: tokenTransferByAddress,
		initFunctions:       []func(){crud.InitTokenTransferByAddressCrud},
	},
}

// Filter the list of transform functions based on the value in the config
func filterTransformers() {

	if len(config.Config.TransformerFunctions) == 0 {
		zap.S().Info("Running all processors.")
		return
	}

	var processors []Processor

	for i, v := range Processors {
		if slices.Contains(config.Config.TransformerFunctions, getFunctionName(v.transformerFunction)) {
			processors = append(processors, Processors[i])
		}
	}

	if len(processors) == 0 {
		zap.S().Fatal("No processors could be found. Update setting=", config.Config.TransformerFunctions)
	}
	zap.S().Info("Running ", len(processors), " processors.")

	Processors = processors
}

func callProcessorInits() {
	for i := 0; i < len(Processors); i++ {
		for f := 0; f < len(Processors[i].initFunctions); f++ {
			Processors[i].initFunctions[f]()
		}
	}
}

var metricsBlockTransformer struct {
	addressesSeen    prometheus.Counter
	addressesIgnored prometheus.Counter
}

func setupBlockTransformerMetrics() {
	metricsBlockTransformer.addressesSeen = metrics.CreateCounter(
		"transformer_addresses_seen",
		"number of addresses transformed",
		nil,
	)
	metricsBlockTransformer.addressesIgnored = metrics.CreateCounter(
		"transformer_addresses_ignored",
		"number of addresses ignored",
		nil,
	)
}

func runBlockProcessors(blockETL *models.BlockETL) {
	var wg sync.WaitGroup
	for _, p := range Processors {
		wg.Add(1)

		f := p.transformerFunction
		go func() {
			defer wg.Done()
			f(blockETL)
		}()
	}

	wg.Wait()
}
