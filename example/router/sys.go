package router

import (
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/example/api/sys"
)

func Sys(party iris.Party) {
	r := party.Party("/sys")
	//Sys

	//系统字典值
	r.Get("/dict", sys.Dict)

	//以下接口 不用写
	//	秘钥
	//r.Get("/key", sys.GetRSAPublicKey)
	// 系统版本
	//r.Get("/version", sys.GetVersion)
	// 错误码
	//r.Get("/code", )
}
