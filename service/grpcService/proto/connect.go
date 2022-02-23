package proto

import (
	"fmt"
	"github.com/xgpc/dsg/frame"
	"google.golang.org/grpc"
)

var GRPCConn *grpc.ClientConn

func GRPCConnect() {
	// 监听端口
	conn, err := grpc.Dial(frame.Config.App.RPCAddr, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("连接服务失败失败: %s", err)
		return
	}
	GRPCConn = conn
}
