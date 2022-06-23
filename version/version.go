package version

import (
	"io/ioutil"

	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/frame"
)

// 版本管理
// 需要配合 git describe --tags --always >> app.version 使用
func Version(ctx iris.Context) {
	this := frame.NewBase(ctx)
	md := map[string]interface{}{}

	f, err := ioutil.ReadFile("app.version")
	if err != nil {
		md["version"] = "nil"
	} else {
		md["version"] = string(f)
	}
	this.SuccessWithData(md)
}
