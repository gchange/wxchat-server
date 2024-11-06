package ginmiddleware

import (
	"fmt"
	"wx-server/internal/logging"

	"github.com/gin-gonic/gin"
)

type logger struct{}

func (l logger) Write(p []byte) (n int, err error) {
	logging.Info(string(p))
	return len(p), nil
}

func Logger() gin.HandlerFunc {
	formatter := func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("[GIN] %v |%3d| %13v | %15s | %s_%s:%s |%-7s %#v\n%s",
			params.TimeStamp.Format("2006/01/02 - 15:04:05"),
			params.StatusCode,
			params.Latency,
			params.ClientIP,
			params.Keys["project"],
			params.Keys["department"],
			params.Keys["user"],
			params.Method,
			params.Path,
			params.ErrorMessage,
		)
	}

	return gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: formatter,
		Output:    logger{},
	})
}
