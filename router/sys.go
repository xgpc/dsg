package router

import (
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/api/sys"
)

func Sys(party iris.Party) {
	r := party.Party("/sys")
	// Sys 系统字典 由各系统返回

	//	秘钥
	r.Get("/key", sys.GetRSAPublicKey)
	// 系统版本
	r.Get("/version", sys.GetVersion)
	// 错误码
	r.Get("/code", sys.GetCode)
}
