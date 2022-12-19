package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type NatsConfig struct {
	Port string `mapstructure:"NATS_PORT"`
	URL  string `mapstructure:"NATS_URL"`
}

func NewNatsConfig() (cfg *NatsConfig) {
	viper.AutomaticEnv()

	if err := parseConfigFile(path); err != nil {
		logrus.Fatalf("parsing configs error: %s", err.Error())
	}

	if err := cfg.parseEnv(); err != nil {
		logrus.Fatalf("parsing nats .env error: %s", err.Error())
	}

	if err := viper.UnmarshalKey("nats", &cfg); err != nil {
		logrus.Fatalf("unmarshaling nats configs error: %s", err.Error())
	}

	cfg.URL = viper.GetString("nats.url")
	cfg.Port = viper.GetString("nats.port")

	return
}

func (c *NatsConfig) parseEnv() error {
	if err := viper.BindEnv("nats.port", "NATS_PORT"); err != nil {
		return err
	}

	return viper.BindEnv("nats.url", "NATS_URL")
}
