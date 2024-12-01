package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	_ "github.com/joho/godotenv/autoload"

	"github.com/yusupovanton/words-of-wisdom-POW/internal/di"
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

	logger := c.GetLogger()
	server := c.GetServer()

	defer func() {
		if panicErr := recover(); panicErr != nil {
			err = fmt.Errorf("panic recovered: %v", panicErr)
			logger.ErrorCtx(ctx, err, "panic recovered!")
		}
	}()

	go func() {
		if err = server.Run(ctx); err != nil {
			logger.ErrorCtx(ctx, err, "server encountered an error")
			cancel()
		}
	}()

	blockUntilContextCancelled(ctx)

	return successExitCode
}

func blockUntilContextCancelled(ctx context.Context) {
	<-ctx.Done()
}
