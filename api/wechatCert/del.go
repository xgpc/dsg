package wechatCert

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/service/httpService"
)

func Del(ctx iris.Context) {
	del("")
}

func del(appID string) {
	param := make(map[string]string)
	param["app_id"] = appID
	marshal, err := json.Marshal(&param)
	if err != nil {

	}
	request, err := httpService.DoRequest("http://localhost:16100/api/wechat/cert/del", httpService.ReqPost, marshal)
	if err != nil {

	}
	fmt.Println(string(request))
	/**/
}
