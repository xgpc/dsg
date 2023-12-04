package main

import (
	"context"
	"fmt"
	"github.com/xgpc/dsg/v2"
)

func RedisDemo() {
	dsg.Load("config.yml")
	fmt.Println(dsg.Conf)

	redisConn := dsg.Redis()
	background := context.Background()

	redisConn.Set(background, "test", "test", 0)
	redisConn.Get(background, "test")

	redisConn.HSet(background, "test", "test", "test")
	redisConn.HGet(background, "test", "test")

	redisConn.HMSet(background, "test", map[string]interface{}{"test": "test"})
	redisConn.HMGet(background, "test", "test")

	redisConn.HGetAll(background, "test")

	redisConn.LPush(background, "test", "test")
	redisConn.LRange(background, "test", 0, 1)

	redisConn.SAdd(background, "test", "test")
	redisConn.SMembers(background, "test")

}
