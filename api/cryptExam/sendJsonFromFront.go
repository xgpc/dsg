package cryptExam

import (
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/frame"
)

// SendJson 发送验证码
// 前端传手机号到后端,前端需要先获取后端的RSA公钥
func SendJson(ctx iris.Context) {
	this := frame.NewBase(ctx)
	bytes := this.CryptReceive()

	this.SuccessWithData(bytes)
	/*
		{
		  "Tag": "RmFwYWVjciFKRCE1M3I1cw==",
		  "Key": "ll077vq6ssT9DCTvWM8o1d+3KBxtay/XS63+0sUe1ZhBakwU06gcDMA9FCohPb1X4dx3dYeETNeHJHf2lJhS27RBIUOou7MITLVSTyd8+7hG9xrGtqkaRkG0RxK0oNwXFK6d29hpAYdoy+D16ShTziCMJ6JWNHdjLbqN5xBwnKQ=",
		  "Data": "uZ563/XQcQ9SxrVkC43TLOBkSwxxbsP2nlK+og=="
		}
	*/
	return
}

//手机号
