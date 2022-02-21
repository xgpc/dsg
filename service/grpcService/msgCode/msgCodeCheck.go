package msgCode

import (
	"context"
	"github.com/xgpc/dsg/service/grpcService/proto"
)

// CheckMsgCode 检查验证码正确性
func CheckMsgCode(mobile, code string) (bool, error) {
	c := proto.NewMsgServiceClient(proto.GRPCConn)
	//调用函数
	reply, err := c.CheckMsgCode(context.Background(), &proto.MsgCode{
		Mobile: mobile,
		Code:   code,
	})
	return reply.GetF(), err
}
