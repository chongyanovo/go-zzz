package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type UserHandler struct {
	//userService UserService
	l *zap.Logger
}

//func NewUserHandler(userService UserService) *UserHandler {
//	return &UserHandler{
//		userService: userService,
//	}
//}

func NewUserHandler(l *zap.Logger) *UserHandler {
	return &UserHandler{
		l: l,
	}
}

func (uh UserHandler) Hello(ctx *gin.Context) {
	ctx.String(http.StatusOK, "hello world")
}
