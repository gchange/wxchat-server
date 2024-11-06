package wxchatserver

import (
	"fmt"
	"sync"
	"wx-server/internal/httpserver"
	"wx-server/internal/httpserver/middleware"
	"wx-server/internal/weixin"
)

type Server struct {
	config Config
	server *httpserver.HttpServer
	wx     *weixin.Client
	ai     sync.Map
}

func NewServer(config *Config) (server *Server, err error) {
	s := httpserver.NewServer(&config.Server)

	wx, err := weixin.NewClient(&config.WeiXin)
	if err != nil {
		return nil, err
	}
	server = &Server{
		config: *config,
		server: s,
		wx:     wx,
		ai:     sync.Map{},
	}

	for name, c := range config.AI {
		fmt.Printf("%s:%+v\n", name, c)
		ai, err := NewAI(&c)
		if err != nil {
			return nil, err
		}
		server.ai.Store(name, ai)
	}

	rg := s.RootRouterGroup()
	rg.Use(middleware.WrapLogger()).Use(middleware.WrapTimer())
	if config.Server.PProf {
		httpserver.WrapPProf(rg)
	}

	unauth := rg.Group("/")
	unauth.POST("login", server.login)

	auth := rg.Group("/")
	//auth.GET("wschat", server.wschat)
	auth.POST("chat", server.chat)
	return
}

func (s *Server) Run() error {
	return s.server.ListenAndServe()
}
