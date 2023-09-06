package models

import (
	"fmt"
	"time"
)

const CAPTCHA = "captcha:"

type RedisStore struct {
}

func (r RedisStore) Set(id, value string) error {
	key := CAPTCHA + id
	err := RedisDb.Set(key, value, time.Minute*2).Err()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (r RedisStore) Get(id string, clear bool) string {
	key := CAPTCHA + id
	val, err := RedisDb.Get(key).Result()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if clear {
		err := RedisDb.Del(key).Err()
		if err != nil {
			fmt.Println(err)
			return ""
		}
	}
	return val
}

// Verify 实现验证captcha的方法
func (r RedisStore) Verify(id, answer string, clear bool) bool {
	v := RedisStore{}.Get(id, clear)
	return v == answer
}
