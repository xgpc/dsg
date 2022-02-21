package user

import (
	"context"
	"github.com/xgpc/dsg/service/grpcService/proto"
)

func UpdateMobile(mobile string, userID uint32) error {
	c := proto.NewUserServiceClient(proto.GRPCConn)
	//调用函数
	_, err := c.UpdateMobile(context.Background(), &proto.User{
		Mobile: mobile,
		Id:     userID,
	})
	return err

}
