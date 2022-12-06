package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"github.com/tuxoo/smart-loader/facade-service/internal/config"
)

type PostgresDB struct {
	cfg  *pgxpool.Config
	pool *pgxpool.Pool
}

func NewPostgresDB(cfg *config.PostgresConfig) *PostgresDB {
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

	return &PostgresDB{
		cfg: pgxConfig,
	}
}

func (p *PostgresDB) Connect() error {
	if pool, err := pgxpool.ConnectConfig(context.Background(), p.cfg); err != nil {
		return err
	} else {
		p.pool = pool
	}

	return nil
}

func (p *PostgresDB) Disconnect() {
	p.pool.Close()
}
