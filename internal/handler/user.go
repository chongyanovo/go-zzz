package handler

import (
	service "github.com/chongyanovo/go-zzz/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type UserHandler struct {
	userService service.UserService
	l           *zap.Logger
}

func NewUserHandler(userService service.UserService, l *zap.Logger) *UserHandler {
	return &UserHandler{
		userService: userService,
		l:           l,
	}
}

func (uh UserHandler) RegisterRoutes(server *gin.Engine) {
	ug := server.Group("/user")
	ug.GET("/hello", uh.Hello)
}

func (uh UserHandler) Hello(ctx *gin.Context) {
	ctx.String(http.StatusOK, "hello world")
}
