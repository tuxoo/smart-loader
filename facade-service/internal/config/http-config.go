package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

//const (
//	path                      = "config/config"
//	defaultHttpPort           = "9000"
//	defaultHttpRWTimeout      = 10 * time.Second
//	defaultMaxHeaderMegabytes = 1
//)

type HTTPConfig struct {
	Host               string `mapstructure:"HTTP_HOST"`
	Port               string `mapstructure:"HTTP_PORT"`
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	MaxHeaderMegabytes int
}

func NewHTTPConfig() (cfg *HTTPConfig) {
	viper.AutomaticEnv()
	preDefaults()

	if err := parseConfigFile(path); err != nil {
		logrus.Fatalf("parsing configs error: %s", err.Error())
	}

	if err := cfg.parseEnv(); err != nil {
		logrus.Fatalf("parsing http .env error: %s", err.Error())
	}

	if err := viper.UnmarshalKey("http", &cfg); err != nil {
		logrus.Fatalf("unmarshaling http configs error: %s", err.Error())
	}

	cfg.Host = viper.GetString("http.host")
	cfg.Port = viper.GetString("http.port")

	return
}

func (c *HTTPConfig) parseEnv() error {
	if err := viper.BindEnv("http.host", "HTTP_HOST"); err != nil {
		return err
	}

	return viper.BindEnv("http.port", "HTTP_PORT")
}
