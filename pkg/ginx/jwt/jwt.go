package jwt

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

var (
	ErrTokenNotExist = errors.New("token不存在")
	ErrTokenExpired  = errors.New("token过期")
	ErrTokenInvalid  = errors.New("token无效")
	AccessKey        = []byte("oIft1b5qZjyLcc0zZo2UrUx5rk3KE0LvZKv73fw502oXd6vfYu1OAQvbSel8whvn")
	AccessHeader     = "Authorization"
)

type Claims interface {
}

type UserClaims struct {
	jwt.RegisteredClaims
	Uid   int64
	Email string
}

// SetJwtToken 设置Token
func SetJwtToken(ctx *gin.Context, id int64, email string) error {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512,
		&UserClaims{
			Uid:   id,
			Email: email,
			RegisteredClaims: jwt.RegisteredClaims{
				//ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
			},
		})
	tokenStr, err := token.SignedString(AccessKey)
	if err != nil {
		return err
	}
	ctx.Header(AccessHeader, "Bearer "+tokenStr)
	return nil
}

// ExtractJwtClaims 从前端请求中，提取tokenClaims
func ExtractJwtClaims(ctx *gin.Context) (*UserClaims, error) {
	tokenStr, err := ExtractToken(ctx)
	if err != nil {
		return nil, err
	}
	uc := &UserClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, uc, func(token *jwt.Token) (any, error) {
		return AccessKey, nil
	})
	if err != nil || token == nil || !token.Valid || uc.Uid == 0 {
		return nil, err
	}
	return uc, nil
}

// ExtractToken 从前端请求中，提取tokenStr
func ExtractToken(ctx *gin.Context) (string, error) {
	tokenStr := ctx.Request.Header.Get(AccessHeader)
	if tokenStr == "" {
		return "", ErrTokenNotExist
	}
	if len(strings.Split(tokenStr, " ")) != 2 {
		return "", ErrTokenInvalid
	}
	return strings.Split(tokenStr, " ")[1], nil
}
