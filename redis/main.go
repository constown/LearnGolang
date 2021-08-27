package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

// 推荐书籍 redis实战

// redis的用处：
//	1、cache缓存
//	2、简单的队列
//	3、排行榜

var redisDb *redis.Client

func initRedis() (err error) {
	redisDb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	_, err = redisDb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := initRedis()
	if err != nil {
		fmt.Printf("connect redis failed, err:%v", err)
		return
	}
	fmt.Println("连接redis成功")

	//	zset
	key := "rank"
	items := []redis.Z{
		redis.Z{Score: 97, Member: "PHP"},
		redis.Z{Score: 98, Member: "Java"},
		redis.Z{Score: 94, Member: "Golang"},
		redis.Z{Score: 96, Member: "JavaScript"},
	}
	redisDb.ZAdd(key, items...)
	// 加分数
	newScore, err := redisDb.ZIncrBy(key, 10, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed,err:%v", err)
		return
	}
	fmt.Printf("golang's score if %f now.\n", newScore)
}
