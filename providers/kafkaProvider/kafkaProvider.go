package kafkaProvider

import (
	"context"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"uploadCSV/models"
	"uploadCSV/providers"
)

const (
	csvFileWriterBatchSize = 1
)

type KafkaProvider struct {
	csvFileWriter *kafka.Writer
}

func NewKafkaProvider() providers.KafkaProvider {
	kafkaHost := "localhost:9092"

	// chatWriter is a kafka writer for chat messages
	csvFileWriter := &kafka.Writer{
		Addr:      kafka.TCP(kafkaHost),
		Topic:     string(models.TopicCSVFileUpload),
		Balancer:  &kafka.RoundRobin{},
		BatchSize: csvFileWriterBatchSize,
	}

	return &KafkaProvider{
		csvFileWriter: csvFileWriter,
	}
}

func (k *KafkaProvider) Publish(topic models.Topic, message []byte, metaData map[models.KafkaHeaders]interface{}) {
	switch topic {
	case models.TopicCSVFileUpload:
		err := k.csvFileWriter.WriteMessages(context.Background(),
			kafka.Message{
				Value: message,
			},
		)
		if err != nil {
			logrus.Errorf("Publish: failed to write csv file upload: %v", err)
		}
	default:
		logrus.Warn("Trying to publish on wrong topic")
	}
}

func (k *KafkaProvider) Close() {
	if err := k.csvFileWriter.Close(); err != nil {
		logrus.Errorf("error closing kafka chat connection")
	}
}
