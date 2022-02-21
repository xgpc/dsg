package router

import (
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/api/sys"
)

func Sys(party iris.Party) {
	r := party.Party("/sys")
	//Sys
	//r.Get("/dict", sys.Get)
	//	秘钥
	r.Get("/rsa", sys.GetRSAPublicKey)
	// 系统版本
	r.Get("/version", sys.GetVersion)
	//r.Get("/byte", sys.Byte)
	//r.Get("/hash", sys.Hash)
	//r.Get("/str", sys.Str)
	//r.Post("/receive", sys.Receive)
}
