package middleware

import (
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/v2/exce"
)

func ExceptionLog(ctx iris.Context) {
	defer func() {
		if err := recover(); err != nil {
			if ctx.IsStopped() {
				return
			}
			exce.DealException(ctx, err)
			ctx.StopExecution()
		}
	}()
	ctx.Next()
}
