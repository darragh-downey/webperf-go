package producer

import (
	"context"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	topic          = ""
	broker1Address = ""
	broker2Address = ""
	broker3Address = ""
)

// derived from https://github.com/confluentinc/confluent-kafka-go
func publish(ctx context.Context) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		fmt.Printf("Error: Failed to initialise Kafka Producer %v", err)
	}

	defer p.Close()

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Error: Delivery failed! %v", ev.TopicPartition)
				} else {
					fmt.Printf("I: Delivered message to %v", ev.TopicPartition)
				}

			}
		}
	}()

	// produce messages to topic asynchronously
	topic := ""
	// the array of strings should be a go channel listening for responses
	// from the poller
	for _, word := range []string{} {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
	}

	p.Flush(15 * 1000)
}
