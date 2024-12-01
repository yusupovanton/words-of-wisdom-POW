package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"github.com/yusupovanton/words-of-wisdom-POW/internal/di"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	container := di.NewContainer(ctx)
	defer container.Close()

	logger := container.GetLogger()
	cfg := container.GetConfig()
	clientUseCase := container.GetClientUseCase()

	ticker := time.NewTicker(cfg.Client.Interval)
	defer ticker.Stop()

	if err := clientUseCase.FetchQuote(ctx); err != nil {
		logger.ErrorCtx(ctx, err, "error fetching quote")
	}

	for {
		select {
		case <-ctx.Done():
			logger.InfoCtx(ctx, "Shutting down client...")
			return
		case <-ticker.C:
			if err := clientUseCase.FetchQuote(ctx); err != nil {
				logger.ErrorCtx(ctx, err, "error fetching quote")
			}
		}
	}
}
