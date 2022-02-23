package cryptExam

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/frame"
)

// ReceiveData 接收前端加密的数据
func ReceiveData(ctx iris.Context) {
	this := frame.NewBase(ctx)

	bytes := this.CryptReceive()

	fmt.Println(string(bytes))
	return
}
