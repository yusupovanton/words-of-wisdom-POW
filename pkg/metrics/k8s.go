package metrics

import (
	"net/http"
	"sync/atomic"

	"github.com/yusupovanton/words-of-wisdom-POW/pkg/clog"
)

type HealthChecker struct {
	isReady   atomic.Value
	isHealthy atomic.Value
	logger    clog.CLog
}

func NewHealthChecker(logger clog.CLog) *HealthChecker {
	hc := &HealthChecker{
		logger: logger,
	}
	hc.isReady.Store(false)
	hc.isHealthy.Store(true)
	return hc
}

func (hc *HealthChecker) SetReady(ready bool) {
	hc.isReady.Store(ready)
}

func (hc *HealthChecker) SetHealthy(healthy bool) {
	hc.isHealthy.Store(healthy)
}

func (hc *HealthChecker) LivenessHandler(w http.ResponseWriter, r *http.Request) {
	converted, ok := hc.isHealthy.Load().(bool)
	if !ok {
		hc.logger.WarnCtx(r.Context(), "liveness response cannot be converted to bool")
		return
	}

	if converted {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("ok"))
		if err != nil {
			hc.logger.ErrorCtx(r.Context(), err, "Failed to write liveness response")
		}

		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	_, err := w.Write([]byte("unhealthy"))
	if err != nil {
		hc.logger.ErrorCtx(r.Context(), err, "Failed to write liveness response")
	}
}

func (hc *HealthChecker) ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	converted, ok := hc.isReady.Load().(bool)
	if !ok {
		hc.logger.WarnCtx(r.Context(), "readiness response cannot be converted to bool")
		return
	}

	if converted {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("ok"))
		if err != nil {
			hc.logger.ErrorCtx(r.Context(), err, "Failed to write readiness response")
		}

		return
	}

	w.WriteHeader(http.StatusServiceUnavailable)
	_, err := w.Write([]byte("not ready"))
	if err != nil {
		hc.logger.ErrorCtx(r.Context(), err, "Failed to write readiness response")
	}
}
