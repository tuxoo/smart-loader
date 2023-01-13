package service

import (
	"bytes"
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/tuxoo/smart-loader/loader-service/internal/client"
	"github.com/tuxoo/smart-loader/loader-service/internal/domain/model"
	"github.com/tuxoo/smart-loader/loader-service/internal/domain/model/const"
)

type MinioService struct {
	minioClient *client.MinioClient
}

func NewMinioService(client *client.MinioClient) *MinioService {
	return &MinioService{
		minioClient: client,
	}
}

func (s *MinioService) Put(ctx context.Context, content []byte, download *model.Download) error {
	if _, err := s.minioClient.Client.PutObject(
		ctx,
		_const.IMAGE_BUCKET,
		download.Id.String(),
		bytes.NewReader(content),
		int64(download.Size),
		minio.PutObjectOptions{},
	); err != nil {
		return err
	}

	return nil
}

func (s *MinioService) Get() error {
	return nil
}
