package cryptExam

import (
	"encoding/json"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/frame"
)

// BackJson 返回Json
// 使用前端的公钥对后端AES加密
func BackJson(ctx iris.Context) {
	this := frame.NewBase(ctx)

	m := map[string]interface{}{
		"Code": 1001,
		"Name": "Tom",
		"Desc": "是一名学生",
	}
	marshal, _ := json.Marshal(m)

	this.CryptSend(marshal)

	return
}
