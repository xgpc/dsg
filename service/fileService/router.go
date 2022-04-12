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
	"github.com/xgpc/dsg/frame"
	"github.com/xgpc/dsg/util/guzzle"
)

func Router(app iris.Party) {

	serviceConf := frame.Config.Microservices
	uploadClient = guzzle.NewClient(guzzle.WithHost(serviceConf.FileAddr))

	mvc.Configure(app, func(m *mvc.Application) {
		m.Register(func(ctx iris.Context) *frame.Base {
			return frame.NewBase(ctx)
		})

		m.Party("/server").Handle(new(UploadController)) // 文件上传
	})
}
