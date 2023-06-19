package sign

import (
	"context"
	"errors"
	"time"
)

const (
	Sign3Day = 7   // 7 => 0000 0111
	Sign7Day = 127 // 127   => 01111 1111
)

// SignUser 人员打卡, 返回今日第几个签到
func SignUser(userID uint32) (int64, error) {
	offset := getTodayNum()
	key := getUserKey(userID)
	result, err := redis().SetBit(context.Background(), key, offset, 1).Result()
	if err != nil {
		panic(err)
	}

	if result == 0 {
		// 每次签到数据考虑存入数据库
		saveDB(userID)
		// 记录今日统计数据
		return SignToday(), nil
	}
	return 0, errors.New("已经打卡")
}

func saveDB(id uint32) {
	db := db()
	if db == nil {
		return
	}
	// TODO: 暂不实现
}

// SignUserCheckToday 今日是否已经打卡
func SignUserCheckToday(userID uint32) bool {
	offset := getTodayNum()
	key := getUserKey(userID)
	result, err := redis().GetBit(context.Background(), key, offset).Result()
	if err != nil {
		panic(err)
	}

	return result == 1
}

// SignUserGetUint8 获取近8日打卡数据
func SignUserGetUint8(userID uint32) []int64 {
	offset := getTodayNum()
	key := getUserKey(userID)

	// 0000 0000 0000 0000 0000 0000 0000 0000
	//                       ^        ^  = 7  从0开始计数
	result, err := redis().BitField(context.Background(), key, "get", "u16", offset-15).Result()
	if err != nil {
		panic(err)
	}

	return result
}

type SignDay struct {
	TimeAt  int64  `json:"time_at"`
	TimeStr string `json:"time_str"`
	IsSign  bool   `json:"is_sign"`
}

type SignInfo struct {
	TimeAt  int64  `json:"time_at"`  // 时间戳
	TimeStr string `json:"time_str"` // 打卡时间
	IsSign  bool   `json:"is_sign"`  //是否打卡
}

type SignData struct {
	Consecutive3   bool       // 连续三天
	Consecutive7   bool       // 连续七天
	ConsecutiveNum int        // 连续签到几天
	List           []SignInfo // 打卡情况
}

// SignUserGetInfo 获取近8日打卡情况
func SignUserGetInfo(userID uint32) (*SignData, error) {
	var res SignData
	list := SignUserGetUint8(userID)
	// 正常情况只有一条
	if len(list) < 1 {
		return nil, errors.New("打卡数据失败, 请联系管理员")
	}

	// 取前14天数据 进行连续打卡判断
	var num int
	data := list[0]
	// 近7天数据
	if data&Sign7Day == Sign7Day {
		res.Consecutive7 = true
	}

	if data&Sign3Day == Sign3Day {
		res.Consecutive3 = true
	}

	// 计算当天
	t := time.Now()
	tcur := time.
		Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	var info SignInfo
	info.IsSign = (data & 1) == 1
	info.TimeAt = tcur.Unix()
	info.TimeStr = tcur.String()
	res.List = append(res.List, info)

	// 0000 0000 0000 0001
	//                  0<-  下标从0开始, 所以是倒叙来的, 但是不计算当天打卡.
	for i := 1; i < 16; i++ {
		var info SignInfo
		t := time.Now()
		tcur := time.
			Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).
			AddDate(0, 0, -1*i)
		data = data >> 1
		info.IsSign = (data & 1) == 1
		info.TimeAt = tcur.Unix()
		info.TimeStr = tcur.String()
		res.List = append(res.List, info)

		if !info.IsSign || num == 1 {
			// 当有一次未打卡, 就不在进行连续打卡计算
			num = 1
		} else {
			res.ConsecutiveNum += 1
		}
	}

	return &res, nil
}
