package models

import (
	//"context"
	"fmt"
	"github.com/go-redis/redis/v7"
)

var (
	RedisDb *redis.Client
)

// 创建redis连接
func init() {
	//var ctx = context.Background()
	ip := Config.Section("redis").Key("ip").String()
	port := Config.Section("redis").Key("port").String()
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", ip, port),
		Password: "",
		DB:       0,
	})
	_, err := RedisDb.Ping().Result()
	if err != nil {
		println(err)
	}
}
