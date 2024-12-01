package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	_ "github.com/joho/godotenv/autoload"

	"github.com/yusupovanton/go-service-template/internal/di"
	"github.com/yusupovanton/go-service-template/pkg/metrics"
)

const (
	successExitCode = 0
	failExitCode    = 1
)

func main() {
	os.Exit(run())
}

func run() int {
	var err error

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	c := di.NewContainer(ctx)
	defer c.Close()

	cfg := c.GetConfig()
	logger := c.GetLogger()
	metricRegistry := c.GetMetricsRegistry()

	defer func() {
		if panicErr := recover(); panicErr != nil {
			err = fmt.Errorf("panic recovered: %v", panicErr)
			logger.ErrorCtx(ctx, err, "panic recovered!")
		}
	}()

	metricsServer := metrics.NewServer(logger, cfg, metricRegistry, metrics.NewHealthChecker(logger))
	defer func() {
		err = metricsServer.Stop(ctx)
		if err != nil {
			logger.ErrorCtx(ctx, err, "failed to stop metrics server")
		}
	}()

	go func() {
		metricsServer.Start(ctx)
	}()

	dbPool := c.GetPostgres()
	defer func() {
		dbPool.Close()
		logger.InfoCtx(ctx, "stopped database connection gracefully")
	}()

	blockUntilContextCancelled(ctx)

	return successExitCode
}

func blockUntilContextCancelled(ctx context.Context) {
	<-ctx.Done()
}
