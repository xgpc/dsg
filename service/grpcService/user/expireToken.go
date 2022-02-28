package user

import (
	"context"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/service/grpcService/proto"
)

func ExpireToken(token string) {
	c := proto.NewUserServiceClient(proto.GRPCConn)
	//调用函数
	_, err := c.ExpireToken(context.Background(), &proto.Token{
		Token: token,
	})
	if err != nil {
		exce.ParseErr(err)
	}
}
