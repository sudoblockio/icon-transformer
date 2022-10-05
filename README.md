<p align="center">
  <h2 align="center">ICON Transformer</h2>
</p>


Icon Transformer lets you convert ICON blockchain data from a kafka topic to a postgres database and redis channels

### Quickstart
```bash
make up
```

### Example Config
```docker-compose
version: "3.7"

x-transformer-env: &transformer-env

  # Kakfa
  KAFKA_BROKER_URL: "kafka:9092"
  KAFKA_BLOCKS_TOPIC: "icon-blocks"         # produced by icon-extractor
  KAFKA_CONTRACTS_TOPIC: "icon-contracts"   # produced by icon-contracts-worker

  # DB
  DB_DRIVER: "postgres"
  DB_HOST: "postgres"
  DB_PORT: "5432"
  DB_USER: "postgres"
  DB_PASSWORD: "changeme"
  DB_DBNAME: "postgres"
  DB_SSL_MODE: "disable"
  DB_TIMEZONE: "UTC"

  # Redis
  REDIS_HOST: "redis"
  REDIS_PORT: "6379"
  REDIS_PASSWORD: ""

services:
  transformer:
    image: sudoblock/icon-transformer:latest
    environment:
      <<: *transformer-env

  transformer-routine:
    image: sudoblock/icon-transformer:latest
    environment:
      <<: *transformer-env
      ROUTINES_RUN_ONLY: "true"

```

#### Transformer
Simply reads from a kafka topic, transforms the message into various views, and inserts into various services. This service will read from the two topics and insert data into a postgres database, count records in redis keys, and stream data into redis channels. All centered around `/src/transformers/transformer.go`
> NOTE: The original icon-explorer project was based on the icon-etl project, which seperated the data into 3 topics; blocks, transactions, logs. This caused many issues downstream in the first version, so the the icon-extractor was created. The icon-extractor combined the 3 streams of data into one and fixed these issues downstream, this service.

![unnamed](https://user-images.githubusercontent.com/77865393/165399000-6be57fce-2101-4b5f-8a16-9d80bbf485ed.png)
