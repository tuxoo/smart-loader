package config

import (
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

type PostgresConfig struct {
	Host            string
	Port            uint
	DB              string
	User            string
	Password        string
	MaxConns        int32
	MinConns        int32
	MaxConnLifetime time.Duration
	MaxConnIdleTime time.Duration
}

func NewPostgresConfig() (cfg *PostgresConfig) {
	viper.AutomaticEnv()

	if err := parseConfigFile(path); err != nil {
		logrus.Fatalf("parsing configs error: %s", err.Error())
	}

	if err := cfg.parseEnv(); err != nil {
		logrus.Fatalf("parsing .env error: %s", err.Error())
	}

	if err := viper.UnmarshalKey("postgres", &cfg); err != nil {
		logrus.Fatalf("unmarshaling configs error: %s", err.Error())
	}

	cfg.Host = viper.GetString("postgres.host")
	cfg.Port = viper.GetUint("postgres.port")
	cfg.DB = viper.GetString("postgres.db")
	cfg.User = viper.GetString("postgres.user")
	cfg.Password = viper.GetString("postgres.password")

	return
}

func (c *PostgresConfig) parseEnv() error {
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
