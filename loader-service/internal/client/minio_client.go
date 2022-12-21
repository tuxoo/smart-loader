package client

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/tuxoo/smart-loader/loader-service/internal/config"
	"github.com/tuxoo/smart-loader/loader-service/internal/domain/model"
)

type MinioClient struct {
	cfg     *config.MinioConfig
	Client  *minio.Client
	buckets map[string]string
}

func NewMinioClient(cfg *config.MinioConfig) *MinioClient {
	return &MinioClient{
		cfg: cfg,
	}
}

func (c *MinioClient) Connect(ctx context.Context) error {
	opt := &minio.Options{
		Creds:  credentials.NewStaticV4(c.cfg.AccessKey, c.cfg.SecretKey, ""),
		Secure: false,
	}

	client, err := minio.New(
		fmt.Sprintf("%s:%s", c.cfg.Host, c.cfg.Port),
		opt,
	)

	if err != nil {
		return err
	}

	c.Client = client

	if ok, err := client.BucketExists(ctx, model.IMAGE_BUCKET); err != nil {
		return err
	} else if !ok {
		if err = client.MakeBucket(ctx, model.IMAGE_BUCKET, minio.MakeBucketOptions{}); err != nil {
			return err
		}
	}

	return nil
}
