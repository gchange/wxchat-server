package httpserver

import (
	"wx-server/internal/logging"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	host   string
	port   string
	engine *gin.Engine
}

func NewServer(config *Config) *HttpServer {
	engine := gin.New()
	engine.Use(gin.Recovery())
	return &HttpServer{
		host:   config.Host,
		port:   config.Port,
		engine: engine,
	}
}

func (s *HttpServer) RootRouterGroup() *gin.RouterGroup {
	return &s.engine.RouterGroup
}

func (s *HttpServer) ListenAndServe() error {
	logging.Info("starting http server on " + s.host + ":" + s.port)
	return s.engine.Run(s.host + ":" + s.port)
}
