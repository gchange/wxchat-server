package wxchatserver

import (
	"wx-server/internal/httpserver"
	"wx-server/internal/weixin"
)

type AIConfig struct {
	BaseUrl string `json:"url" yaml:"url" mapstructure:"base-url"`
	Model   string `json:"model" yaml:"model" mapstructure:"model"`
	Key     string `json:"key" yaml:"key" mapstructure:"key"`
}

type Config struct {
	Server    httpserver.Config   `json:"server" yaml:"server" mapstructure:"server"`
	SecretKey []byte              `json:"secret-key" yaml:"secret-key" mapstructure:"secret-key"`
	WeiXin    weixin.Config       `json:"weixin" yaml:"weixin" mapstructure:"weixin"`
	AI        map[string]AIConfig `json:"ai" yaml:"ai" mapstructure:"ai"`
}
