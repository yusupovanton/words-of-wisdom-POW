package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/yusupovanton/go-service-template/internal/config"
)

func NewPostgres(ctx context.Context, cfg config.Config) (*pgxpool.Pool, error) {
	pool, err := initDb(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return pool, nil
}

func initDb(ctx context.Context, cfg config.Config) (*pgxpool.Pool, error) {
	poolConfig, err := pgxpool.ParseConfig(cfg.DBConfig.Conn)
	if err != nil {
		return nil, fmt.Errorf("unable to parse database connection string: %w", err)
	}

	pool, err := pgxpool.ConnectConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	if err = pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}

	return pool, nil
}
