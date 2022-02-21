package user

import (
	"context"
	"github.com/xgpc/dsg/service/grpcService/proto"
)

func ExpireToken(token string) error {
	c := proto.NewUserServiceClient(proto.GRPCConn)
	//调用函数
	_, err := c.ExpireToken(context.Background(), &proto.Token{
		Token: token,
	})
	return err
}
