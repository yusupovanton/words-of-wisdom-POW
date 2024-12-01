package usecase

import (
	"context"
	"fmt"

	"github.com/yusupovanton/words-of-wisdom-POW/pkg/clog"
	"github.com/yusupovanton/words-of-wisdom-POW/pkg/metrics"
)

//go:generate ../../../bin/mockery --name quoteGetterClient
type quoteGetterClient interface {
	GetQuote(ctx context.Context) (string, error)
}

type QuoteUseCase struct {
	client   quoteGetterClient
	logger   clog.CLog
	registry metrics.Registry
	series   metrics.Series
}

func NewQuoteUseCase(client quoteGetterClient, logger clog.CLog, registry metrics.Registry) *QuoteUseCase {
	return &QuoteUseCase{
		client:   client,
		logger:   logger,
		registry: registry,
		series:   metrics.NewSeries(metrics.SeriesTypeUseCase, "fetch_quote"),
	}
}

func (uc *QuoteUseCase) FetchQuote(ctx context.Context) error {
	ctx, series := uc.series.WithOperation(ctx, "fetch_quote")

	quote, err := uc.client.GetQuote(ctx)
	if err != nil {
		uc.registry.Inc(series.Error("get_quote"))
		uc.logger.ErrorCtx(ctx, err, "Failed to get quote")
		return fmt.Errorf("failed to get quote: %w", err)
	}

	uc.logger.InfoCtx(ctx, fmt.Sprintf("fetched quote successfully: %s", quote))
	uc.registry.Inc(series.Success())
	return nil
}
