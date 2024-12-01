package di

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/yusupovanton/go-service-template/internal/config"
	"github.com/yusupovanton/go-service-template/pkg/clog"
	"github.com/yusupovanton/go-service-template/pkg/metrics"
	"github.com/yusupovanton/go-service-template/pkg/postgres"
)

type Container struct {
	ctx    context.Context
	logger clog.CLog
	cfg    config.Config

	metricsServer metrics.Server
	registry      metrics.Registry
	dbPool        *pgxpool.Pool

	closeFns []func()
}

func NewContainer(ctx context.Context) *Container {
	cfg := config.MustNew()
	logger := clog.NewCustomLogger(cfg)
	registry := metrics.NewRegistry(cfg)

	return &Container{
		ctx:      ctx,
		cfg:      cfg,
		logger:   logger,
		registry: registry,
		closeFns: make([]func(), 0),
	}
}

func (c *Container) GetConfig() config.Config {
	//nolint:gocritic
	return get(&c.cfg, func() config.Config {
		return config.MustNew()
	})
}

func (c *Container) GetLogger() clog.CLog {
	return get(&c.logger, func() clog.CLog {
		return clog.NewCustomLogger(c.GetConfig())
	})
}

func (c *Container) GetMetricsRegistry() metrics.Registry {
	return get(&c.registry, func() metrics.Registry {
		return metrics.NewRegistry(c.GetConfig())
	})
}

func (c *Container) GetMetricsServer() metrics.Server {
	return get(&c.metricsServer, func() metrics.Server {
		server := metrics.NewServer(c.GetLogger(), c.GetConfig(), c.GetMetricsRegistry(), metrics.NewHealthChecker(c.GetLogger()))
		c.addCloseFn(func() {
			if err := server.Stop(c.ctx); err != nil {
				c.GetLogger().ErrorCtx(c.ctx, err, "could not stop server")
			}
		})
		return server
	})
}

func (c *Container) GetPostgres() *pgxpool.Pool {
	return get(&c.dbPool, func() *pgxpool.Pool {
		dbPool, err := postgres.NewPostgres(c.ctx, c.cfg)
		if err != nil {
			c.GetLogger().ErrorCtx(c.ctx, err, "Failed to initialize PostgreSQL")
		}
		c.addCloseFn(func() { dbPool.Close() })
		return dbPool
	})
}

func (c *Container) addCloseFn(fn func()) {
	c.closeFns = append(c.closeFns, fn)
}

func (c *Container) Close() {
	var wg sync.WaitGroup
	for _, fn := range c.closeFns {
		wg.Add(1)
		go func(fn func()) {
			defer wg.Done()
			fn()
		}(fn)
	}
	wg.Wait()
}

func get[T comparable](obj *T, builder func() T) T {
	if *obj != *new(T) {
		return *obj
	}
	*obj = builder()
	return *obj
}
