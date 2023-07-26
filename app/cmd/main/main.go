package main

import (
	"kafka-service/app/cmd/server"
	"kafka-service/app/config"
	"kafka-service/app/domains/v1/modules"
	"kafka-service/app/domains/v1/modules/kafka/logic"
)

func main() {
	cfg := config.LoadConfig()
	s := server.InitServer(cfg)

	modules.Modules(s)

	logic.NewKafkaProducer(cfg)
	logic.NewKafkaConsumer(cfg)
	server.Run(s)
}
