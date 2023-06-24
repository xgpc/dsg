/**
 * @Author: smono
 * @Description:
 * @File:  isSuper
 * @Version: 1.0.0
 * @Date: 2022/9/25 12:58
 */

package admin

import "github.com/xgpc/dsg/exce"

func IsSuper(userID, subjectID uint32) bool {
	info := getUser(userID, subjectID)
	if info.IsSuper == 1 {
		return true
	}
	return false
}

func CheckSuper(userID, subjectID uint32) {
	if !IsSuper(userID, subjectID) {
		exce.ThrowSys(exce.CodeUserNoAuth, "您不是超级管理员")
	}
}
