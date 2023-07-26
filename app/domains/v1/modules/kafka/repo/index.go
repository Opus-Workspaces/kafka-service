package repo

import (
	"context"
	"kafka-service/app/domains/v1/modules/kafka/model"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaRepo struct {
	kafka *kafka.Producer
}

func NewKafkaRepo(kafka *kafka.Producer) model.KafkaConsumerRepo {
	return &KafkaRepo{kafka: kafka}
}

func (k KafkaRepo) FetchMessage(ctx context.Context) (*kafka.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (k KafkaRepo) Commit(ctx context.Context, msg ...*kafka.Message) error {
	//TODO implement me
	panic("implement me")
}
