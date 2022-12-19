package client

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/tuxoo/smart-loader/loader-service/internal/config"
)

type MinioClient struct {
	cfg    *config.MinioConfig
	Client *minio.Client
}

func NewMinioClient(cfg *config.MinioConfig) *MinioClient {
	return &MinioClient{
		cfg: cfg,
	}
}

func (c *MinioClient) Connect() error {
	opt := &minio.Options{
		Creds:  credentials.NewStaticV4(c.cfg.AccessKey, c.cfg.SecretKey, ""),
		Secure: true,
	}

	client, err := minio.New(c.cfg.Host, opt)
	if err != nil {
		return err
	}

	c.Client = client

	//if ok, err := client.BucketExists(ctx, model.IMAGE_BUCKET); err != nil {
	//	return err
	//} else if !ok {
	//	if err = client.MakeBucket(ctx, model.IMAGE_BUCKET, minio.MakeBucketOptions{}); err != nil {
	//		return err
	//	}
	//}

	return nil
}
