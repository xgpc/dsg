package wechat

import (
	"context"
	"github.com/xgpc/dsg/service/grpcService/proto"
)

// LoginSub 公众号登录
func LoginSub(appID, code string, sysCode uint32) (string, error) {
	c := proto.NewWechatServiceClient(proto.GRPCConn)
	//调用函数
	reply, err := c.Login(context.Background(), &proto.WechatLogin{
		AppID:     appID,
		Code:      code,
		SysCode:   sysCode,
		LoginType: 4,
	})
	return reply.GetToken(), err
}

// LoginMini 小程序登录
func LoginMini(appID, code string, sysCode uint32) (string, error) {
	c := proto.NewWechatServiceClient(proto.GRPCConn)
	//调用函数
	reply, err := c.Login(context.Background(), &proto.WechatLogin{
		AppID:     appID,
		Code:      code,
		SysCode:   sysCode,
		LoginType: 3,
	})
	return reply.GetToken(), err
}
