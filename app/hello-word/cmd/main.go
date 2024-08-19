package main

import (
	"context"
	"fmt"
	"github.com/chongyanovo/go-zzz/pkg/ginx/middleware/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	app, err := InitApp()
	if err != nil {
		panic(err)
	}

	engine := app.Server
	engine.Use(logger.
		NewBuilder(func(ctx context.Context, log *logger.AccessLog) {
			app.Logger.Info("Http请求", zap.Any("日志", log))
		}).
		AllowRequestBody().
		AllowResponseBody().
		Build())

	ug := engine.Group("/user")
	ug.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})
	engine.Run(fmt.Sprintf("%s:%d",
		app.Config.ServerConfig.Host,
		app.Config.ServerConfig.Port))

}
