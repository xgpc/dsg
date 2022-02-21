package msgCode

import (
	"context"
	"github.com/xgpc/dsg/service/grpcService/proto"
)

// SendMsgCode 发送验证码
func SendMsgCode(mobile string) (msg, code string, err error) {
	c := proto.NewMsgServiceClient(proto.GRPCConn)
	//调用函数
	reply, err := c.SendMsgCode(context.Background(), &proto.Mobile{
		Mobile: mobile,
	})
	if err != nil {
		return "", "", err
	}
	return reply.GetMsg(), reply.GetCode(), err
}
