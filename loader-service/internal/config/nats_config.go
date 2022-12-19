package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type NatsConfig struct {
	Port string `mapstructure:"NATS_PORT"`
	HOST string `mapstructure:"NATS_HOST"`
}

func NewNatsConfig() (cfg *NatsConfig) {
	viper.AutomaticEnv()

	if err := parseConfigFile(path); err != nil {
		logrus.Fatalf("parsing configs error: %s", err.Error())
	}

	if err := cfg.parseEnv(); err != nil {
		logrus.Fatalf("parsing .env error: %s", err.Error())
	}

	if err := viper.UnmarshalKey("nats", &cfg); err != nil {
		logrus.Fatalf("unmarshaling configs error: %s", err.Error())
	}

	cfg.HOST = viper.GetString("nats.host")
	cfg.Port = viper.GetString("nats.port")

	return
}

func (c *NatsConfig) parseEnv() error {
	if err := viper.BindEnv("nats.port", "NATS_PORT"); err != nil {
		return err
	}

	return viper.BindEnv("nats.host", "NATS_HOST")
}
