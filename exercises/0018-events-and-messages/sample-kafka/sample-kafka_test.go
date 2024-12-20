package sample_kafka

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go/modules/kafka"
	"testing"
	"time"
)

type SampleKafkaTestSuite struct {
	suite.Suite
	kafkaContainer *kafka.KafkaContainer
	config         *sarama.Config
	context        *context.Context
}

func (suite *SampleKafkaTestSuite) SetupTest() {
	ctx := context.Background()
	suite.context = &ctx
	var err error

	suite.kafkaContainer, err = kafka.Run(*suite.context,
		"confluentinc/confluent-local:7.5.0",
		kafka.WithClusterID("sample-cluster"))

	require.Nil(suite.T(), err)

	suite.config = sarama.NewConfig()
	suite.config.Producer.Return.Successes = true
}

func TestBootstrap(t *testing.T) {
	suite.Run(t, new(SampleKafkaTestSuite))
}

func (suite *SampleKafkaTestSuite) Test01Producer() {
	brokers, err := suite.kafkaContainer.Brokers(*suite.context)

	require.Nil(suite.T(), err)

	var producer *sarama.SyncProducer
	producer = NewProducer(brokers, suite.config)
	require.NotNil(suite.T(), producer)

	msg := &sarama.ProducerMessage{
		Topic:     "sample-topic",
		Key:       sarama.StringEncoder("unique-key"),
		Value:     sarama.StringEncoder("unique-value"),
		Timestamp: time.Time{},
	}

	partition, offset, err := (*producer).SendMessage(msg)

	require.NotNil(suite.T(), partition)
	require.NotNil(suite.T(), offset)
	require.Nil(suite.T(), err)
}

func (suite *SampleKafkaTestSuite) Test02Consumer() {
	brokers, err := suite.kafkaContainer.Brokers(*suite.context)

	// send a message

	var producer *sarama.SyncProducer
	producer = NewProducer(brokers, suite.config)
	require.NotNil(suite.T(), producer)

	msg := &sarama.ProducerMessage{
		Topic:     "sample-topic",
		Key:       sarama.StringEncoder("unique-key"),
		Value:     sarama.StringEncoder("unique-value"),
		Timestamp: time.Time{},
	}

	_, _, err = (*producer).SendMessage(msg)

	require.Nil(suite.T(), err)

	// consumer setup

	consumer := NewConsumer(brokers, suite.config)
	defer (*consumer).Close()

	require.NotNil(suite.T(), consumer)

	partitionConsumer, err := (*consumer).
		ConsumePartition("sample-topic", 0, sarama.OffsetOldest)

	require.NotNil(suite.T(), partitionConsumer)
	require.Nil(suite.T(), err)

	var message *sarama.ConsumerMessage
	message = <-partitionConsumer.Messages() // consumer hangs here waiting for messages
	require.NotNil(suite.T(), message)
	assert.Equal(suite.T(), "unique-value", string(message.Value))

}
