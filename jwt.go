package dsg

import (
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/v2/exce"
	"github.com/xgpc/dsg/v2/pkg/jwt"
	"time"
)

// jwt
var _jetKey string

func OptionJwt(JwtKey string) func() error {
	return func() error {
		_jetKey = JwtKey
		return nil
	}
}

// Login 中间件 login
func Login(ctx iris.Context) {
	MiddlewareJwt(ctx)
}

// MiddlewareJwt 中间件 login
func MiddlewareJwt(ctx iris.Context) {
	p := NewBase(ctx)

	token := ctx.GetHeader("token")
	parseToken, err := jwt.ParseToken(_jetKey, token)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, "解析token失败"+err.Error())
	}
	p.SetMyId(parseToken.UserID)
}

// CreateToken 创建Token
func CreateToken(UserID uint32, ExpiresAt time.Duration) string {
	ExpiresNum := time.Now().Add(ExpiresAt).Unix()
	token, err := jwt.MakeToken(_jetKey, UserID, ExpiresNum)
	if err != nil {
		exce.ThrowSys(exce.CodeSysBusy, "创建token失败,"+err.Error())
	}
	return token
}

// ParseToken 解析Token
func ParseToken(token string) *jwt.MapClaims {
	parseToken, err := jwt.ParseToken(_jetKey, token)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, "解析token失败"+err.Error())
	}
	return parseToken
}
