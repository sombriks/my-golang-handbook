package sample_kafka

import (
	"github.com/IBM/sarama"
	"log"
)

func NewProducer(brokerList []string, config *sarama.Config) *sarama.SyncProducer {
	producer, err := sarama.NewSyncProducer(brokerList, config)

	if err != nil {
		log.Fatal(err)
	}

	return &producer
}
