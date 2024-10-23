package configs

import (
	"io"
	"time"

	"github.com/minio/minio-go"
)

type Minio struct {
	Client *minio.Client
}

func NewMinio(useSsl bool, secret, accessKey, endPoint string) *Minio {
	minioClient, err := minio.New(endPoint, accessKey, secret, useSsl)
	if err != nil {

	}
	return &Minio{
		Client: minioClient,
	}
}

func (m *Minio) UploadFile(bucketName, objectName string, reader io.Reader, objectSize int64,
	opts minio.PutObjectOptions) error {
	_, err := m.Client.PutObject(bucketName, objectName, reader, objectSize,
		opts)
	return err
}

func (m *Minio) GetPresignedURL(bucketName, objectName string, expires time.Duration) (string, error) {
	url, err := m.Client.PresignedGetObject(bucketName, objectName, expires, nil)
	if err != nil {
		return "", err
	}
	return url.String(), err
}

func (m *Minio) DeleteFile(bucketName, objectName string) error {
	return m.Client.RemoveObject(bucketName, objectName)
}
