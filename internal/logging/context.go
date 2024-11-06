package logging

import (
	"context"
)

type logKey string

const ContextLog logKey = "contextLog"

func WithValue(ctx context.Context, logger *ZapLogger) context.Context {
	return context.WithValue(ctx, ContextLog, logger)
}

func WithContext(ctx context.Context) *ZapLogger {
	v := ctx.Value(string(ContextLog))
	if v == nil {
		return DefaultLogger
	}
	cl, ok := v.(*ZapLogger)
	if ok {
		return cl
	}
	return DefaultLogger
}
