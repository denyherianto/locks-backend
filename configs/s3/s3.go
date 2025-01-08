package s3

import (
	"errors"
	"os"

	"github.com/gofiber/storage/s3/v2"
)

// NewS3Client initializes an S3 client
func OpenS3Client() (*s3.Storage, error) {
	storage := s3.New(s3.Config{
		Endpoint: os.Getenv("S3_ENDPOINT"),
		Region:   os.Getenv("S3_REGION"),
		Bucket:   os.Getenv("S3_BUCKET"),
		Credentials: s3.Credentials{
			AccessKey:       os.Getenv("S3_ACCESSKEY"),
			SecretAccessKey: os.Getenv("S3_SECRETKEY"),
		},
	})

	if storage == nil {
		return nil, errors.New("failed to initialize S3 storage")
	}

	return storage, nil
}
