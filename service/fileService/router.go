// Package fileService
// @Author:        asus
// @Description:   $
// @File:          router.go
// @Data:          2022/4/1116:51
//
package fileService

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/xgpc/dsg"
	"github.com/xgpc/dsg/env"
	"github.com/xgpc/dsg/util/guzzle"
)

func Router(app iris.Party) {

	serviceConf := env.Config.Microservices
	uploadClient = guzzle.NewClient(guzzle.WithHost(serviceConf.FileAddr))

	mvc.Configure(app, func(m *mvc.Application) {
		m.Register(func(ctx iris.Context) *dsg.Base {
			return dsg.NewBase(ctx)
		})

		m.Party("/server").Handle(new(UploadController)) // 文件上传
	})
}
