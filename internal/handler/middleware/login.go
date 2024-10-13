package middleware

import (
	jwt2 "github.com/ChongYanOvO/little-blue-book/pkg/ginx/jwt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

// LoginBuilder JWT 登录校验
type LoginBuilder struct {
	paths []string
}

func NewLoginBuilder() *LoginBuilder {
	return &LoginBuilder{}
}

func (l *LoginBuilder) IgnorePaths(path string) *LoginBuilder {
	l.paths = append(l.paths, path)
	return l
}

func (l *LoginBuilder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		for _, path := range l.paths {
			if ctx.Request.URL.Path == path {
				ctx.Next()
				return
			}
		}

		_, err := jwt2.ExtractToken(ctx)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		uc, err := jwt2.ExtractJwtClaims(ctx)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		now := time.Now()
		if uc.ExpiresAt.Sub(now) < time.Minute {
			uc.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute))
			if err := jwt2.SetJwtToken(ctx, uc.Uid, uc.Email); err != nil {
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		}
	}
}
