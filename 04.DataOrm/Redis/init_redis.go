package model

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis"
)

// RedisClient Redis缓存客户端单例
var RedisClient *redis.Client

// Redis 在中间件中初始化redis链接
func Redis() {
	db, _ := strconv.ParseUint(os.Getenv("REDIS_DB"), 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PW"),
		DB:       int(db),
	})

	_, err := client.Ping().Result()

	if err != nil {
		log.Panicf("连接Redis不成功 %s", err)
	}

	RedisClient = client
}

func RedisPing() {
	pong, err := RedisClient.Ping().Result()
	fmt.Println(pong, err)
}

func Example() {
	// 设置键值
	err := RedisClient.Set("key", "value", 0).Err()
	if err != nil {
		log.Panicln(err)
	}

	// 获取KEY的值
	val, err := RedisClient.Get("key").Result()
	fmt.Println(val, err)

	val2, err := RedisClient.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}
