package user

import (
	"context"
	"fmt"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/service/grpcService/proto"
)

// Get 用户ID和手机号
func Get(id uint32) (uint32, string) {
	c := proto.NewUserServiceClient(proto.GRPCConn)
	r, err := c.GetByID(context.Background(), &proto.UserID{
		Id: id,
	})

	if err != nil {
		exce.ParseErr(err)
	}
	return r.GetId(), r.GetMobile()
}

// GetInfoByToken 通过token获取userID和openID
func GetInfoByToken(token string) (userID uint32, openID string) {
	c := proto.NewUserServiceClient(proto.GRPCConn)
	//调用函数
	reply, err := c.GetInfoByToken(context.Background(), &proto.Token{
		Token: token,
	})
	if err != nil {
		exce.ParseErr(err)
	}
	return reply.GetUserID(), reply.GetOpenID()
}

// GetIDByMobile 通过手机号查询
func GetIDByMobile(mobile string) uint32 {
	c := proto.NewUserServiceClient(proto.GRPCConn)
	//调用函数
	reply, err := c.GetIDByMobile(context.Background(), &proto.Mobile{
		Mobile: mobile,
	})
	if err != nil {
		exce.ParseErr(err)
	}
	return reply.GetId()
}

// ListByIDs 多用户ID查询
func ListByIDs(id []uint32) []proto.User {
	c := proto.NewUserServiceClient(proto.GRPCConn)

	stream, err := c.ListByIDs(context.Background())
	if err != nil {
		exce.ParseErr(err)
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
