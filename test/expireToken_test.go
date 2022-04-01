// Package user
// @Author:        asus
// @Description:   $
// @File:          expireToken_test.go
// @Data:          2022/4/116:59
//
package test

import (
	"fmt"
	"github.com/xgpc/dsg/service/grpcService/proto"
	"github.com/xgpc/dsg/service/grpcService/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
)

func TestExpireTokenByID(t *testing.T) {

	// 监听端口
	conn, err := grpc.Dial("127.0.0.1:9200", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("连接服务失败失败: %s", err)
		return
	}
	proto.GRPCConn = conn

	type args struct {
		token  string
		userID uint32
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"1", args{
			token:  "1234567",
			userID: 7,
		}}, {"2", args{
			token:  "1234567",
			userID: 7,
		}}, {"3", args{
			token:  "333",
			userID: 0,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user.ExpireTokenByID(tt.args.token, tt.args.userID)
		})
	}
}
