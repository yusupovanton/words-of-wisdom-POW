package usecase

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/yusupovanton/words-of-wisdom-POW/pkg/clog"
	"github.com/yusupovanton/words-of-wisdom-POW/pkg/metrics"
)

type tcpClient interface {
	Send(message string) error
	Receive() (string, error)
}

type QuoteUseCase struct {
	client   tcpClient
	logger   clog.CLog
	registry metrics.Registry
	series   metrics.Series
}

func NewQuoteUseCase(client tcpClient, logger clog.CLog, registry metrics.Registry) *QuoteUseCase {
	return &QuoteUseCase{
		client:   client,
		logger:   logger,
		registry: registry,
		series:   metrics.NewSeries(metrics.SeriesTypeUseCase, "fetch_quote"),
	}
}

func (uc *QuoteUseCase) FetchQuote(ctx context.Context) error {
	ctx, series := uc.series.WithOperation(ctx, "fetch_quote")

	challengeMessage, err := uc.client.Receive()
	if err != nil {
		uc.registry.Inc(series.Error("receive_challenge"))
		return fmt.Errorf("failed to receive challenge: %w", err)
	}
	uc.logger.InfoCtx(ctx, "Received challenge", "message", challengeMessage)

	parts := strings.Split(challengeMessage, "\n")
	if len(parts) < 2 {
		uc.registry.Inc(series.Error("invalid_challenge"))
		return fmt.Errorf("invalid challenge format")
	}
	prefix := strings.TrimSpace(parts[0])
	difficulty := len(strings.TrimSpace(parts[1]))

	nonce := uc.solveChallenge(prefix, difficulty)

	err = uc.client.Send(nonce)
	if err != nil {
		uc.registry.Inc(series.Error("send_nonce"))
		return fmt.Errorf("failed to send nonce: %w", err)
	}
	uc.logger.InfoCtx(ctx, "Submitted nonce", "nonce", nonce)

	quote, err := uc.client.Receive()
	if err != nil {
		uc.registry.Inc(series.Error("receive_quote"))
		return fmt.Errorf("failed to receive quote: %w", err)
	}
	uc.logger.InfoCtx(ctx, "Received quote", "quote", quote)

	return nil
}

func (uc *QuoteUseCase) solveChallenge(prefix string, difficulty int) string {
	requiredPrefix := strings.Repeat("0", difficulty)
	nonce := 0
	for {
		data := fmt.Sprintf("%s%d", prefix, nonce)
		hash := sha256.Sum256([]byte(data))
		hashHex := hex.EncodeToString(hash[:])
		if strings.HasPrefix(hashHex, requiredPrefix) {
			return fmt.Sprintf("%d", nonce)
		}
		nonce++
	}
}
