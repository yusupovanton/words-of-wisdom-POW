package clog_test

import (
	"bytes"
	"context"
	"log/slog"
	"testing"
	"time"

	"github.com/yusupovanton/words-of-wisdom-POW/internal/config"
	"github.com/yusupovanton/words-of-wisdom-POW/pkg/clog"
)

func BenchmarkCustomLogger(b *testing.B) {
	var buf bytes.Buffer

	cfg := config.Config{
		Log: &config.Log{
			Level:     slog.LevelDebug,
			Dest:      &buf,
			AddSource: false,
		},
	}

	logger := clog.NewCustomLogger(cfg)

	ctx := logger.AddKeysValuesToCtx(context.Background(), map[string]interface{}{
		"userID":    12345,
		"userName":  "testuser",
		"timestamp": time.Now(),
		"data":      []int{1, 2, 3},
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.InfoCtx(ctx, "Some test message")
	}
}
