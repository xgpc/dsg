/**
 * @Author: smono
 * @Description:
 * @File:  isAdmin
 * @Version: 1.0.0
 * @Date: 2022/9/25 12:49
 */

package admin

import "github.com/xgpc/dsg/exce"

func IsAdmin(userID, subjectID uint32) bool {
	info := GetUser(userID, subjectID)
	return info.IsAdmin
}

func CheckAdmin(userID, subjectID uint32) {
	if !IsAdmin(userID, subjectID) {
		exce.ThrowSys(exce.CodeUserNoAuth, "您不是管理员, 无权操作")
	}
}
