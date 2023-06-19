package sign

type Config struct {
	NumDaysAgo int64 `json:"num_days_ago"` // 距今多少天, 为第一天(至少为当前天数 - 20天, 防止偏移超出范围)
}
