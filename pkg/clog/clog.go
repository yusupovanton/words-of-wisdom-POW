package clog

import (
	"context"
	"fmt"
	"log/slog"
	"sync"

	"github.com/yusupovanton/go-service-template/internal/config"
)

func NewCustomLogger(cfg config.Config) CLog {
	return &CustomLogger{
		Logger: slog.New(
			slog.NewJSONHandler(
				cfg.Log.Dest,
				&slog.HandlerOptions{
					AddSource: cfg.Log.AddSource,
					Level:     cfg.Log.Level,
				})),
		mu: &sync.RWMutex{},
	}
}

type CustomLogger struct {
	*slog.Logger

	mu *sync.RWMutex
}

// ErrorCtx logs an error message with fmt.SprintF()
func (l *CustomLogger) ErrorCtx(ctx context.Context, err error, msg string, args ...any) {
	l.With(convertToAttrs(l.fieldsFromCtx(ctx))...).With(slog.String("error", err.Error())).ErrorContext(ctx, fmt.Sprintf(msg, args...))
}

// InfoCtx logs an informational message with fmt.SprintF()
func (l *CustomLogger) InfoCtx(ctx context.Context, msg string, args ...any) {
	l.With(convertToAttrs(l.fieldsFromCtx(ctx))...).InfoContext(ctx, fmt.Sprintf(msg, args...))
}

// DebugCtx logs a debug message with fmt.SprintF()
func (l *CustomLogger) DebugCtx(ctx context.Context, msg string, args ...any) {
	l.With(convertToAttrs(l.fieldsFromCtx(ctx))...).DebugContext(ctx, fmt.Sprintf(msg, args...))
}

// WarnCtx logs a debug message with fmt.SprintF()
func (l *CustomLogger) WarnCtx(ctx context.Context, msg string, args ...any) {
	l.With(convertToAttrs(l.fieldsFromCtx(ctx))...).WarnContext(ctx, fmt.Sprintf(msg, args...))
}

// convertToAttrs converts a map of custom fields to a slice of slog.Attr
func convertToAttrs(fields map[string]interface{}) []any {
	attrs := make([]any, len(fields))

	i := 0
	for k, v := range fields {
		attrs[i] = slog.Any(k, v)
		i++
	}

	return attrs
}
