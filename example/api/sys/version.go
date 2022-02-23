package sys

import (
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/frame"
	"github.com/xgpc/dsg/service/sysService"
)

// GetVersion 获取后端最后一次启动时间
func GetVersion(ctx iris.Context) {
	this := frame.NewBase(ctx)
	version := sysService.GetSetSysVersion()
	this.SuccessWithData(version)
	return
}
