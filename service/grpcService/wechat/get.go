package wechat

import (
	"context"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/service/grpcService/proto"
)

func GetAccessToken(appID string) string {
	c := proto.NewWechatServiceClient(proto.GRPCConn)
	//调用函数
	reply, err := c.GetAccessToken(context.Background(), &proto.AppID{
		AppID: appID,
	})
	if err != nil {
		exce.ParseErr(err)
	}
	return reply.GetToken()
}
