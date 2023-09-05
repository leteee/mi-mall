package models

import (
	//"context"
	"github.com/go-redis/redis/v7"
)

var (
	RedisDb *redis.Client
)

// 创建redis连接
func init() {
	//var ctx = context.Background()
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	_, err := RedisDb.Ping().Result()
	if err != nil {
		println(err)
	}
}
