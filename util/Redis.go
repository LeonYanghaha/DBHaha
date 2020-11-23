package util

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var ctxList = make(map[string]RedisInfo)

func PutCtx(id string, dbInfo RedisInfo) {
	ctxList[id] = dbInfo
}

func GetCtx(id string) bool {
	if _, ok := ctxList[id]; ok {
		return true
	}
	return false
}

type RedisInfo struct {
	Addr      string
	Password  string
	DB        int
	MyClient  redis.Client
	MyContext context.Context
}

func GetRedisClient(redisInfo RedisInfo) RedisInfo {

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisInfo.Addr,
		Password: redisInfo.Password,
		DB:       redisInfo.DB,
	})
	redisInfo.MyClient = *rdb
	redisInfo.MyContext = context.Background()
	return redisInfo
}

// 获取所有的key
func (r RedisInfo) Keys(pattern string) *string {
	s := r.MyClient.Keys(r.MyContext, "*")
	//TODO 明天从这里开始 。。。。。
	println(s)
	return nil
}

//设置指定 key 的值
func (r RedisInfo) Set(k string, v interface{}, ttl time.Duration) *redis.StatusCmd {
	s := r.MyClient.Set(r.MyContext, k, v, ttl)
	return s
}

//获取指定 key 的值
func (r RedisInfo) Get(k string) *redis.StringCmd {
	return r.MyClient.Get(r.MyContext, k)
}

//返回 key 中字符串值的子字符
func (r RedisInfo) GetRange(k string, start, end int64) *redis.StringCmd {
	return r.MyClient.GetRange(r.MyContext, k, start, end)
}

//将给定 key 的值设为 value ，并返回 key 的旧值(old value)
func (r RedisInfo) GetSet(k string, v interface{}) *redis.StringCmd {
	return r.MyClient.GetSet(r.MyContext, k, v)
}

//对 key 所储存的字符串值，获取指定偏移量上的位(bit)
func (r RedisInfo) GetBit(k string, offset int64) *redis.IntCmd {
	return r.MyClient.GetBit(r.MyContext, k, offset)
}

//获取所有(一个或多个)给定 key 的值
func (r RedisInfo) MGet(keys ...string) *redis.SliceCmd {
	return r.MyClient.MGet(r.MyContext, keys...)
}

//对 key 所储存的字符串值，设置或清除指定偏移量上的位(bit)
func (r RedisInfo) SetBit(key string, offset int64, value int) *redis.IntCmd {
	return r.MyClient.SetBit(r.MyContext, key, offset, value)
}

//将值 value 关联到 key ，并将 key 的过期时间设为 seconds (以秒为单位)
// TODO SETEX
//func (r RedisInfo) SETEX(key string, seconds int, value interface{}) *redis.StringCmd {
//	return r.MyClient.se
//}

//只有在 key 不存在时设置 key 的值
func (r RedisInfo) SetNX(key string, value interface{}, ttl time.Duration) *redis.BoolCmd {
	return r.MyClient.SetNX(r.MyContext, key, value, ttl)
}

//用 value 参数覆写给定 key 所储存的字符串值，从偏移量 offset 开始
func (r RedisInfo) SetRange(key string, offset int64, value string) *redis.IntCmd {
	return r.MyClient.SetRange(r.MyContext, key, offset, value)
}

//返回 key 所储存的字符串值的长度
func (r RedisInfo) StrLen(key string) *redis.IntCmd {
	return r.MyClient.StrLen(r.MyContext, key)
}

//删除一个或多个哈希表字段
func (r RedisInfo) Hdel(key string, field ...string) *redis.IntCmd {
	return r.MyClient.HDel(r.MyContext, key, field...)
}

//查看哈希表 key 中，指定的字段是否存在
func (r RedisInfo) HExists(key, field string) *redis.BoolCmd {
	return r.MyClient.HExists(r.MyContext, key, field)
}

//获取存储在哈希表中指定字段的值
func (r RedisInfo) HGet(key, field string) *redis.StringCmd {
	return r.MyClient.HGet(r.MyContext, key, field)
}

//获取在哈希表中指定 key 的所有字段和值
func (r RedisInfo) HGetAll(key string) *redis.StringStringMapCmd {
	return r.MyClient.HGetAll(r.MyContext, key)
}

//为哈希表 key 中的指定字段的整数值加上增量 increment
func (r RedisInfo) HIncrBy(key, field string, increment int64) *redis.IntCmd {
	return r.MyClient.HIncrBy(r.MyContext, key, field, increment)
}

//为哈希表 key 中的指定字段的浮点数值加上增量 increment
func (r RedisInfo) HIncrByFloat(key, field string, increment float64) *redis.FloatCmd {
	return r.MyClient.HIncrByFloat(r.MyContext, key, field, increment)
}

//获取所有哈希表中的字段
func (r RedisInfo) HKyes(key string) *redis.StringSliceCmd {
	return r.MyClient.HKeys(r.MyContext, key)
}

//获取哈希表中字段的数量
func (r RedisInfo) HLen(key string) *redis.IntCmd {
	return r.MyClient.HLen(r.MyContext, key)
}

//获取所有给定字段的值
func (r RedisInfo) HMGet(key string, field ...string) *redis.SliceCmd {
	return r.MyClient.HMGet(r.MyContext, key, field...)
}

//同时将多个 field-value (域-值)对设置到哈希表 key 中
func (r RedisInfo) HMSet(key string, value ...interface{}) *redis.BoolCmd {
	return r.MyClient.HMSet(r.MyContext, key, value...)
}

//将哈希表 key 中的字段 field 的值设为 value
func (r RedisInfo) HSet(key string, value ...interface{}) *redis.IntCmd {
	return r.MyClient.HSet(r.MyContext, key, value...)
}

//只有在字段 field 不存在时，设置哈希表字段的值
func (r RedisInfo) HSetNX(key string, field string, value interface{}) *redis.BoolCmd {
	return r.MyClient.HSetNX(r.MyContext, key, field, value)
}

//获取哈希表中所有值
func (r RedisInfo) HVals(key string) *redis.StringSliceCmd {
	return r.MyClient.HVals(r.MyContext, key)
}

//迭代哈希表中的键值对。
func (r RedisInfo) HScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	return r.MyClient.HScan(r.MyContext, key, cursor, match, count)
}

func ExecCmd(id, cmd, param string) {
	if !GetCtx(id) {
		fmt.Println("严重错误")
		return
	}
	var temp = ctxList[id]
	temp.Keys("")
}
