package frame

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type method interface {
	ErrRedisGet(key string) (string, bool)
	RedisGet(key string) string

	ErrRedisSet(key, value string, sec int) bool
	RedisSet(key, value string, sec int)

	ErrRedisDel(key string) bool
	RedisDel(key string)
}

var ctx = context.Background()

type redisConn struct {
	Conn *redis.Client
}

func NewRedisConn(conn *redis.Client) *redisConn {
	return &redisConn{Conn: conn}
}

func RedisConn() *redisConn {
	conn := RedisDefault()
	conn.Do(context.Background(), "select", 0)
	return &redisConn{Conn: conn}
}

func (r *redisConn) ErrRedisGet(key string) (string, bool) {
	if r.Conn == nil {
		LogError("redis get -> redis nil", key)
		return "", false
	}

	str, err := r.Conn.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", true
	}
	if err != nil {
		LogError("redis get -> error:"+err.Error(), key)
		return "", false
	}

	return str, true
}

func (conn *redisConn) RedisGet(key string) string {
	str, flag := conn.ErrRedisGet(key)
	if !flag {
		//log.ThrowSys("RedisGet", key)
	} //	todo
	return str
}

//
//// ***************************** set *****************************
//

func (conn *redisConn) ErrRedisSet(key, value string, sec int) bool {
	if conn == nil {
		LogError("redis set -> redis nil", key+" => "+value)
		return false
	}

	err := conn.Conn.Set(ctx, key, value, time.Duration(sec)*time.Second).Err()
	if err != nil {
		LogError("redis set -> error:"+err.Error(), key+" => "+value)
		return false
	}
	return true
}

func (conn *redisConn) RedisSet(key, value string, sec int) {
	if !conn.ErrRedisSet(key, value, sec) {
		//log.ThrowSys("RedisSet", key+" => "+value)
	} //	todo
}

//// ***************************** del *****************************

func (conn *redisConn) ErrRedisDel(key string) bool {
	if conn.Conn == nil {
		LogError("redis del -> redis nil", key)
		return false
	}

	err := conn.Conn.Del(ctx, key).Err()
	if err != nil {
		LogError("redis del -> error:"+err.Error(), key)
		return false
	}
	return true
}

func (conn *redisConn) RedisDel(key string) {
	if !conn.ErrRedisDel(key) {
		//log.ThrowSys("RedisDel", key)
	} //	todo
}

//// ***************************** scan *****************************

func (conn *redisConn) ErrRedisScan(cursor uint64, match string, count int64) ([]string, uint64, bool) {
	if conn == nil {
		LogError("redis scan -> redis nil", match)
		return nil, 0, false
	}

	keys, newCursor, err := conn.Conn.Scan(ctx, cursor, match, count).Result()
	if err == redis.Nil {
		return []string{}, 0, true
	}

	if err != nil {
		LogError("redis scan -> error:"+err.Error(), match)
		return nil, 0, false
	}

	return keys, newCursor, true
}

func (conn *redisConn) RedisScan(cursor uint64, match string, count int64) ([]string, uint64) {
	keys, newCursor, flag := conn.ErrRedisScan(cursor, match, count)
	if !flag {
		//log.ThrowSys("RedisScan", match)
	} //	todo
	return keys, newCursor
}

//// ***************************** hscan *****************************

func (conn *redisConn) ErrRedisHScan(key string, cursor uint64, match string, count int64) ([]string, uint64, bool) {
	if conn == nil {
		LogError("redis hscan -> redis nil", match)
		return nil, 0, false
	}

	keys, newCursor, err := conn.Conn.HScan(ctx, key, cursor, match, count).Result()
	if err == redis.Nil {
		return []string{}, 0, true
	}

	if err != nil {
		LogError("redis hscan -> error:"+err.Error(), match)
		return nil, 0, false
	}

	return keys, newCursor, true
}

