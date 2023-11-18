package service

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Producer struct {
	*kafka.Producer
	deliveryChan chan kafka.Event
}

func NewProducer(producer *kafka.Producer) *Producer {
	return &Producer{
		producer,
		make(chan kafka.Event, 1000),
	}
}

func (c Producer) SendMsg(topic, value string) {
	err := c.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(value)},
		c.deliveryChan,
	)
	if err != nil {
		fmt.Println(err)

	}
	go func() {
		for e := range c.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Successfully produced record to topic %s partition [%d] @ offset %v\n",
						*ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.TopicPartition.Offset)
				}
			}
		}
	}()
}
