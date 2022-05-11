package proto

import (
	"fmt"
	"github.com/xgpc/dsg/frame"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var GRPCConn *grpc.ClientConn

func GRPCConnect() {
	// 监听端口
	conn, err := grpc.Dial(frame.Config.Microservices.RPCAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("连接服务失败失败: %s", err)
		return
	}
	GRPCConn = conn
}
