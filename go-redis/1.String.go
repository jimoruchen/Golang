package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var DB *redis.Client

func RedisClient() {
	client := redis.NewClient(&redis.Options{
		Addr:        "localhost:6379",
		Password:    "",
		DB:          0,
		DialTimeout: 1 * time.Second,
	})
	err := client.Ping().Err()
	if err != nil {
		panic("redis连接失败")
	}
	DB = client
}

func RedisString() {
	DB.Set("name", "xxx", 0)
	stringCmd := DB.Get("name")
	fmt.Println(stringCmd.Val()) //字符串
	fmt.Println(stringCmd.Result())
	fmt.Println(stringCmd.Int()) //数字
	fmt.Println(stringCmd.Err())
}

func RedisString1() {
	DB.Set("name", "xxx", 0)
	DB.Set("age", 18, 0)
	fmt.Println(DB.Get("name").Val()) //xxx
	fmt.Println(DB.Get("age").Int())  //18 <nil>

	fmt.Println(DB.Exists("name").Val())   //1
	fmt.Println(DB.Incr("age").Val())      //19
	fmt.Println(DB.IncrBy("age", 1).Val()) //20
	fmt.Println(DB.Decr("age").Val())      //19
	fmt.Println(DB.DecrBy("age", 1).Val()) //18
	fmt.Println(DB.Del("age").Val())       //1

	fmt.Println(DB.TTL("name")) //-1s
	DB.Set("name", "yyy", 5*time.Second)
	fmt.Println(DB.TTL("name")) //5s
	time.Sleep(5 * time.Second)
	fmt.Println(DB.TTL("name")) //-2s

	DB.Set("name", "zzz", 0)
	DB.Expire("name", 2*time.Second)
	time.Sleep(2 * time.Second)
	fmt.Println(DB.TTL("name")) //-2s
}

func main() {
	RedisClient()
	RedisString1()
}
