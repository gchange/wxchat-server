package wxchatserver

import (
	"wx-server/internal/ai"
	"wx-server/internal/httpserver"
	"wx-server/internal/httpserver/middleware"
	"wx-server/internal/weixin"
)

type Server struct {
	config Config
	server *httpserver.HttpServer
	wx     *weixin.Client
	qwen   ai.AI
}

func NewServer(config *Config) (server *Server, err error) {
	s := httpserver.NewServer(&config.Server)

	wx, err := weixin.NewClient(&config.WeiXin)
	if err != nil {
		return nil, err
	}
	qwen, err := ai.NewQwen(&config.Qwen)
	if err != nil {
		return nil, err
	}
	server = &Server{
		config: *config,
		server: s,
		wx:     wx,
		qwen:   qwen,
	}

	rg := s.RootRouterGroup()
	rg.Use(middleware.WrapLogger()).Use(middleware.WrapTimer())
	if config.Server.PProf {
		httpserver.WrapPProf(rg)
	}

	unauth := rg.Group("/")
	unauth.POST("login", server.login)

	auth := rg.Group("/")
	auth.GET("wschat", server.wschat)
	auth.POST("chat", server.chat)
	return
}

func (s *Server) Run() error {
	return s.server.ListenAndServe()
}
