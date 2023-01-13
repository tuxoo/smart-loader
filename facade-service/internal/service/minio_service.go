package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/tuxoo/smart-loader/facade-service/internal/client"
)

type MinioService struct {
	minioClient *client.MinioClient
}

func NewMinioService(client *client.MinioClient) *MinioService {
	return &MinioService{
		minioClient: client,
	}
}

func (s *MinioService) Get(ctx context.Context, downloadId uuid.UUID) (*minio.Object, error) {
	return s.minioClient.Client.GetObject(ctx, "images", downloadId.String(), minio.GetObjectOptions{})
}
