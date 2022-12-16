package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"github.com/tuxoo/smart-loader/loader-service/internal/config"
)

type PostgresDB struct {
	cfg  *pgxpool.Config
	pool *pgxpool.Pool
}

func NewPostgresDB(cfg *config.PostgresConfig) *PostgresDB {
	pgxConfig, err := pgxpool.ParseConfig("")
	if err != nil {
		logrus.Fatalf("parsing pgx configs error: %s", err.Error())
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

	return &PostgresDB{
		cfg: pgxConfig,
	}
}

func (p *PostgresDB) Connect(ctx context.Context) error {
	pool, err := pgxpool.ConnectConfig(ctx, p.cfg)
	if err != nil {
		logrus.Fatalf("error occured on connecting to postgres: %s", err.Error())
	} else {
		p.pool = pool
	}

	return nil
}

func (p *PostgresDB) Disconnect() {
	p.pool.Close()
}
