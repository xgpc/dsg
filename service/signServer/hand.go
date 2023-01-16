package signServer

import (
	redis2 "github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	_db  *gorm.DB
	_rdb *redis2.Client
)

const (
	TagTodayTotal = "sign:today" // 今日签到
	TagUser       = "sign:user"  // 签到人
)

func Init(db *gorm.DB, rdb *redis2.Client) {
	_rdb = rdb
	_db = db
}

func db() *gorm.DB {
	return _db
}

func redis() *redis2.Client {
	return _rdb
}
