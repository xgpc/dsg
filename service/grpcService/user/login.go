package user

import (
	"context"
	"github.com/xgpc/dsg/service/grpcService/proto"
)

func Login(loginType uint32, device, mobile, code string) (string, error) {
	c := proto.NewUserServiceClient(proto.GRPCConn)
	//调用函数
	reply, err := c.Login(context.Background(), &proto.Login{
		//登录方式，只能web和app,见models/constant.go
		LoginType: loginType,
		//设备信息
		Device: device,
		Mobile: mobile,
		//验证码
		Code: code,
	})
	return reply.GetToken(), err

}
