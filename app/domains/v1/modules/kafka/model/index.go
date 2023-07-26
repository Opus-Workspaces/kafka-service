package model

import (
	"context"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaProducer struct {
	Writer *kafka.Producer
}

type KafkaConsumer struct {
	reader          *kafka.Consumer
	boostrapServers []string
}

type KafkaConsumerLogic interface {
	FetchMessage(ctx context.Context) (*kafka.Message, error)
	TransformMessage(ctx context.Context) (string, error)
}

type KafkaConsumerRepo interface {
	FetchMessage(ctx context.Context) (*kafka.Message, error)
	Commit(ctx context.Context, msg ...*kafka.Message) error
}
