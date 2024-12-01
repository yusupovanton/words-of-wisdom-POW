package metrics

import (
	"context"
	"net/http"
	"runtime"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/yusupovanton/go-service-template/internal/config"
	"github.com/yusupovanton/go-service-template/pkg/clog"
)

const (
	metricsEndpoint   = "/metrics"
	livenessEndpoint  = "/healthz"
	readinessEndpoint = "/readyz"
)

var (
	memAllocGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "go_mem_stats_alloc_bytes",
		Help: "Number of bytes allocated and still in use.",
	})
	memSysGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "go_mem_stats_sys_bytes",
		Help: "Number of bytes obtained from the system.",
	})
	memHeapAllocGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "go_mem_stats_heap_alloc_bytes",
		Help: "Number of heap bytes allocated and still in use.",
	})
	memHeapSysGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "go_mem_stats_heap_sys_bytes",
		Help: "Number of heap bytes obtained from the system.",
	})
)

type server struct {
	logger      clog.CLog
	cfg         config.Config
	registry    Registry
	healthCheck *HealthChecker
	httpServer  *http.Server
	stopCh      chan struct{}
}

func NewServer(logger clog.CLog, cfg config.Config, registry Registry, healthCheck *HealthChecker) Server {
	registry.PrometheusRegistry().MustRegister(memAllocGauge, memSysGauge, memHeapAllocGauge, memHeapSysGauge)

	return &server{
		logger:      logger,
		cfg:         cfg,
		registry:    registry,
		healthCheck: healthCheck,
		stopCh:      make(chan struct{}),
	}
}

func (s *server) Start(ctx context.Context) {
	mux := http.NewServeMux()

	mux.Handle(metricsEndpoint, promhttp.HandlerFor(s.registry.PrometheusRegistry(), promhttp.HandlerOpts{}))
	mux.HandleFunc(livenessEndpoint, s.healthCheck.LivenessHandler)
	mux.HandleFunc(readinessEndpoint, s.healthCheck.ReadinessHandler)

	ctx = s.logger.AddKeysValuesToCtx(ctx, map[string]interface{}{
		"metrics_address": s.cfg.Metrics.Address,
	})

	s.httpServer = &http.Server{
		Addr:              s.cfg.Metrics.Address,
		Handler:           mux,
		ReadHeaderTimeout: 1 * time.Second,
		ReadTimeout:       2 * time.Second,
		WriteTimeout:      2 * time.Second,
		IdleTimeout:       30 * time.Second,
	}

	go s.collectMemoryStats(ctx)

	go func() {
		s.logger.InfoCtx(ctx, "starting metrics server, address: %s", s.cfg.Metrics.Address)
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.ErrorCtx(ctx, err, "failed to start metrics server")
		}
	}()
}

func (s *server) Stop(ctx context.Context) error {
	close(s.stopCh)

	if s.httpServer != nil {
		err := s.httpServer.Shutdown(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *server) collectMemoryStats(ctx context.Context) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			s.logger.InfoCtx(ctx, "stopping memory stats collection")
			return
		case <-s.stopCh:
			s.logger.InfoCtx(ctx, "memory stats collection stopped by server stop")
			return
		case <-ticker.C:
			var memStats runtime.MemStats
			runtime.ReadMemStats(&memStats)

			memAllocGauge.Set(float64(memStats.Alloc))
			memSysGauge.Set(float64(memStats.Sys))
			memHeapAllocGauge.Set(float64(memStats.HeapAlloc))
			memHeapSysGauge.Set(float64(memStats.HeapSys))
		}
	}
}
