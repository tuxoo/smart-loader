package service

import (
	"context"
	"github.com/tuxoo/smart-loader/loader-service/internal/domain/model"
	"github.com/tuxoo/smart-loader/loader-service/internal/domain/repository"
)

type DownloadService struct {
	repository repository.IDownloadRepository
}

func NewDownloadService(repository repository.IDownloadRepository) *DownloadService {
	return &DownloadService{
		repository: repository,
	}
}

func (s *DownloadService) GetByHash(ctx context.Context, hash string) (*model.Download, error) {
	return s.repository.FindByHash(ctx, hash)
}

func (s *DownloadService) SaveOne(ctx context.Context, download *model.Download) error {
	return s.repository.SaveOne(ctx, download)
}