func (conn *redisConn) RedisHScan(key string, cursor uint64, match string, count int64) ([]string, uint64) {
	keys, newCursor, flag := conn.ErrRedisHScan(key, cursor, match, count)
	if !flag {
		//log.ThrowSys("RedisHScan", match)
	} //	todo
	return keys, newCursor
}

//// ***************************** hget *****************************

func (conn *redisConn) ErrRedisHGet(key, field string) (string, bool) {
	if conn == nil {
		LogError("redis hget -> redis nil", key+" => "+field)
		return "", false
	}

	str, err := conn.Conn.HGet(ctx, key, field).Result()
	if err == redis.Nil {
		return "", true
	}
	if err != nil {
		LogError("redis hget -> error:"+err.Error(), key+" => "+field)
		return "", false
	}

	return str, true
}

func (conn *redisConn) RedisHGet(key, field string) string {
	str, flag := conn.ErrRedisHGet(key, field)
	if !flag {
		//log.ThrowSys("RedisHGet", key+" => "+field)
	} //	todo
	return str
}

//// ***************************** hdel *****************************

func (conn *redisConn) ErrRedisHDel(key, field string) bool {
	if conn == nil {
		LogError("redis hdel -> redis nil", key+" => "+field)
		return false
	}
	err := conn.Conn.HDel(ctx, key, field).Err()
	if err != nil {
		LogError("redis hdel -> error:"+err.Error(), key+" => "+field)
		return false
	}

	return true
}

func (conn *redisConn) RedisHDel(key, field string) {
	if !conn.ErrRedisHDel(key, field) {
		//log.ThrowSys("RedisHDel", key+" => "+field)
	} //	todo
}

//// ***************************** hset *****************************

func (conn *redisConn) ErrRedisHSet(key, field, value string) bool {
	if conn == nil {
		LogError("redis hset -> redis nil", key+" => "+field)
		return false
	}

	err := conn.Conn.HSet(ctx, key, field, value).Err()
	if err != nil {
		LogError("redis hset -> error:"+err.Error(), key+" => "+field)
		return false
	}
	return true
}

func (conn *redisConn) RedisHSet(key, field, value string) {
	if !conn.ErrRedisHSet(key, field, value) {
		//log.ThrowSys("RedisHSet", key+" => "+field)
	} //	todo
}

//// ***************************** rpush *****************************

func (conn *redisConn) ErrRedisPush(key, value string) bool {
	if conn == nil {
		LogError("redis rpush -> redis nil", key)
		return false
	}

	err := conn.Conn.RPush(ctx, key, value).Err()
	if err != nil {
		LogError("redis rpush -> error:"+err.Error(), key+" => "+value)
		return false
	}
	return true
}

func (conn *redisConn) RedisPush(key, value string) {
	if !conn.ErrRedisPush(key, value) {
		//log.ThrowSys("RedisPush", key)
		//	todo
	}
}

//// ***************************** lpop *****************************

func (conn *redisConn) ErrRedisPop(key string) (string, bool) {
	if conn == nil {
		LogError("redis lpop -> redis nil", key)
		return "", false
	}

	bs, err := conn.Conn.LPop(ctx, key).Bytes()
	if err == redis.Nil {
		return "", true
	}
	if err != nil {
		LogError("redis lpop -> error:"+err.Error(), key)
		return "", false
	}
	return string(bs), true
}

func (conn *redisConn) RedisPop(key string) string {
	str, flag := conn.ErrRedisPop(key)
	if !flag {
		//log.ThrowSys("RedisPop", key)
		//	todo
	}
	return str
}

// ***************************** lrange *****************************

func (conn *redisConn) ErrRedisLRange(key string, start, end int64) ([]string, bool) {
	if conn == nil {
		LogError("redis lrange -> redis nil", key)
		return []string{}, false
	}

	datum, err := conn.Conn.LRange(ctx, key, start, end).Result()
	if err == redis.Nil {
		return []string{}, true
	}
	if err != nil {
		LogError("redis lrange -> error:"+err.Error(), key)
		return []string{}, false
	}
	return datum, true
}

