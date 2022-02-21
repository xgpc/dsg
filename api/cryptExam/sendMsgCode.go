package cryptExam

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/frame"
	"github.com/xgpc/dsg/service/grpcService/msgCode"
)

// SendCode 发送验证码
// 前端传手机号到后端,前端需要先获取后端的RSA公钥
func SendCode(ctx iris.Context) {
	this := frame.NewBase(ctx)

	bytes := this.CryptReceive()

	var mobile = string(bytes)

	fmt.Println("mobile:::", mobile)
	return
	msg, code, err := msgCode.SendMsgCode(mobile)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("msg:", msg)
	fmt.Println("code:", code)
	return
}
