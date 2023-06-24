/**
 * @Author: smono
 * @Description:
 * @File:  middleware
 * @Version: 1.0.0
 * @Date: 2022/9/28 21:32
 */

package apps

import (
	"company/admin/services/user"
	"github.com/kataras/iris/v12"
	"github.com/xgpc/dsg/exce"
	"github.com/xgpc/dsg/pkg"
	"github.com/xgpc/dsg/util"
	"strconv"
)

// 所有权限管理 基于会话的ID来操作

func Check(ctx iris.Context) {

	uuid := ctx.GetHeader("app-key")
	if uuid != _uuid {
		exce.ThrowSys(exce.CodeRequestError, "服务器与客户端秘钥对接失败, 请检查")
	}

	sessionData := ctx.GetHeader("company-data")
	var sessionInfo user.SessionData
	err := util.JsonDecode([]byte(sessionData), &sessionInfo)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}
	SetSession(ctx, sessionInfo)

	ctx.Next()
}

func Login(ctx iris.Context) {

	token := ctx.GetHeader("token")
	if token != "" {
		info := GetSession(token, ctx)
		if info.UserID == 0 {
			exce.ThrowSys(exce.CodeUserNoLogin)
		}
		ctx.Values().Set("mid", info.UserID)
		ctx.Values().Set("session", info)
	} else {
		exce.ThrowSys(exce.CodeUserNoLogin)
	}

	ctx.Next()
}

func SetSession(ctx iris.Context, data user.SessionData) {
	ctx.Values().Set("session", data)
	// 同步兼容 myId()
	ctx.Values().Set("mid", data.UserID)
}

func GetSession(key string, ctx iris.Context) SessionData {

	token, err := pkg.ParseToken(key)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	var info SessionData
	info.UserID = token.UserID

	header := ctx.GetHeader("subject_id")
	if header == "" {
		header = "0"
	}

	subjectID, err := strconv.Atoi(header)
	if err != nil {
		exce.ThrowSys(exce.CodeUserNoAuth, "subject 解析失败, 请检查 header中是否携带subject_id")
	}

	info.SubjectID = uint32(subjectID)

	return info
}

type SessionData struct {
	UserID    uint32 `json:"user_id"`
	SubjectID uint32 `json:"subject_id"`
}

func (p *SessionData) String() string {
	marshal, err := util.JsonEncode(p)
	if err != nil {
		exce.ThrowSys(exce.CodeRequestError, err.Error())
	}

	return string(marshal)
}