func (conn *redisConn) RedisLRange(key string, start, end int64) []string {
	datum, flag := conn.ErrRedisLRange(key, start, end)
	if !flag {
		//log.ThrowSys("RedisLRange", key)
	} //	todo
	return datum
}

// ***************************** expire *****************************

func (conn *redisConn) ErrRedisExpire(key string, lifetime int) bool {
	if conn == nil {
		LogError("redis expire -> redis nil", key)
		return false
	}

	err := conn.Conn.Expire(ctx, key, time.Duration(lifetime)*time.Second).Err()
	if err != nil {
		LogError("redis expire -> error:"+err.Error(), key)
		return false
	}
	return true
}

func (conn *redisConn) RedisExpire(key string, lifetime int) {
	if !conn.ErrRedisExpire(key, lifetime) {
		//log.ThrowSys("RedisExpire", key)
	} //	todo
}

// ***************************** exists *****************************

func (conn *redisConn) ErrRedisExists(key string) (bool, bool) {
	if conn == nil {
		LogError("redis exists -> redis nil", key)
		return false, false
	}

	v, err := conn.Conn.Exists(ctx, key).Result()
	if err != nil {
		LogError("redis exists -> error:"+err.Error(), key)
		return false, false
	}
	return v > 0, true
}

func (conn *redisConn) RedisExists(key string) bool {
	f, flag := conn.ErrRedisExists(key)
	if !flag {
		//log.ThrowSys("RedisExists", key)
	} //	todo
	return f
}

//// ***************************** hexists *****************************

func (conn *redisConn) ErrRedisHExists(key, field string) (bool, bool) {
	if conn == nil {
		LogError("redis hexists -> redis nil", key)
		return false, false
	}

	v, err := conn.Conn.HExists(ctx, key, field).Result()
	if err != nil {
		LogError("redis hexists -> error:"+err.Error(), key)
		return false, false
	}
	return v, true
}

func (conn *redisConn) RedisHExists(key, field string) bool {
	f, flag := conn.ErrRedisHExists(key, field)
	if !flag {
		//log.ThrowSys("RedisHExists", key)
	} //	todo
	return f
}

// ***************************** incr *****************************

func (conn *redisConn) ErrRedisIncr(key string) (int64, bool) {
	if conn == nil {
		LogError("redis incr -> redis nil", key)
		return 0, false
	}

	v, err := conn.Conn.Incr(ctx, key).Result()
	if err != nil {
		LogError("redis incr -> error:"+err.Error(), key)
		return 0, false
	}
	return v, true
}

func (conn *redisConn) RedisIncr(key string) int64 {
	v, flag := conn.ErrRedisIncr(key)
	if !flag {
		//log.ThrowSys("RedisIncr", key)
	} //	todo
	return v
}

// ***************************** sadd *****************************

func (conn *redisConn) ErrRedisSadd(key string, members ...interface{}) (int64, bool) {
	if conn == nil {
		LogError("redis Sadd -> redis nil", key)
		return 0, false
	}

	v, err := conn.Conn.SAdd(ctx, key, members...).Result()
	if err != nil {
		LogError("redis Sadd -> error:"+err.Error(), key)
		return 0, false
	}
	return v, true
}

func (conn *redisConn) RedisSadd(key string, members ...interface{}) int64 {
	v, flag := conn.ErrRedisSadd(key, members...)
	if !flag {
		//log.ThrowSys("RedisSadd", key)
	} //	todo
	return v
}

func (conn *redisConn) ErrRedisSMembers(key string) ([]string, bool) {
	if conn == nil {
		LogError("redis SMembers -> redis nil", key)
		return nil, false
	}

	v, err := conn.Conn.SMembers(ctx, key).Result()
	if err != nil {
		LogError("redis SMembers -> error:"+err.Error(), key)
		return nil, false
	}
	return v, true
}

