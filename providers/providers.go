package providers

import (
	"context"
	"mime/multipart"
)

type StorageProvider interface {
	Upload(ctx context.Context, file multipart.File, fileName, contentType string) (string, error)
	GetSharableURL() (string, error)
}

type KafkaProvider interface {
}