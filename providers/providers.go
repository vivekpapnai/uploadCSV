package providers

import (
	"context"
	"mime/multipart"
	"uploadCSV/models"
)

type StorageProvider interface {
	Upload(ctx context.Context, file multipart.File, fileName, contentType string) (string, error)
	GetSharableURL() (string, error)
}

type KafkaProvider interface {
	Publish(topic models.Topic, message []byte, metaData map[models.KafkaHeaders]interface{})
}
