package service

import (
	"github.com/xgpc/dsg/service/cryptService"
)

func CacheStart() {
	//微信公众号信息缓存
	//wechatService.SetWeChat()

	cryptService.SetRsaKey()
}
