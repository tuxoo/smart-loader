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

type (
	Config struct {
		HTTP     HTTPConfig
		Postgres PostgresConfig
		Nats     NatsConfig
	}

	HTTPConfig struct {
		Host               string        `mapstructure:"host"`
		Port               string        `mapstructure:"port"`
		ReadTimeout        time.Duration `mapstructure:"readTimeout"`
		WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
		MaxHeaderMegabytes int           `mapstructure:"maxHeaderMegabytes"`
	}

	NatsConfig struct {
		Port string `mapstructure:"port"`
		URL  string `mapstructure:"uri"`
	}
)

func NewConfig() (*Config, error) {
	viper.AutomaticEnv()
	preDefaults()

	if err := parseConfigFile(path); err != nil {
		return nil, err
	}

	if err := parseEnv(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshalConfig(&cfg); err != nil {
		return nil, err
	}

	setFromEnv(&cfg)

	return &cfg, nil
}

func preDefaults() {
	viper.SetDefault("http.port", defaultHttpPort)
	viper.SetDefault("http.max_header_megabytes", defaultMaxHeaderMegabytes)
	viper.SetDefault("http.timeouts.read", defaultHttpRWTimeout)
	viper.SetDefault("http.timeouts.write", defaultHttpRWTimeout)
}

func parseConfigFile(filepath string) error {
	path := strings.Split(filepath, "/")

	viper.AddConfigPath(path[0])
	viper.SetConfigName(path[1])

	return viper.ReadInConfig()
}

func parseEnv() error {
	if err := parseHttpEnv(); err != nil {
		return err
	}

	if err := parsePostgresEnv(); err != nil {
		return err
	}

	if err := parseNatsEnv(); err != nil {
		return err
	}

	return nil
}

func parseHttpEnv() error {
	if err := viper.BindEnv("http.host", "HTTP_HOST"); err != nil {
		return err
	}

	return viper.BindEnv("http.port", "HTTP_PORT")
}

func parsePostgresEnv() error {
	if err := viper.BindEnv("postgres.host", "POSTGRES_HOST"); err != nil {
		return err
	}

	if err := viper.BindEnv("postgres.port", "POSTGRES_PORT"); err != nil {
		return err
	}

	if err := viper.BindEnv("postgres.db", "POSTGRES_DB"); err != nil {
		return err
	}

	if err := viper.BindEnv("postgres.user", "POSTGRES_USER"); err != nil {
		return err
	}

	if err := viper.BindEnv("postgres.password", "POSTGRES_PASSWORD"); err != nil {
		return err
	}

	return viper.BindEnv("postgres.sslmode", "POSTGRES_SSLMODE")
}

func parseNatsEnv() error {
	if err := viper.BindEnv("nats.port", "NATS_PORT"); err != nil {
		return err
	}

	return viper.BindEnv("nats.url", "NATS_URL")
}

func unmarshalConfig(cfg *Config) error {
	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("postgres", &cfg.Postgres); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("nats", &cfg.Nats); err != nil {
		return err
	}

	return nil
}

func setFromEnv(cfg *Config) {
	cfg.HTTP.Host = viper.GetString("http.host")
	cfg.HTTP.Port = viper.GetString("http.port")

	cfg.Postgres.Host = viper.GetString("postgres.host")
	cfg.Postgres.Port = viper.GetUint("postgres.port")
	cfg.Postgres.DB = viper.GetString("postgres.db")
	cfg.Postgres.User = viper.GetString("postgres.user")
	cfg.Postgres.Password = viper.GetString("postgres.password")

	cfg.Nats.URL = viper.GetString("nats.url")
	cfg.Nats.Port = viper.GetString("nats.port")
}
