package sign

import (
	"context"
	"time"
)

var CurrentKey string

// SignToday 统计今日签到人数
// 返回当前第几个签到
func SignToday() int64 {
	key := getKeyTotalToday()
	result, err := redis().Incr(context.Background(), key).Result()
	if err != nil {
		panic(err)
	}

	// 防止试过过多
	if CurrentKey != key {
		// 一个月失效
		_, err := redis().Expire(context.Background(), key, time.Hour*24*30).Result()
		if err != nil {
			panic(err)
		}
		// 更换key
		CurrentKey = key
	}

	return result
}

// SignGetTodayNum  统计今日签到人数
func SignGetTodayNum() int64 {

	key := getKeyTotalToday()
	result, err := redis().Get(context.Background(), key).Int64()
	if err != nil {
		panic(err)
	}
	return result
}
