package server

import (
	"context"

	"github.com/yusupovanton/words-of-wisdom-POW/pkg/clog"
	"github.com/yusupovanton/words-of-wisdom-POW/pkg/metrics"
	randomChoice "github.com/yusupovanton/words-of-wisdom-POW/pkg/random_choice"
)

//go:generate ../../../bin/mockery --name storage
type storage interface {
	GetQuoteByID(id int) (string, error)
	QuotesLength() int
}

type UseCase struct {
	logger  clog.CLog
	storage storage

	registry metrics.Registry
	series   metrics.Series
}

// NewUseCase creates a new instance of UseCase with the given storage and metrics registry.
func NewUseCase(logger clog.CLog, storage storage, registry metrics.Registry) *UseCase {
	return &UseCase{
		logger:   logger,
		storage:  storage,
		registry: registry,
		series:   metrics.NewSeries(metrics.SeriesTypeUseCase, "get_random_quote"),
	}
}

// GetRandomQuote retrieves a random quote from the storage.
func (uc *UseCase) GetRandomQuote(ctx context.Context) (string, error) {
	ctx, series := uc.series.WithOperation(ctx, "get_random_quote")

	length := uc.storage.QuotesLength()
	id, err := randomChoice.RandomInt(1, length)
	if err != nil {
		uc.logger.ErrorCtx(ctx, err, "could not choose random integer")
		uc.registry.Inc(series.Error("internal"))

		return "", err
	}

	quote, err := uc.storage.GetQuoteByID(id)
	if err != nil {
		uc.logger.ErrorCtx(ctx, err, "could not get quote")
		uc.registry.Inc(series.Error("repository"))

		return "", err
	}

	uc.registry.Inc(series.Success())
	return quote, nil
}
