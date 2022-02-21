package wechat

import (
	"context"
	"github.com/xgpc/dsg/service/grpcService/proto"
)

func GetAccessToken(appID string) (string, error) {
	c := proto.NewWechatServiceClient(proto.GRPCConn)
	//调用函数
	reply, err := c.GetAccessToken(context.Background(), &proto.AppID{
		AppID: appID,
	})
	return reply.GetToken(), err
}
