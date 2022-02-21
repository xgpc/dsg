package cryptExam

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/frame"
	"github.com/xgpc/dsg/service/grpcService/user"
)

// BackMobile 返回手机号
// 使用前端的公钥对后端AES加密
func BackMobile(ctx iris.Context) {
	var param struct {
		UserID uint32 `validate:"required"`
	}
	this := frame.NewBase(ctx)
	this.Init(&param)

	// 01 gRPC查询手机号
	userID, mobile, err := user.Get(param.UserID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userID)
	fmt.Println(mobile)

	this.CryptSend([]byte(mobile))

	return
}
