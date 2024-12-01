package di

import (
	"context"
	"sync"

	"github.com/yusupovanton/words-of-wisdom-POW/internal/config"
	"github.com/yusupovanton/words-of-wisdom-POW/internal/handlers"
	"github.com/yusupovanton/words-of-wisdom-POW/internal/repository"
	"github.com/yusupovanton/words-of-wisdom-POW/internal/server"
	"github.com/yusupovanton/words-of-wisdom-POW/internal/usecase"
	"github.com/yusupovanton/words-of-wisdom-POW/pkg/clog"
	"github.com/yusupovanton/words-of-wisdom-POW/pkg/metrics"
)

type Container struct {
	ctx    context.Context
	logger clog.CLog
	cfg    config.Config

	metricsServer metrics.Server
	registry      metrics.Registry

	repo    *repository.Repository
	useCase *usecase.UseCase
	handler *handlers.GetQuoteHandler
	server  *server.Server

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
		mSrv := metrics.NewServer(c.GetLogger(), c.GetConfig(), c.GetMetricsRegistry(), metrics.NewHealthChecker(c.GetLogger()))
		c.addCloseFn(func() {
			if err := mSrv.Stop(c.ctx); err != nil {
				c.GetLogger().ErrorCtx(c.ctx, err, "could not stop server")
			}
		})
		return mSrv
	})
}

func (c *Container) GetRepository() *repository.Repository {
	return get(&c.repo, func() *repository.Repository {
		return repository.New()
	})
}

func (c *Container) GetUseCase() *usecase.UseCase {
	return get(&c.useCase, func() *usecase.UseCase {
		return usecase.NewUseCase(c.GetRepository())
	})
}

func (c *Container) GetQuoteHandler() *handlers.GetQuoteHandler {
	return get(&c.handler, func() *handlers.GetQuoteHandler {
		return handlers.NewGetQuoteHandler(
			c.GetLogger(),
			c.GetMetricsRegistry(),
			c.GetUseCase(),
			c.GetConfig().POW.Complexity,
			c.GetConfig().POW.Prefix,
		)
	})
}

func (c *Container) GetServer() *server.Server {
	return get(&c.server, func() *server.Server {
		return server.NewServer(
			c.GetConfig().TCPServer.Port,
			c.GetLogger(),
			c.GetQuoteHandler(),
		)
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