package service

import (
	"archive/zip"
	"bytes"
	"context"
	"github.com/google/uuid"
	"github.com/tuxoo/smart-loader/facade-service/internal/domain/repository"
	"io"
)

type DownloadService struct {
	repository   repository.IDownloadRepository
	minioService IMinioService
}

func NewDownloadService(repository repository.IDownloadRepository, minioService IMinioService) *DownloadService {
	return &DownloadService{
		repository:   repository,
		minioService: minioService,
	}
}

func (s *DownloadService) GetDownloadZip(ctx context.Context, jobId uuid.UUID, userId int) ([]byte, error) {
	downloads, err := s.repository.FindAllByJobId(ctx, jobId, userId)
	if err != nil {
		return nil, err
	}

	bytesBuffer := new(bytes.Buffer)
	zipWriter := zip.NewWriter(bytesBuffer)

	for _, download := range downloads {
		minioObject, err := s.minioService.Get(ctx, download.Id)
		if err != nil {
			return nil, err
		}

		fileWriter, _ := zipWriter.Create(uuid.NewString() + ".jpg")
		if _, err = io.Copy(fileWriter, minioObject); err != nil {
			return nil, err
		}

		if err = minioObject.Close(); err != nil {
			return nil, err
		}
	}

	zipWriter.Close()

	return bytesBuffer.Bytes(), nil
}
