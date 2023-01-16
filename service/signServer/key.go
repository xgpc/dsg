package signServer

import (
	"fmt"
	"time"
)

// getUserKey  sign:user:userID
func getUserKey(userID uint32) string {

	return fmt.Sprintf("%s:%d", TagUser, userID)
}

// sign:today:total:year:day
func getKeyTotalToday() string {
	t := time.Now().Format("2006-01-02")
	return fmt.Sprintf("%s:%s", TagTodayTotal, t)
}

func getKeyTotalYesterday() string {
	t := time.Now().AddDate(0, 0, -1)
	return fmt.Sprintf("%s:%d:%d", TagTodayTotal, t.Year(), t.Day())
}
