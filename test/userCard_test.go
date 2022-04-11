// Package userCard
// @Author:        asus
// @Description:   $
// @File:          userCard_test.go
// @Data:          2022/4/116:59
//
package test

import (
	"fmt"
	"github.com/xgpc/dsg/service/grpcService/proto"
	"github.com/xgpc/dsg/service/grpcService/userCard"
	"github.com/xgpc/dsg/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"reflect"
	"testing"
)

func TestGetUserCard(t *testing.T) {

	// 监听端口
	conn, err := grpc.Dial("127.0.0.1:9200", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("连接服务失败失败: %s", err)
		return
	}
	proto.GRPCConn = conn

	type args struct {
		UserID uint32
	}
	tests := []struct {
		name string
		args args
		want *proto.UserCardRes
	}{
		// TODO: Add test cases.
		{"1", args{UserID: 7}, &proto.UserCardRes{UserID: 7, RealName: "段", UserImage: "123", CardID: "321"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := userCard.GetUserCard(tt.args.UserID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserCard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateUserCard(t *testing.T) {
	// 监听端口
	conn, err := grpc.Dial("127.0.0.1:9200", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("连接服务失败失败: %s", err)
		return
	}
	proto.GRPCConn = conn
	type args struct {
		UserID uint32
		data   map[string]interface{}
	}

	data := util.StructToMapByRef(userCard.UserCard{
		Gender:    false,
		RealName:  "老段",
		UserImage: "321",
		CardID:    "1111111",
		CardType:  "15",
		StartAt:   1,
		EndAt:     2,
		CardCode:  "1",
		AreaCode:  "2",
		Address:   "3",
		Email:     "4",
	})

	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"1", args{
			UserID: 7,
			data:   data,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userCard.UpdateUserCard(tt.args.UserID, tt.args.data)
		})
	}
}

func TestAddress(t *testing.T) {
	// 监听端口
	conn, err := grpc.Dial("127.0.0.1:9200", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("连接服务失败失败: %s", err)
		return
	}
	proto.GRPCConn = conn
	type args struct {
		UserID uint32
		data   map[string]interface{}
	}

	res := userCard.GetAddressList(7)
	fmt.Println(res)

	userCard.UpAddress(7, []string{})

	res = userCard.GetAddressList(7)
	fmt.Println(res)

	userCard.InstallAddress(7, "昆明4楼")
	res = userCard.GetAddressList(7)
	fmt.Println(res)
}
