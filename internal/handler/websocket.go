package handler

import (
	ws "github.com/chongyanovo/go-zzz/pkg/ginx/websocket"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type WebSocketHandler struct {
	logger           *zap.Logger
	webSocketManager map[string]*ws.Manager
}

func NewWebSocketHandler(logger *zap.Logger, webSocketManager map[string]*ws.Manager) *WebSocketHandler {
	return &WebSocketHandler{
		logger:           logger,
		webSocketManager: webSocketManager,
	}
}

func (wh *WebSocketHandler) RegisterRoutes(server *gin.Engine) {
	wsr := server.Group("/ws")
	wsr.GET("/monitor", wh.MonitorWebSocketHandler)
}

func (wh *WebSocketHandler) MonitorWebSocketHandler(ctx *gin.Context) {
	manager := wh.webSocketManager["monitor"]
	err := ws.DefaultServeWs(manager, ctx.Writer, ctx.Request)
	if err != nil {
		return
	}
	for range time.NewTicker(time.Second * 5).C {
		manager.Broadcast <- []byte("监控数据")
	}
}
