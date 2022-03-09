package wechat

import (
	"context"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/service/grpcService/proto"
)

// BindUser 绑定用户
func BindUser(mobile string, openID string, sysCode uint32) uint32 {
	c := proto.NewWechatServiceClient(proto.GRPCConn)
	//调用函数
	r, err := c.BindUser(context.Background(), &proto.WechatMobile{
		Mobile:  mobile,
		OpenID:  openID,
		SysCode: sysCode,
	})
	if err != nil {
		exce.ParseErr(err)
	}
	return r.GetId()
}
