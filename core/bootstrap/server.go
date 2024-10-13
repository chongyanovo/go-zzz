package bootstrap

import (
	"github.com/chongyanovo/go-zzz/internal/handler"
	"github.com/gin-gonic/gin"
)

// ServerConfig server配置
type ServerConfig struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	Port int    `mapstructure:"port" json:"port" yaml:"port"`
}

func NewServer(middlewares []gin.HandlerFunc, uh *handler.UserHandler, wh *handler.WebSocketHandler) *gin.Engine {
	server := gin.Default()
	server.Use(middlewares...)
	uh.RegisterRoutes(server)
	wh.RegisterRoutes(server)
	return server
}
