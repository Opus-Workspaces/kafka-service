package consumer

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	cfg "kafka-service/app/config"
	"os"
	"os/signal"
	"syscall"
)

func NewKafkaConsumer(config cfg.Config) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM) // catch SIGINT and SIGTERM to terminate consumer gracefully

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":        config.Kafka.BootstrapServers,
		"group.id":                 config.Kafka.GroupID,
		"auto.offset.reset":        config.Kafka.AutoOffsetReset,
		"session.timeout.ms":       config.Kafka.SessionTimeoutMs,
		"enable.auto.offset.store": config.Kafka.EnableAutoOffsetStore,
	})

	if err != nil {
		panic(err)
	}

	defer c.Close()

	fmt.Printf("Created Consumer %v\n\n", c)

	err = c.SubscribeTopics(
		[]string{config.Kafka.Topic, "^aRegex.*[Tt]opic"},
		nil)

	if err != nil {
		panic(err)
	}

	run := true

	for run == true {
		select {
		case sig := <-sigChan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev := c.Poll(100)
			if ev == nil {
				continue
			}
			switch e := ev.(type) {
			case *kafka.Message:
				fmt.Printf("%% Message on %s:\n%s\n", e.TopicPartition, string(e.Value))
				if e.Headers != nil {
					fmt.Printf("%% Headers: %v\n", e.Headers)
				}

				_, err := c.StoreMessage(e)
				if err != nil {
					fmt.Fprintf(os.Stderr, "%% Error storing offset after message %s:\n",
						e.TopicPartition)
				}
			case kafka.Error:
				fmt.Fprintf(os.Stderr, "%% Error: %v: %v\n", e.Code(), e)
				if e.Code() == kafka.ErrAllBrokersDown {
					run = false
				}
			default:
				fmt.Printf("Ignored %v\n", e)
			}
		}
	}

	fmt.Printf("Closing consumer\n")
	c.Close()
}
