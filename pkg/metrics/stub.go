package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type registryStub struct{}

func NewRegistryStub() Registry {
	return &registryStub{}
}

func (s *registryStub) Inc(_ string, _ prometheus.Labels) {}

func (s *registryStub) RecordDuration(_ string, _ prometheus.Labels, _ float64) {}

func (s *registryStub) PrometheusRegistry() *prometheus.Registry {
	return nil
}
