//go:build wireinject
// +build wireinject

package main

import (
	"github.com/chongyanovo/go-zzz/core"
	"github.com/chongyanovo/go-zzz/core/bootstrap"
	handler "github.com/chongyanovo/go-zzz/internal/handler"
	repository "github.com/chongyanovo/go-zzz/internal/repository"
	dao "github.com/chongyanovo/go-zzz/internal/repository/dao"
	service "github.com/chongyanovo/go-zzz/internal/service"
	"github.com/google/wire"
)

var BootstrapProvider = wire.NewSet(
	bootstrap.NewViper,
	bootstrap.NewConfig,
	bootstrap.NewMysql,
	bootstrap.NewRedis,
	bootstrap.NewZap,
	bootstrap.NewMiddlewares,
	bootstrap.NewServer,
	bootstrap.NewWebSocketManager,
	handler.NewWebSocketHandler,
	core.NewApplication,
)

var UserProvider = wire.NewSet(
	handler.NewUserHandler,
	service.NewUserService,
	repository.NewUserRepository,
	dao.NewUserDao,
)

func InitApp() (*core.Application, error) {
	wire.Build(
		UserProvider,
		BootstrapProvider,
	)
	return &core.Application{}, nil
}
