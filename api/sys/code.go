package sys

import (
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
)

// GetCode 错误码
func GetCode(ctx iris.Context) {
	this := frame.NewBase(ctx)
	this.SuccessWithData(exce.ErrString)
	return
}
