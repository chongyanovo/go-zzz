package logger

import (
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/atomic"
	"io"
	"time"
)

// MiddleWareBuilder 日志打印中间件
type MiddleWareBuilder struct {
	allowRequestBody  *atomic.Bool
	allowResponseBody *atomic.Bool
	loggerFunc        func(ctx context.Context, log *AccessLog)
}

// NewBuilder 创建日志打印中间件构造器
func NewBuilder(loggerFunc func(ctx context.Context, log *AccessLog)) *MiddleWareBuilder {
	return &MiddleWareBuilder{
		allowRequestBody:  atomic.NewBool(false),
		allowResponseBody: atomic.NewBool(false),
		loggerFunc:        loggerFunc,
	}
}

// AllowRequestBody 允许打印请求体
func (b *MiddleWareBuilder) AllowRequestBody() *MiddleWareBuilder {
	b.allowRequestBody.Store(true)
	return b
}

// AllowResponseBody 允许打印响应体
func (b *MiddleWareBuilder) AllowResponseBody() *MiddleWareBuilder {
	b.allowResponseBody.Store(true)
	return b
}

// Build 构建中间件
func (b *MiddleWareBuilder) Build() gin.HandlerFunc {
	start := time.Now()
	return func(ctx *gin.Context) {
		url := ctx.Request.URL.String()
		if len(url) > 1024 {
			url = url[:1024] + "..."
		}
		log := AccessLog{
			Method: ctx.Request.Method,
			Url:    url,
		}
		if b.allowRequestBody.Load() && ctx.Request.Body != nil {
			body, _ := ctx.GetRawData()
			ctx.Request.Body = io.NopCloser(bytes.NewReader(body))
			if len(body) > 1024 {
				body = body[:1024]
			}
			log.RequestBody = string(body)
		}

		if b.allowResponseBody.Load() {
			ctx.Writer = &responseWriter{
				ResponseWriter: ctx.Writer,
				log:            &log,
			}
		}
		log.Duration = time.Since(start).String()
		defer func() {
			b.loggerFunc(ctx, &log)
		}()
		ctx.Next()

	}
}

// responseWriter 装饰gin.ResponseWriter获取响应体
type responseWriter struct {
	log *AccessLog
	gin.ResponseWriter
}

// Write 获取响应体
func (w *responseWriter) Write(data []byte) (int, error) {
	w.log.ResponseBody = string(data)
	return w.ResponseWriter.Write(data)
}

// WriteString 获取响应体
func (w *responseWriter) WriteString(data string) (n int, err error) {
	w.log.ResponseBody = data
	return w.ResponseWriter.WriteString(data)
}

// WriteHeader 获取响应状态码
func (w *responseWriter) WriteHeader(statusCode int) {
	w.log.Status = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

// AccessLog 请求日志
type AccessLog struct {
	Duration     string
	Method       string
	Url          string
	Status       int
	RequestBody  string
	ResponseBody string
}
