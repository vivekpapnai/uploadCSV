package storageProvider

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/sirupsen/logrus"
	"mime/multipart"
	"os"
	"uploadCSV/providers"
)

var AccessKeyID string
var SecretAccessKey string

const AWSRegion = "ap-south-1"

//GetEnvWithKey : get env value
func GetEnvWithKey(key string) string {
	return os.Getenv(key)
}

type AWSStorage struct {
	Session *session.Session
}

func NewStorageProvider() providers.StorageProvider {
	AccessKeyID = os.Getenv("AWS_ACCESS_KEY_ID")
	SecretAccessKey = GetEnvWithKey("AWS_SECRET_ACCESS_KEY")

	sess, err := session.NewSession(
		&aws.Config{
			Region:      aws.String(AWSRegion),
			Credentials: credentials.NewStaticCredentials(AccessKeyID, SecretAccessKey, ""),
		})
	if err != nil {
		logrus.Fatal(err)
	}

	return &AWSStorage{Session: sess}
}

func (as AWSStorage) Upload(ctx context.Context, file multipart.File, fileName, contentType string) (string, error) {
	sess := as.Session
	uploader := s3manager.NewUploader(sess)

	BucketName := os.Getenv("BUCKET_NAME")

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(BucketName),
		ACL:         aws.String("public-read"),
		Key:         aws.String(fileName),
		ContentType: aws.String(contentType),
		Body:        file,
	})

	if err != nil {
		fmt.Println("Upload", err)
		return "", err
	}

	filepath := "https://" + BucketName + "." + "s3-" + AWSRegion + ".amazonaws.com/" + fileName

	return filepath, nil
}

func (as AWSStorage) GetSharableURL() (string, error) {
	return "", nil
}
