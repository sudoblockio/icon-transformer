package config

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type configType struct {
	Name        string `envconfig:"NAME" required:"false" default:"icon-transformer"`
	NetworkName string `envconfig:"NETWORK_NAME" required:"false" default:"mainnnet"`

	// Logging
	LogLevel         string `envconfig:"LOG_LEVEL" required:"false" default:"INFO"`
	LogToFile        bool   `envconfig:"LOG_TO_FILE" required:"false" default:"false"`
	LogFileName      string `envconfig:"LOG_FILE_NAME" required:"false" default:"etl.log"`
	LogFormat        string `envconfig:"LOG_FORMAT" required:"false" default:"console"`
	LogIsDevelopment bool   `envconfig:"LOG_IS_DEVELOPMENT" required:"false" default:"true"`

	// Icon node service
	IconNodeServiceURL           []string      `envconfig:"ICON_NODE_SERVICE_URL" required:"false" default:"https://api.icon.community/api/v3"`
	IconNodeRpcRetrySleepSeconds time.Duration `envconfig:"ICON_NODE_RPC_SLEEP_SECONDS" required:"false" default:"1s"`
	IconNodeRpcRetryAttempts     int           `envconfig:"ICON_NODE_RPC_RETRY_ATTEMPTS" required:"false" default:"20"`

	// Kafka
	KafkaBrokerURL string `envconfig:"KAFKA_BROKER_URL" required:"false" default:"localhost:29092"`

	// Kafka Topics
	// NOTE add to string array in kafka/consumer.go
	KafkaBlocksTopic      string `envconfig:"KAFKA_BLOCKS_TOPIC" required:"false" default:"icon-blocks"`
	KafkaContractsTopic   string `envconfig:"KAFKA_CONTRACTS_TOPIC" required:"false" default:"icon-contracts"`
	KafkaDeadMessageTopic string `envconfig:"KAFKA_DEAD_MESSAGE_TOPIC" required:"false" default:"icon-blocks-dead"`

	// Consumer Group
	ConsumerGroup                string `envconfig:"CONSUMER_GROUP" required:"false" default:"blocks-consumer-group"`
	ConsumerGroupBalanceStrategy string `envconfig:"CONSUMER_GROUP_BALANCE_STRATEGY" required:"false" default:"BalanceStrategyRange"`

	// Consumer Tail
	ConsumerIsTail bool   `envconfig:"CONSUMER_IS_TAIL" required:"false" default:"false"`
	ConsumerJobID  string `envconfig:"CONSUMER_JOB_ID" required:"false" default:""`

	// Consumer Partition
	ConsumerIsPartitionConsumer  bool   `envconfig:"CONSUMER_IS_PARTITION_CONSUMER" required:"false" default:"false"`
	ConsumerPartition            int    `envconfig:"CONSUMER_PARTITION" required:"false" default:"0"`
	ConsumerPartitionTopic       string `envconfig:"CONSUMER_PARTITION_TOPIC" required:"false" default:"blocks"`
	ConsumerPartitionStartOffset int    `envconfig:"CONSUMER_PARTITION_START_OFFSET" required:"false" default:"1"`

	// DB
	DbDriver             string `envconfig:"DB_DRIVER" required:"false" default:"postgres"`
	DbHost               string `envconfig:"DB_HOST" required:"false" default:"localhost"`
	DbPort               string `envconfig:"DB_PORT" required:"false" default:"5432"`
	DbUser               string `envconfig:"DB_USER" required:"false" default:"postgres"`
	DbPassword           string `envconfig:"DB_PASSWORD" required:"false" default:"changeme"`
	DbName               string `envconfig:"DB_DBNAME" required:"false" default:"postgres"`
	DbSslmode            string `envconfig:"DB_SSL_MODE" required:"false" default:"disable"`
	DbTimezone           string `envconfig:"DB_TIMEZONE" required:"false" default:"UTC"`
	DbMaxIdleConnections int    `envconfig:"DB_MAX_IDLE_CONNECTIONS" required:"false" default:"2"`
	DbMaxOpenConnections int    `envconfig:"DB_MAX_OPEN_CONNECTIONS" required:"false" default:"10"`

	// GORM
	GormLoggingThresholdMilli int `envconfig:"GORM_LOGGING_THRESHOLD_MILLI" required:"false" default:"250"`

	// Redis
	RedisHost                     string `envconfig:"REDIS_HOST" required:"false" default:"localhost"`
	RedisPort                     string `envconfig:"REDIS_PORT" required:"false" default:"6379"`
	RedisPassword                 string `envconfig:"REDIS_PASSWORD" required:"false" default:""`
	RedisSentinelClientMode       bool   `envconfig:"REDIS_SENTINEL_CLIENT_MODE" required:"false" default:"false"`
	RedisSentinelClientMasterName string `envconfig:"REDIS_SENTINEL_CLIENT_MASTER_NAME" required:"false" default:"master"`
	RedisKeyPrefix                string `envconfig:"REDIS_KEY_PREFIX" required:"false" default:"icon_"`

	// Redis Channels
	// NOTE must add to redis client manually
	// src/redis/client.go:63
	RedisBlocksChannel         string `envconfig:"REDIS_BLOCKS_CHANNEL" required:"false" default:"blocks"`
	RedisTransactionsChannel   string `envconfig:"REDIS_TRANSACTIONS_CHANNEL" required:"false" default:"transactions"`
	RedisLogsChannel           string `envconfig:"REDIS_LOGS_CHANNEL" required:"false" default:"logs"`
	RedisTokenTransfersChannel string `envconfig:"REDIS_TOKEN_TRANSFERS_CHANNEL" required:"false" default:"token_transfers"`

	// Transformer
	TransformerServiceCallThreshold  time.Duration `envconfig:"TRANSFORMER_SERVICE_CALL_THRESHOLD" required:"false" default:"1h"`
	TransformerRedisChannelThreshold time.Duration `envconfig:"TRANSFORMER_REDIS_CHANNEL_THRESHOLD" required:"false" default:"15s"`

	// Routines
	RoutinesRunOnly       bool          `envconfig:"ROUTINES_RUN_ONLY" required:"false" default:"false"`
	RoutinesSleepDuration time.Duration `envconfig:"ROUTINES_SLEEP_DURATION" required:"false" default:"1h"`
	RoutinesBatchSize     int           `envconfig:"ROUTINES_BATCH_SIZE" required:"false" default:"1000"`
	RoutinesNumWorkers    int           `envconfig:"ROUTINES_NUM_WORKERS" required:"false" default:"1"`

	// FindMissing
	FindMissingRunOnly              bool   `envconfig:"FIND_MISSING_RUN_ONLY" required:"false" default:"false"`
	FindMissingStartBlock           int64  `envconfig:"FIND_MISSING_START_BLOCK" required:"false" default:"0"`
	FindMissingEndBlock             int64  `envconfig:"FIND_MISSING_END_BLOCK" required:"false" default:"0"`
	FindMissingExtractorAPILocation string `envconfig:"FIND_MISSING_EXTRACTOR_API_LOCATION" required:"false" default:"http://localhost:8000/api/v1"`

	// Redis Recovery
	RedisRecoveryRunOnly bool `envconfig:"REDIS_RECOVERY_RUN_ONLY" required:"false" default:"false"`
}

// Config - runtime config struct
var Config configType

// ReadEnvironment - Read and store runtime config
func ReadEnvironment() {
	err := envconfig.Process("", &Config)
	if err != nil {
		log.Fatalf("ERROR: envconfig - %s\n", err.Error())
	}
}
