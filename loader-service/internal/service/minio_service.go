package service

import (
	"github.com/tuxoo/smart-loader/loader-service/internal/client"
)

type MinioService struct {
	client *client.MinioClient
}

func NewMinioService(client *client.MinioClient) *MinioService {
	return &MinioService{
		client: client,
	}
}

func (s *MinioService) Save() error {
	return nil
}

func (s *MinioService) Get() error {
	return nil
}
