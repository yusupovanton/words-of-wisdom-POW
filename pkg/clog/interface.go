package clog

import "context"

type CLog interface {
	AddKeysValuesToCtx(ctx context.Context, kv map[string]interface{}) context.Context
	GetFieldByKey(ctx context.Context, key string) (interface{}, bool)
	ErrorCtx(ctx context.Context, err error, msg string, args ...any)
	InfoCtx(ctx context.Context, msg string, args ...any)
	DebugCtx(ctx context.Context, msg string, args ...any)
	WarnCtx(ctx context.Context, msg string, args ...any)
}
