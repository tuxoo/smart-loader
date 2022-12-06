package config

import (
	"github.com/spf13/viper"
	"strings"
	"time"
)

const (
	path                      = "config/config"
	defaultHttpPort           = "9000"
	defaultHttpRWTimeout      = 10 * time.Second
	defaultMaxHeaderMegabytes = 1
)

func preDefaults() {
	viper.SetDefault("http.port", defaultHttpPort)
	viper.SetDefault("http.max_header_megabytes", defaultMaxHeaderMegabytes)
	viper.SetDefault("http.timeouts.read", defaultHttpRWTimeout)
	viper.SetDefault("http.timeouts.write", defaultHttpRWTimeout)
}

func parseConfigFile(filepath string) error {
	configPath := strings.Split(filepath, "/")

	viper.AddConfigPath(configPath[0])
	viper.SetConfigName(configPath[1])

	return viper.ReadInConfig()
}
