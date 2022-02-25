package sys

import (
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/frame"
	"github.com/xgpc/dsg/service/cryptService"
)

// GetRSAPublicKey 前端获取公钥
func GetRSAPublicKey(ctx iris.Context) {
	this := frame.NewBase(ctx)
	info := cryptService.RSAKey.Public
	//base64
	this.SuccessWithData(info)
	return
}
