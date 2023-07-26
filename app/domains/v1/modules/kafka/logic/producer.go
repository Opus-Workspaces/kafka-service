package logic

import (
	"fmt"
	"kafka-service/app/config"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func NewKafkaProducer(config config.Config) {

	p, err := kafka.NewProducer(
		&kafka.ConfigMap{
			"bootstrap.servers": config.Kafka.BootstrapServers,
		})

	if err != nil {
		panic(err)
	}

	defer p.Close()

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					panic(ev.TopicPartition.Error)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	words := []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"}
	topic := config.Kafka.Topic
	for _, word := range words {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
	}

	p.Flush(15 * 1000) // wait for delivery report or error
}
