package core

import (
	"github.com/chongyanovo/go-zzz/core/bootstrap"
	"github.com/chongyanovo/go-zzz/pkg/ginx/websocket"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Application struct {
	Config           *bootstrap.Config
	Viper            *viper.Viper
	DB               *gorm.DB
	Redis            redis.Cmdable
	Logger           *zap.Logger
	Server           *gin.Engine
	WebSocketManager map[string]*websocket.Manager
}

// NewApplication 初始化 Application
func NewApplication(config *bootstrap.Config,
	viper *viper.Viper,
	db *gorm.DB,
	redis redis.Cmdable,
	logger *zap.Logger,
	server *gin.Engine,
	wsm map[string]*websocket.Manager) *Application {
	return &Application{
		Config:           config,
		Viper:            viper,
		DB:               db,
		Redis:            redis,
		Logger:           logger,
		Server:           server,
		WebSocketManager: wsm,
	}
}
