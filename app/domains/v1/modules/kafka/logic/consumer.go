package logic

import (
	"context"
	"fmt"
	"kafka-service/app/config"
	"kafka-service/app/domains/v1/modules/kafka/model"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaHandler struct {
	KafkaRepo model.KafkaConsumerRepo
}

func NewKafkaHandler(kafkaRepo model.KafkaConsumerRepo) model.KafkaConsumerLogic {
	return &KafkaHandler{
		KafkaRepo: kafkaRepo,
	}
}
func (k *KafkaHandler) FetchMessage(ctx context.Context) (*kafka.Message, error) {
	return k.KafkaRepo.FetchMessage(ctx)
}

func (k *KafkaHandler) TransformMessage(ctx context.Context) (string, error) {
	return "pong", nil
}

func NewKafkaConsumer(config config.Config) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": config.Kafka.BootstrapServers,
		"group.id":          config.Kafka.GroupID,
		"auto.offset.reset": config.Kafka.AutoOffsetReset,
	})

	if err != nil {
		panic(err)
	}

	defer c.Close()

	err = c.SubscribeTopics(
		[]string{config.Kafka.Topic, "^aRegex.*[Tt]opic"},
		nil)

	if err != nil {
		panic(err)
	}

	run := true

	for run {
		msg, err := c.ReadMessage(time.Second)

		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else if !err.(kafka.Error).IsTimeout() {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

}
