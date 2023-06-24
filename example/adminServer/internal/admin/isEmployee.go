package admin

import "github.com/xgpc/dsg/exce"

func IsEmployee(userID, subjectID uint32) bool {
	info := GetUser(userID, subjectID)
	return info.UserID != 0
}

func CheckEmployee(userID, subjectID uint32) {
	if !IsAdmin(userID, subjectID) {
		exce.ThrowSys(exce.CodeUserNoAuth, "您不该主体员工")
	}
}
