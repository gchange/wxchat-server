package wxchatserver

import (
	"wx-server/internal/ai"
	"wx-server/internal/httpserver"
	"wx-server/internal/weixin"
)

type Config struct {
	Server    httpserver.Config `json:"server" yaml:"server"`
	SecretKey []byte            `json:"secret_key" yaml:"secret_key"`
	WeiXin    weixin.Config     `json:"weixin" yaml:"weixin"`
	Qwen      ai.QwenConfig     `json:"qwen" yaml:"qwen"`
}
