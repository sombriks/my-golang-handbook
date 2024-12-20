package sample_kafka

import (
	"github.com/IBM/sarama"
	"log"
)

func NewConsumer(addrs []string, config *sarama.Config) *sarama.Consumer {

	consumer, err := sarama.NewConsumer(addrs, config)
	if err != nil {
		log.Fatal(err)
	}

	return &consumer
}
