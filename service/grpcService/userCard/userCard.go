// Package userCard
// @Author:        asus
// @Description:   $
// @File:          userCard
// @Data:          2022/4/116:46
package userCard

import (
	"context"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/service/grpcService/proto"
	"github.com/xgpc/dsg/util"
)

func GetUserCard(UserID uint32) *proto.UserCardRes {
	c := proto.NewUserCardClient(proto.GRPCConn)
	res, err := c.GetUserCard(context.Background(), &proto.UserCardReq{UserID: UserID})
	if err != nil {
		exce.ParseErr(err)
	}
	return res
}

type UserCard struct {
	Gender    bool   `json:"gender"`     // 性别
	RealName  string `json:"real_name"`  // 真实姓名
	UserImage string `json:"user_image"` // 用户头像
	CardID    string `json:"card_id"`    // 身份证号码
	CardType  string `json:"card_type"`  // 证件类型: 第一代15位  第二代18位, 考虑老人情况
	StartAt   uint32 `json:"start_at"`   // 有效期:开始时间
	EndAt     uint32 `json:"end_at"`     // 有效期:结束时间  0为长期有效
	CardCode  string `json:"card_code"`  // 籍贯代码
	AreaCode  string `json:"area_code"`  // 行政区域代码
	Address   string `json:"address"`    // 常用地址
	Email     string `json:"email"`      // 邮箱
}

func UpdateUserCard(UserID uint32, data map[string]interface{}) {
	body, err := util.JsonEncode(data)
	if err != nil {
		exce.ThrowSys(exce.CodeSysBusy, err.Error())
	}

	c := proto.NewUserCardClient(proto.GRPCConn)
	_, err = c.UpdateUserCard(context.Background(), &proto.UpdateUserCardReq{
		UserID: UserID,
		Body:   body,
	})
	if err != nil {
		exce.ParseErr(err)
	}
}

func GetAddressList(userID uint32) *proto.UserCardRes {
	c := proto.NewUserCardClient(proto.GRPCConn)
	res, err := c.GetAddressList(context.Background(), &proto.UserCardReq{UserID: userID})
	if err != nil {
		exce.ParseErr(err)
	}
	return res
}

func InstallAddress(UserID uint32, data string) {
	c := proto.NewUserCardClient(proto.GRPCConn)
	_, err := c.InstallAddress(context.Background(), &proto.UpdateUserCardReq{
		UserID: UserID,
		Body:   []byte(data),
	})
	if err != nil {
		exce.ParseErr(err)
	}
}

func UpAddress(UserID uint32, data []string) {
	body, err := util.JsonEncode(data)
	if err != nil {
		exce.ThrowSys(exce.CodeSysBusy, err.Error())
	}

	c := proto.NewUserCardClient(proto.GRPCConn)
	_, err = c.UpAddress(context.Background(), &proto.UpdateUserCardReq{
		UserID: UserID,
		Body:   body,
	})
	if err != nil {
		exce.ParseErr(err)
	}
}
