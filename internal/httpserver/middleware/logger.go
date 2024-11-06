package middleware

import (
	"wx-server/internal/logging"
	"wx-server/internal/random"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	XTraceId = "X-Trace-Id"
)

// WrapLogger wraps logger with trace id
func WrapLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := c.GetHeader(XTraceId)
		if traceID == "" {
			traceID = genTraceId()
		}
		l := logging.WithContext(c)
		zl := l.With(zap.String("X-Trace-Id", traceID))
		zl.SetLogger(zl.GetLogger().WithOptions(zap.AddCallerSkip(-1)))

		c.Set(string(logging.ContextLog), zl)
		c.Next()
	}
}

func genTraceId() string {
	return random.RandString(21)
}
