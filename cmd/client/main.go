package main

import (
	"context"
	"os"
	"os/signal"

	_ "github.com/joho/godotenv/autoload"

	"github.com/yusupovanton/words-of-wisdom-POW/internal/di"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	container := di.NewContainer(ctx)
	defer container.Close()

	logger := container.GetLogger()

	clientUseCase := container.GetClientUseCase()

	if err := clientUseCase.FetchQuote(ctx); err != nil {
		logger.ErrorCtx(ctx, err, "error fetching quote")
	}
}
