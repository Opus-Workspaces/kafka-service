### Kafka Setup

#### docker-compose.yml

```yaml
---
version: "3"
services:
  zookeeper:
    image: confluentinc/cp-zookeeper:7.4.1
    hostname: zookeeper
    container_name: zookeeper
    environment:
      ZOOKEEPER_CLIENT_PORT: 2181
      ZOOKEEPER_TICK_TIME: 2000

  broker:
    image: confluentinc/cp-kafka:7.4.1
    container_name: broker
    ports:
      - "9092:9092"
    depends_on:
      - zookeeper
    environment:
      KAFKA_BROKER_ID: 1
      KAFKA_ZOOKEEPER_CONNECT: "zookeeper:2181"
      KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: PLAINTEXT:PLAINTEXT,PLAINTEXT_INTERNAL:PLAINTEXT
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://localhost:9092,PLAINTEXT_INTERNAL://broker:29092
      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 1
      KAFKA_TRANSACTION_STATE_LOG_MIN_ISR: 1
      KAFKA_TRANSACTION_STATE_LOG_REPLICATION_FACTOR: 1
```

Now start the Kafka broker with the new docker compose command (see the [Docker Compose documentation for more details](https://docs.docker.com/compose/cli-command/#new-docker-compose-command)):

```bash
  docker compose up -d

  ## Stop the Kafka broker
    docker compose down
```

### Configuration

Paste the following configuration into a file named `getting-started.properties`:

```properties
bootstrap.servers=localhost:9092
```

### Create a Topic

```bash
  docker compose exec broker kafka-topics --create --topic purchases --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1
```
