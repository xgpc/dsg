package user

import (
	"fmt"
	"github.com/xgpc/dsg"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/frame"
	"github.com/xgpc/dsg/service/grpcService/wechat"

	"github.com/kataras/iris/v12"
)

func Login(ctx iris.Context) {
	var param struct {
		Code string
	}
	this := frame.NewBase(ctx)
	this.Init(&param)
	appID := dsg.Config.Wechat.AppID
	fmt.Println("appID:::", appID)
	// 小程序登录
	token, err := wechat.LoginMini(appID, param.Code, 1)
	if err != nil {
		exce.ThrowSys(exce.CodeSysError, err.Error())
	}
	this.SuccessWithData(token)
}
