package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type configType struct {
	Name        string `envconfig:"NAME" required:"false" default:"blocks-service"`
	NetworkName string `envconfig:"NETWORK_NAME" required:"false" default:"mainnnet"`

	// Logging
	LogLevel         string `envconfig:"LOG_LEVEL" required:"false" default:"INFO"`
	LogToFile        bool   `envconfig:"LOG_TO_FILE" required:"false" default:"false"`
	LogFileName      string `envconfig:"LOG_FILE_NAME" required:"false" default:"etl.log"`
	LogFormat        string `envconfig:"LOG_FORMAT" required:"false" default:"console"`
	LogIsDevelopment bool   `envconfig:"LOG_IS_DEVELOPMENT" required:"false" default:"true"`

	// Kafka
	KafkaBrokerURL   string `envconfig:"KAFKA_BROKER_URL" required:"false" default:"localhost:29092"`
	KafkaBlocksTopic string `envconfig:"KAFKA_BLOCKS_TOPIC" required:"false" default:"icon-blocks"`

	// Consumer Group
	ConsumerGroup                string `envconfig:"CONSUMER_GROUP" required:"false" default:"blocks-consumer-group"`
	ConsumerGroupBalanceStrategy string `envconfig:"CONSUMER_GROUP_BALANCE_STRATEGY" required:"false" default:"BalanceStrategySticky"`

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
