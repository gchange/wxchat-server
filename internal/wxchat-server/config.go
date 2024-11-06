package wxchatserver

import (
	"wx-server/internal/httpserver"
	"wx-server/internal/weixin"
)

type AIConfig struct {
	BaseUrl string `json:"base_url" yaml:"base_url"`
	Model   string `json:"model" yaml:"model"`
	Key     string `json:"key" yaml:"key"`
}

type Config struct {
	Server    httpserver.Config   `json:"server" yaml:"server"`
	SecretKey []byte              `json:"secret_key" yaml:"secret_key"`
	WeiXin    weixin.Config       `json:"weixin" yaml:"weixin"`
	AI        map[string]AIConfig `json:"ai" yaml:"ai"`
}
