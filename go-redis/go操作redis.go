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
		Password:    "123456",
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

func RedisList() {
	DB.RPush("list", "zhangsan", "lisi", "wangwu", "xiaoming")
	fmt.Println(DB.LLen("list"))
	fmt.Println(DB.LRange("list", 0, -1).Val())
}

func RedisHash() {
	DB.HSet("info", "name", "zhangsan")
	DB.HSet("info", "age", 18)
	fmt.Println(DB.HGet("info", "name").Val())   //zhangsan
	fmt.Println(DB.HGet("info", "age").Val())    //18
	fmt.Println(DB.HGetAll("info").Val())        //map[age:18 name:zhangsan]
	fmt.Println(DB.HKeys("info").Val())          //[name age]
	fmt.Println(DB.HLen("info").Val())           //2
	fmt.Println(DB.HDel("info", "name").Val())   //1
	fmt.Println(DB.HKeys("info").Val())          //[age]
	fmt.Println(DB.HExists("info", "age").Val()) //true
}

func RedisSet() {
	DB.SAdd("set", "a", "b", "c", "d")
	fmt.Println(DB.SIsMember("set", "a").Val()) //true
	fmt.Println(DB.SMembers("set").Val())       //[a b c d]
	fmt.Println(DB.SRem("set", "d").Val())      //1
	fmt.Println(DB.SCard("set").Val())          //3

	DB.SAdd("set1", 1, 2, 3)
	DB.SAdd("set2", 2, 3, 4)
	fmt.Println(DB.SDiff("set1", "set2").Val())  //[1]
	fmt.Println(DB.SInter("set1", "set2").Val()) //[2 3]
	fmt.Println(DB.SUnion("set1", "set2").Val()) //[1 2 3 4]
}

func RedisZset() {
	DB.ZAdd("class", redis.Z{Score: 80, Member: "zhangsan"}, redis.Z{Score: 40, Member: "lisi"}, redis.Z{Score: 30, Member: "wangwu"})
	fmt.Println(DB.ZCard("class").Val())                                               //3
	fmt.Println(DB.ZRange("class", 0, -1).Val())                                       //[wangwu lisi zhangsan]
	fmt.Println(DB.ZScore("class", "lisi").Val())                                      //40
	fmt.Println(DB.ZRank("class", "zhangsan").Val())                                   //2
	fmt.Println(DB.ZCount("class", "20", "50").Val())                                  //2
	fmt.Println(DB.ZRangeByScore("class", redis.ZRangeBy{Min: "20", Max: "50"}).Val()) //[wangwu lisi]
	fmt.Println(DB.ZRevRangeWithScores("class", 0, -1).Val())                          //[{80 zhangsan} {40 lisi} {30 wangwu}]
}

func RedisPipeLine() {
	DB.Pipelined(func(tx redis.Pipeliner) error {
		tx.Set("age", 18, 0)
		return nil
	})
	fmt.Println(DB.Get("age").Int())
}

func RedisPipeLine1() {
	tx := DB.Pipeline()
	// 添加命令到管道
	tx.Set("age", 18, 0)
	getCmd := tx.Get("age")
	// 执行管道
	_, err := tx.Exec()
	if err != nil {
		fmt.Println("执行失败:", err)
		return
	}
	if age, err := getCmd.Result(); err == nil {
		fmt.Println("age的值:", age) // 输出: age的值: 18
	} else {
		fmt.Println("获取失败:", err)
	}
}

func RedisWatch() {
	DB.Watch(func(tx *redis.Tx) error {
		_, err := tx.Pipelined(func(pipe redis.Pipeliner) error {
			time.Sleep(5 * time.Second)
			pipe.Set("age", 19, 0)
			return nil
		})
		if err != nil {
			fmt.Println("事务不成功")
			return err
		}
		fmt.Println(tx.Get("age").Int())
		return nil
	}, "age")
}

func main() {
	RedisClient()
	//RedisString()
	//RedisString1()
	//RedisList()
	//RedisHash()
	//RedisSet()
	//RedisZset()
	//RedisPipeLine()
	//RedisPipeLine1()
	RedisWatch()
}
