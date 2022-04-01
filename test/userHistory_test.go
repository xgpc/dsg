// Package userHistory
// @Author:        asus
// @Description:   $
// @File:          userHistory_test.go
// @Data:          2022/4/116:58
//
package test

import (
	"fmt"
	"github.com/xgpc/dsg/service/grpcService/proto"
	"github.com/xgpc/dsg/service/grpcService/userHistory"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	// 监听端口
	conn, err := grpc.Dial("127.0.0.1:9200", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("连接服务失败失败: %s", err)
		return
	}
	proto.GRPCConn = conn

	type args struct {
		userID  uint32
		AppName string
		Body    []byte
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "1", args: args{
			userID:  7,
			AppName: "xxx",
			Body:    []byte(`{"1":2}`),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userHistory.Add(tt.args.userID, tt.args.AppName, tt.args.Body)
		})
	}
}

func TestQuery(t *testing.T) {
	// 监听端口
	conn, err := grpc.Dial("127.0.0.1:9200", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("连接服务失败失败: %s", err)
		return
	}
	proto.GRPCConn = conn

	type args struct {
		userID  uint32
		StartAt uint32
		EndAt   uint32
		SysCode []string
	}
	tests := []struct {
		name string
		args args
		want *proto.UserHistoryQueryRes
	}{
		// TODO: Add test cases.
		{
			"1", args{
				userID:  7,
				StartAt: 1648815236,
				EndAt:   1648815394,
			}, nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := userHistory.Query(tt.args.userID, tt.args.StartAt, tt.args.EndAt, tt.args.SysCode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Query() = %v, want %v", got, tt.want)
			}
		})
	}
}
