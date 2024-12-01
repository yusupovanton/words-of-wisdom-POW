package metrics

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
)

// Registry defines the interface for managing metrics.
type Registry interface {
	Inc(metricName string, labels prometheus.Labels)
	RecordDuration(metricName string, labels prometheus.Labels, duration float64)
	PrometheusRegistry() *prometheus.Registry
}

type Server interface {
	Start(ctx context.Context)
	Stop(ctx context.Context) error
}
