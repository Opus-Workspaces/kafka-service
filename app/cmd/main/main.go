package main

import (
	"kafka-service/app/cmd/server"
	"kafka-service/app/config"
	"kafka-service/app/domains/v1/modules"
	"kafka-service/app/internal/kafka/consumer"
	"kafka-service/app/internal/kafka/producer"
)

func main() {
	cfg := config.LoadConfig()
	s := server.InitServer(cfg)

	modules.Modules(s)

	producer.NewKafkaProducer(cfg)
	consumer.NewKafkaConsumer(cfg)

	server.Run(s)
}
