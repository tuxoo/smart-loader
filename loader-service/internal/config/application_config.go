package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type AppConfig struct {
	HashSalt string
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

	return
}

func (c *AppConfig) parseEnv() error {
	return viper.BindEnv("app.hashSalt", "APP_HASH_SALT")
}
