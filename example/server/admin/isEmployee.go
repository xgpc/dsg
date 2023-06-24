package admin

import "github.com/xgpc/dsg/exce"

func IsEmployee(userID, subjectID uint32) bool {
	info := getUser(userID, subjectID)
	if info.UserID != 0 {
		return true
	}
	return false
}

func CheckEmployee(userID, subjectID uint32) {
	if !IsAdmin(userID, subjectID) {
		exce.ThrowSys(exce.CodeUserNoAuth, "您不该主体员工")
	}
}
