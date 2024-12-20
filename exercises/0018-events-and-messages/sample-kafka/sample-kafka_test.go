package sample_kafka

import (
	"context"
	"github.com/IBM/sarama"
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

func (suite *SampleKafkaTestSuite) TestProducer() {
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

func (suite *SampleKafkaTestSuite) TestConsumer() {
	brokers, err := suite.kafkaContainer.Brokers(*suite.context)

	require.Nil(suite.T(), err)

	NewConsumer(brokers, suite.config)
}
