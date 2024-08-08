package bootstrap

import (
	"github.com/chongyanovo/go-zzz/bootstrap/internal"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Application struct {
	Config *internal.Config
	Viper  *viper.Viper
	DB     *gorm.DB
	Redis  redis.Cmdable
	Logger *zap.Logger
	Server *gin.Engine
}

// NewApplication 初始化 Application
func NewApplication(config *internal.Config,
	viper *viper.Viper,
	db *gorm.DB,
	redis redis.Cmdable,
	logger *zap.Logger,
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
