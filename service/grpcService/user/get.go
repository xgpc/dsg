package user

import (
	"context"
	"fmt"
	"github.com/xgpc/dsg/service/grpcService/proto"
)

// Get 用户ID和手机号
func Get(id uint32) (uint32, string, error) {
	c := proto.NewUserServiceClient(proto.GRPCConn)
	r, err := c.GetByID(context.Background(), &proto.UserID{
		Id: id,
	})

	return r.GetId(), r.GetMobile(), err
}

// GetInfoByToken 通过token获取userID和openID
func GetInfoByToken(token string) (userID uint32, openID string, err error) {
	c := proto.NewUserServiceClient(proto.GRPCConn)
	//调用函数
	reply, err := c.GetInfoByToken(context.Background(), &proto.Token{
		Token: token,
	})
	return reply.GetUserID(), reply.GetOpenID(), err
}

// GetIDByMobile 通过手机号查询
func GetIDByMobile(mobile string) (uint32, error) {
	c := proto.NewUserServiceClient(proto.GRPCConn)
	//调用函数
	reply, err := c.GetIDByMobile(context.Background(), &proto.Mobile{
		Mobile: mobile,
	})
	return reply.GetId(), err

}

// ListByIDs 多用户ID查询
func ListByIDs(id []uint32) []proto.User {
	c := proto.NewUserServiceClient(proto.GRPCConn)

	stream, err := c.ListByIDs(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	// 传输参数
	for _, i := range id {
		err := stream.Send(&proto.UserID{Id: i})
		if err != nil {
			fmt.Println(err)
		}
	}
	// 关闭
	err = stream.CloseSend()
	if err != nil {
		fmt.Println(err)
	}
	// 接收响应
	var list []proto.User
	for {
		recv, err := stream.Recv()
		if err != nil {
			//fmt.Println(err)
			break
		}
		list = append(list, proto.User{Id: recv.GetU().GetId(), Mobile: recv.GetU().GetMobile()})
	}
	return list
}
