package wechatCert

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/service/httpService"
)

func Add(ctx iris.Context) {
	add("", "", "")
}

func add(appID, appSecret, appName string) {
	param := make(map[string]string)
	param["app_id"] = appID
	param["app_secret"] = appSecret
	param["app_name"] = appName
	marshal, err := json.Marshal(&param)
	if err != nil {

	}
	request, err := httpService.DoRequest("http://localhost:16100/api/wechat/cert/add", httpService.ReqPost, marshal)
	if err != nil {

	}
	fmt.Println(string(request))
}
