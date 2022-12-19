package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type MinioConfig struct {
	Host      string
	AccessKey string
	SecretKey string
}

func NewMinioConfig() (cfg *MinioConfig) {
	viper.AutomaticEnv()
	preDefaults()

	if err := parseConfigFile(path); err != nil {
		logrus.Fatalf("parsing configs error: %s", err.Error())
	}

	if err := cfg.parseEnv(); err != nil {
		logrus.Fatalf("parsing app .env error: %s", err.Error())
	}

	if err := viper.UnmarshalKey("minio", &cfg); err != nil {
		logrus.Fatalf("unmarshaling app configs error: %s", err.Error())
	}

	cfg.Host = viper.GetString("minio.host")
	cfg.AccessKey = viper.GetString("minio.accessKey")
	cfg.SecretKey = viper.GetString("minio.secretKey")

	return
}

func (c *MinioConfig) parseEnv() error {
	if err := viper.BindEnv("minio.host", "MINIO_HOST"); err != nil {
		return err
	}
	if err := viper.BindEnv("minio.accessKey", "MINIO_ACCESS_KEY"); err != nil {
		return err
	}

	return viper.BindEnv("minio.secretKey", "MINIO_SECRET_KEY")
}
