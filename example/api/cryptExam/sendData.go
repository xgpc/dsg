package cryptExam

import (
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/frame"
)

// SendData 发送加密数据给前端
func SendData(ctx iris.Context) {
	var param struct {
		UserID uint32 `validate:"required"`
	}
	this := frame.NewBase(ctx)
	this.Init(&param)

	this.CryptSend([]byte("data"))

	return
}