func (conn *redisConn) RedisSMembers(key string) []string {
	v, flag := conn.ErrRedisSMembers(key)
	if !flag {
		//log.ThrowSys("RedisSMembers", key)
	} //	todo
	return v
}

// set 类型
//***************************** 有序合集 set *****************************
func (conn *redisConn) ErrRedisZADD(key string, members ...*redis.Z) (int64, bool) {
	if conn == nil {
		LogError("redis ZAdd -> redis nil", key)
		return 0, false
	}

	v, err := conn.Conn.ZAdd(ctx, key, members...).Result()
	if err != nil {
		LogError("redis ZAdd -> error:"+err.Error(), key)
		return 0, false
	}
	return v, true
}

func (conn *redisConn) RedisZADD(key string, members ...*redis.Z) int64 {
	v, flag := conn.ErrRedisZADD(key, members...)
	if !flag {
		//log.ThrowSys("RedisZIncrBy", key)
		return 0
	}
	return v
}

func (conn *redisConn) ErrRedisZIncrBy(key string, increment float64, member string) (float64, bool) {
	if conn == nil {
		LogError("redis ZIncrBy -> redis nil", key)
		return 0, false
	}

	v, err := conn.Conn.ZIncrBy(ctx, key, increment, member).Result()
	if err != nil {
		LogError("redis ZIncrBy -> error:"+err.Error(), key)
		return 0, false
	}
	return v, true
}

func (conn *redisConn) RedisZIncrBy(key string, increment float64, member string) float64 {

	v, flag := conn.ErrRedisZIncrBy(key, increment, member)
	if !flag {
		//log.ThrowSys("RedisZIncrBy", key)
		return 0
	}
	return v
}

//***************************** ZRange *****************************
func (conn *redisConn) ErrRedisZRange(key string, start, stop int64) ([]string, bool) {
	if conn == nil {
		LogError("redis ZRange -> redis nil", key)
		return nil, false
	}

	v, err := conn.Conn.ZRange(ctx, key, start, stop).Result()
	if err != nil {
		LogError("redis ZRange -> error:"+err.Error(), key)
		return nil, false
	}
	return v, true
}

func (conn *redisConn) RedisZRange(key string, start, stop int64) []string {

	v, flag := conn.ErrRedisZRange(key, start, stop)
	if !flag {
		//log.ThrowSys("RedisZRange", key)
		return nil
	}
	return v
}

//***************************** ZRevRange *****************************
func (conn *redisConn) ErrRedisZRevRange(key string, start, stop int64) ([]string, bool) {
	if conn == nil {
		LogError("redis ZRevRange -> redis nil", key)
		return nil, false
	}

	v, err := conn.Conn.ZRevRange(ctx, key, start, stop).Result()
	if err != nil {
		LogError("redis ZRevRange -> error:"+err.Error(), key)
		return nil, false
	}
	return v, true
}

func (conn *redisConn) RedisZRevRange(key string, start, stop int64) []string {

	v, flag := conn.ErrRedisZRevRange(key, start, stop)
	if !flag {
		//log.ThrowSys("RedisZRange", key)
		return nil
	}
	return v
}

//***************************** ZRevRangeWithScores *****************************
func (conn *redisConn) ErrRedisZRevRangeWithScores(key string, start, stop int64) ([]redis.Z, bool) {
	if conn == nil {
		LogError("redis ZRevRangeWithScores -> redis nil", key)
		return nil, false
	}

	v, err := conn.Conn.ZRevRangeWithScores(ctx, key, start, stop).Result()
	if err != nil {
		LogError("redis ZRevRangeWithScores -> error:"+err.Error(), key)
		return nil, false
	}
	return v, true
}

func (conn *redisConn) RedisZRevRangeWithScores(key string, start, stop int64) []redis.Z {

	v, flag := conn.ErrRedisZRevRangeWithScores(key, start, stop)
	if !flag {
		//log.ThrowSys(1000, "RedisZRange", key)
		return nil
	}
	return v
}
