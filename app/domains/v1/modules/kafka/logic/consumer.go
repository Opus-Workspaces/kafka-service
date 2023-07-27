package logic

import (
	"context"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"kafka-service/app/domains/v1/modules/kafka/model"
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
