package handler

import (
	"github.com/chongyanovo/go-zzz/core/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	//userService UserService
	l logger.Logger
}

//func NewUserHandler(userService UserService) *UserHandler {
//	return &UserHandler{
//		userService: userService,
//	}
//}

func NewUserHandler(l logger.Logger) *UserHandler {
	return &UserHandler{
		l: l,
	}
}

func (uh UserHandler) Hello(ctx *gin.Context) {
	ctx.String(http.StatusOK, "hello world")
}
