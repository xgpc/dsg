package jwt

// package jwt
// @Author: dsg
// @Description:
// @File:  jwt.go
// @Date: 2023/5/16 19:22

import (
	"errors"
	goJwt "github.com/golang-jwt/jwt/v5"

	"github.com/xgpc/dsg/v2/exce"
	"time"
)

type MapClaims struct {
	UserID uint32 `json:"user_id"`
	goJwt.RegisteredClaims
}

func (m MapClaims) Validate() error {
	if m.UserID == 0 {
		return exce.CodeUserNoLogin
	}

	ExpiresTime, err := m.GetExpirationTime()
	if err != nil {
		return errors.New("jwt-GetExpirationTime error:, " + err.Error())
	}
	if ExpiresTime.Unix() <= time.Now().Unix() {
		return errors.New("token 已过期")
	}

	return nil
}

// MakeToken 创建token
func MakeToken(key []byte, userID uint32, ExpiresNum int64) (string, error) {
	claims := MapClaims{
		UserID: userID,
		RegisteredClaims: goJwt.RegisteredClaims{
			//iss: 签发者
			//sub: 面向的用户
			//aud: 接收方
			//exp: 过期时间
			ExpiresAt: goJwt.NewNumericDate(time.Unix(ExpiresNum, 0)),
			//nbf: 生效时间
			//iat: 签发时间
			//jti: 唯一身份标识
		},
	}
	token := goJwt.NewWithClaims(goJwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}

// ParseToken 解析token
func ParseToken(key []byte, tokenStr string) (*MapClaims, error) {

	// 声明一个空的数据声明
	iJwtCustomClaims := MapClaims{}
	//ParseWithClaims是NewParser().ParseWithClaims()的快捷方式
	//第一个值是token ，
	//第二个值是我们之后需要把解析的数据放入的地方，
	//第三个值是Keyfunc将被Parse方法用作回调函数，以提供用于验证的键。函数接收已解析但未验证的令牌。
	_, err := goJwt.ParseWithClaims(tokenStr, &iJwtCustomClaims, func(token *goJwt.Token) (interface{}, error) {
		return key, nil
	})

	// 判断 是否为空 或者是否无效只要两边有一处是错误 就返回无效token
	if err != nil {
		return nil, err
	}

	//if iJwtCustomClaims.ExpiresAt.Unix() <= time.Now().Unix() {
	//    return nil, errors.New("token已过期")
	//}

	return &iJwtCustomClaims, err
}
