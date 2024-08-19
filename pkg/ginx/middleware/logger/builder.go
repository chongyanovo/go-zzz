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
	urlMaxLength      *atomic.Int32
	bodyMaxLength     *atomic.Int32
	loggerFunc        func(ctx context.Context, log *AccessLog)
}

// NewBuilder 创建日志打印中间件构造器
func NewBuilder(loggerFunc func(ctx context.Context, log *AccessLog)) *MiddleWareBuilder {
	return &MiddleWareBuilder{
		allowRequestBody:  atomic.NewBool(false),
		allowResponseBody: atomic.NewBool(false),
		urlMaxLength:      atomic.NewInt32(1024),
		bodyMaxLength:     atomic.NewInt32(1024),
		loggerFunc:        loggerFunc,
	}
}

// AllowRequestBody 允许打印请求体
func (b *MiddleWareBuilder) AllowRequestBody(flag bool) *MiddleWareBuilder {
	b.allowRequestBody.Store(flag)
	return b
}

// AllowResponseBody 允许打印响应体
func (b *MiddleWareBuilder) AllowResponseBody(flag bool) *MiddleWareBuilder {
	b.allowResponseBody.Store(flag)
	return b
}

// UrlLength 设置url最大长度
func (b *MiddleWareBuilder) UrlLength(length int) *MiddleWareBuilder {
	b.urlMaxLength.Store(int32(length))
	return b
}

// BodyLength 设置请求体最大长度
func (b *MiddleWareBuilder) BodyLength(length int) *MiddleWareBuilder {
	b.bodyMaxLength.Store(int32(length))
	return b
}

// Build 构建中间件
func (b *MiddleWareBuilder) Build() gin.HandlerFunc {
	start := time.Now()
	return func(ctx *gin.Context) {
		url := ctx.Request.URL.String()
		urlMaxLength := int(b.urlMaxLength.Load())
		if len(url) > urlMaxLength {
			url = url[:urlMaxLength] + "..."
		}
		log := AccessLog{
			Method: ctx.Request.Method,
			Url:    url,
		}
		if b.allowRequestBody.Load() && ctx.Request.Body != nil {
			body, _ := ctx.GetRawData()
			ctx.Request.Body = io.NopCloser(bytes.NewReader(body))
			bodyMaxLength := int(b.bodyMaxLength.Load())
			if len(body) > bodyMaxLength {
				body = body[:bodyMaxLength]
			}
			log.RequestBody = string(body)
		}

		if b.allowResponseBody.Load() {
			ctx.Writer = &responseWriter{
				ResponseWriter: ctx.Writer,
				log:            &log,
				bodyMaxLength:  atomic.NewInt32(1024),
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
	bodyMaxLength *atomic.Int32
}

// Write 获取响应体
func (w *responseWriter) Write(data []byte) (int, error) {
	bodyMaxLength := int(w.bodyMaxLength.Load())
	if len(data) > bodyMaxLength {
		data = data[:bodyMaxLength]
	}
	w.log.ResponseBody = string(data)
	return w.ResponseWriter.Write(data)
}

// WriteString 获取响应体
func (w *responseWriter) WriteString(data string) (n int, err error) {
	bodyMaxLength := int(w.bodyMaxLength.Load())
	if len(data) > bodyMaxLength {
		data = data[:bodyMaxLength]
	}
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
