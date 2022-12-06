package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"github.com/tuxoo/smart-loader/facade-service/internal/config"
)

func NewPostgresPool(cfg *config.PostgresConfig) *pgxpool.Pool {
	pgxConfig, err := pgxpool.ParseConfig("")
	if err != nil {
		logrus.Fatalf("parsing postgres configs error: %s", err.Error())
	}

	pgxConfig.ConnConfig.Host = cfg.Host
	pgxConfig.ConnConfig.Port = uint16(cfg.Port)
	pgxConfig.ConnConfig.Database = cfg.DB
	pgxConfig.ConnConfig.User = cfg.User
	pgxConfig.ConnConfig.Password = cfg.Password

	pgxConfig.MaxConns = cfg.MaxConns
	pgxConfig.MinConns = cfg.MinConns
	pgxConfig.MaxConnLifetime = cfg.MaxConnLifetime
	pgxConfig.MaxConnIdleTime = cfg.MaxConnIdleTime

	pool, err := pgxpool.ConnectConfig(context.Background(), pgxConfig)
	if err != nil {
		logrus.Fatalf("postgres initializing error: %s", err.Error())
	}
	return pool
}
