package signServer

import "time"

const (
	OneDayTime = 3600 * 24
)

// 获取1970年1月1日 距今多少天 - 19300
func getTodayNum() int64 {
	t := time.Now().Unix()
	d := int64(t / OneDayTime)
	tm := t % OneDayTime

	if tm > 0 {
		d = d + 1
	}
	// 数字太多不利于存储, 因为会获取近7天的数据, 所以 尽量不要让减去的数据接近 1970-1-1 至今的天 靠近
	return d - 19300
}
