package bootstrap

import (
	"github.com/chongyanovo/go-zzz/pkg/ginx/websocket"
)

// WebSocketConfig websocket配置
type WebSocketConfig struct {
	Biz []string `mapstructure:"biz" json:"biz" yaml:"biz"`
}

func NewWebSocketManager(config *Config) map[string]*websocket.Manager {
	manager := make(map[string]*websocket.Manager)
	for _, biz := range config.WebSocketConfig.Biz {
		manager[biz] = websocket.NewWsManager()
	}
	return manager
}
