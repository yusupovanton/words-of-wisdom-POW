package clog

import "context"

type fieldMapType struct{}

var fieldMap = fieldMapType{}

func (l *CustomLogger) AddKeysValuesToCtx(ctx context.Context, kv map[string]interface{}) context.Context {
	fields := ctx.Value(fieldMap)

	if fields == nil {
		return context.WithValue(ctx, fieldMap, kv)
	}

	l.mu.Lock()

	converted, ok := fields.(map[string]interface{})
	if !ok {
		return ctx
	}

	for k, v := range kv {
		if v != nil {
			converted[k] = v
		}
	}

	l.mu.Unlock()

	return context.WithValue(ctx, fieldMap, fields)
}

func (l *CustomLogger) fieldsFromCtx(ctx context.Context) map[string]interface{} {
	fm := ctx.Value(fieldMap)
	if fm == nil {
		return nil
	}

	converted, ok := fm.(map[string]interface{})
	if !ok {
		return nil
	}

	return converted
}

// GetFieldByKey retrieves a specific value by key from the fields in the context
func (l *CustomLogger) GetFieldByKey(ctx context.Context, key string) (interface{}, bool) {
	fields := l.fieldsFromCtx(ctx)
	if fields == nil {
		return nil, false
	}

	value, exists := fields[key]
	return value, exists
}
