package router

import (
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/api/wechatCert"
)

func Wechat(party iris.Party) {
	r := party.Party("/wechat/cert")
	r.Post("/add", wechatCert.Add)
	r.Post("/del", wechatCert.Del)
}
