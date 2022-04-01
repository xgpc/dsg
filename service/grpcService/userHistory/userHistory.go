// Package userHistory
// @Author:        asus
// @Description:   $
// @File:          userHistory
// @Data:          2022/4/116:53
//
package userHistory

import (
	"context"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/service/grpcService/proto"
)

func Add(userID uint32, SysCode string, Body []byte) {
	c := proto.NewUserHistoryClient(proto.GRPCConn)

	_, err := c.Add(context.Background(), &proto.UserHistoryAddReq{
		UserID:  userID,
		SysCode: SysCode,
		Body:    Body,
	})

	if err != nil {
		exce.ParseErr(err)
	}
}

func Query(userID, StartAt, EndAt uint32, SysCode []string) *proto.UserHistoryQueryRes {
	c := proto.NewUserHistoryClient(proto.GRPCConn)

	res, err := c.Query(context.Background(), &proto.UserHistoryQueryReq{
		UserID:  userID,
		StartAt: StartAt,
		EndAt:   EndAt,
		SysCode: SysCode,
	})

	if err != nil {
		exce.ParseErr(err)
	}
	return res
}
