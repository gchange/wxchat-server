package middleware

import (
	"time"
	"wx-server/internal/logging"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// WrapTimer times the process of the request
func WrapTimer() gin.HandlerFunc {
	return func(c *gin.Context) {
		l := logging.WithContext(c)
		l = l.With(
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("remote addr", c.Request.RemoteAddr),
		)
		start := time.Now()
		l.Info("request in")
		c.Next()
		l.With(
			zap.Int("status", c.Writer.Status()),
			zap.Int("size", c.Writer.Size()),
			zap.Int64("duration(ms)", time.Since(start).Milliseconds()),
		).Info("request out")
	}
}
