package producer

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	cfg "kafka-service/app/config"
	"time"
)

func NewKafkaProducer(config cfg.Config) {

	p, err := kafka.NewProducer(
		&kafka.ConfigMap{
			"bootstrap.servers": config.Kafka.BootstrapServers,
		})

	if err != nil {
		panic(err)
	}

	defer p.Close()

	fmt.Printf("Created Producer %v\n\n", p)

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case kafka.Error:
				fmt.Printf("%% Error: %v\n", ev)
			default:
				fmt.Printf("Ignored event: %v\n", ev)
			}
		}
	}()

	msgCount := 0

	for msgCount < 10 {
		value := fmt.Sprintf("Message %d", msgCount)

		deliveryChan := make(chan kafka.Event)
		go func() {
			for e := range deliveryChan {
				switch ev := e.(type) {
				case *kafka.Message:
					m := ev
					if m.TopicPartition.Error != nil {
						fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
					} else {
						fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
							*m.TopicPartition.Topic,
							m.TopicPartition.Partition,
							m.TopicPartition.Offset)
					}
				default:
					fmt.Printf("Ignored event: %s\n", ev)
				}
				close(deliveryChan)
			}
		}()

		err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &config.Kafka.Topic,
				Partition: kafka.PartitionAny},
			Value:   []byte(value),
			Headers: []kafka.Header{{Key: "myTestHeader", Value: []byte("header values are binary")}},
		}, deliveryChan)

		if err != nil {
			close(deliveryChan)
			if err.(kafka.Error).Code() == kafka.ErrQueueFull {
				time.Sleep(time.Second)
				continue
			}
			fmt.Printf("Failed to produce message: %v\n", err)
		}

		msgCount++
	}

	for p.Flush(100000) > 0 {
		fmt.Print("Still waiting for flush outstanding messages\n", err)
	}

	p.Close()
}
