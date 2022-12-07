package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type AppConfig struct {
	HashSalt         string
	UriPartitionSize int
}

func NewAppConfig() (cfg *AppConfig) {
	viper.AutomaticEnv()
	preDefaults()

	if err := parseConfigFile(path); err != nil {
		logrus.Fatalf("parsing configs error: %s", err.Error())
	}

	if err := cfg.parseEnv(); err != nil {
		logrus.Fatalf("parsing app .env error: %s", err.Error())
	}

	if err := viper.UnmarshalKey("app", &cfg); err != nil {
		logrus.Fatalf("unmarshaling app configs error: %s", err.Error())
	}

	cfg.HashSalt = viper.GetString("app.hashSalt")
	cfg.UriPartitionSize = viper.GetInt("app.uriPartitionSize")

	return
}

func (c *AppConfig) parseEnv() error {
	if err := viper.BindEnv("app.hashSalt", "APP_HASH_SALT"); err != nil {
		return err
	}

	return viper.BindEnv("app.uriPartitionSize", "APP_URI_PARTITION_SIZE")
}
