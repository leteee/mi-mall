package models

import (
	//"context"
	"fmt"
	"github.com/go-redis/redis/v7"
	"gopkg.in/ini.v1"
	"os"
)

var (
	RedisDb *redis.Client
)

// 创建redis连接
func init() {

	//读取.ini里面的数据库配置
	config, iniErr := ini.Load("./conf/app.ini")
	if iniErr != nil {
		fmt.Printf("Fail to read file: %v", iniErr)
		os.Exit(1)
	}
	//var ctx = context.Background()
	ip := config.Section("redis").Key("ip").String()
	port := config.Section("redis").Key("port").String()
	password := config.Section("redis").Key("password").String()
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", ip, port),
		Password: password,
		DB:       0,
	})
	_, err := RedisDb.Ping().Result()
	if err != nil {
		println(err)
	}
}
