package core

import (
	"github.com/chongyanovo/go-zzz/core/bootstrap"
	"github.com/chongyanovo/go-zzz/core/logger"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Application struct {
	Config *bootstrap.Config
	Viper  *viper.Viper
	DB     *gorm.DB
	Redis  redis.Cmdable
	Logger logger.Logger
	Server *gin.Engine
}

// NewApplication 初始化 Application
func NewApplication(config *bootstrap.Config,
	viper *viper.Viper,
	db *gorm.DB,
	redis redis.Cmdable,
	logger logger.Logger,
	server *gin.Engine) Application {
	return Application{
		Config: config,
		Viper:  viper,
		DB:     db,
		Redis:  redis,
		Logger: logger,
		Server: server,
	}
}
