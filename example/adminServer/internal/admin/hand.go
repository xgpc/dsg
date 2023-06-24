/**
 * @Author: smono
 * @Description:
 * @File:  hand
 * @Version: 1.0.0
 * @Date: 2022/9/25 12:47
 */

package admin

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/xgpc/dsg/pkg/json"
	"gorm.io/gorm"
)

var (
	_db      *gorm.DB
	_redis   *redis.Client
	redisTag *string
)

type Info struct {
	UserID    uint32 `json:"user_id"`
	SubjectID uint32 `json:"subject_id"`
	IsAdmin   bool   `json:"is_admin"`
	IsSuper   bool   `json:"is_super"`
	IsEmploy  bool   `json:"is_employ"`
}

func (*Info) TableName() string {
	return "admin" + *redisTag
}

func GetUser(userID, SubjectID uint32) Info {
	key := *redisTag + ":" + string(SubjectID)
	resString := Redis().HGet(context.Background(), key, string(userID)).String()

	var info Info

	err := json.Decode([]byte(resString), &info)
	if err != nil {
		panic(err)
	}
	return info
}

func SetUser(info Info) error {
	marshal, err := json.Encode(info)
	if err != nil {
		panic(err)
	}
	_, err = Redis().HSet(context.Background(), *redisTag+":"+string(info.SubjectID), info.UserID, string(marshal)).Result()
	return err

}

func Init(db *gorm.DB, redisConn *redis.Client) {
	_db = db
	_redis = redisConn
}

func DB() *gorm.DB {
	return _db
}

func Redis() *redis.Client {
	return _redis
}
