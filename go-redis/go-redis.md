# Golang

## g0-redis

### 连接redis
```go
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

func main() {
	RedisClient()
	fmt.Println(DB.Get("name"))
}
```

<hr>

### redis五大类型

#### String

字符串 String 是 Redis 最简单的数据结构，可以存储字符串、整数或者浮点数。最常见的应用场景就是对象缓存。

##### 基本操作
```shell
127.0.0.1:6379> set name xxx
OK
127.0.0.1:6379> get name
"xxx"
127.0.0.1:6379> exists name
(integer) 1
127.0.0.1:6379> exists name1
(integer) 0
127.0.0.1:6379> del name
(integer) 1
127.0.0.1:6379> get name
(nil)
127.0.0.1:6379> getset name xxx   # 如果不存在值，则返回 nil，set是会执行的
(nil)
127.0.0.1:6379> getset name xxx   # 如果存在值，获取原来的值，并设置新的值
"xxx"
```

##### 批量操作
```shell
127.0.0.1:6379> mset name1 xxx name2 yyy
OK
127.0.0.1:6379> get name1
"xxx"
127.0.0.1:6379> get name2
"yyy"
```

##### 计数
自增1，自减1
```shell
127.0.0.1:6379> get n
(nil)
127.0.0.1:6379> incr n
(integer) 1
127.0.0.1:6379> get n
"1"
127.0.0.1:6379> incr n
(integer) 2
127.0.0.1:6379> get n
"2"
127.0.0.1:6379> decr n
(integer) 1
127.0.0.1:6379> get n
"1"
127.0.0.1:6379> decr n
(integer) 0
127.0.0.1:6379> get n
"0"
```

##### 自增n，自减n
```shell
127.0.0.1:6379> get n
"0"
127.0.0.1:6379> incrby n 10
(integer) 10
127.0.0.1:6379> get n
"10"
127.0.0.1:6379> decrby n 10
(integer) 0
127.0.0.1:6379> get n
"0"
```

##### 过期操作
setex 单位是秒
```shell
127.0.0.1:6379> setex name 10 xxx
OK
127.0.0.1:6379> ttl name    # 查看还有多久过期
(integer) 7
127.0.0.1:6379> ttl name    # -2表示已过期
(integer) -2
127.0.0.1:6379> get name
(nil)
```

##### expire
设置一个key的过期时间
```shell
127.0.0.1:6379> set name yyy
OK
127.0.0.1:6379> ttl name    # 现在是永不过期
(integer) -1
127.0.0.1:6379> expire name 10
(integer) 1
127.0.0.1:6379> ttl name
(integer) 5
127.0.0.1:6379> ttl name    # 现在是过期了
(integer) -2
```

使用场景

登录
限流
计数器

<hr>

#### List

<img src="https://s2.loli.net/2025/10/30/Tl5xBkp6FOSLygG.png"  alt="">

##### 基本操作
```shell
127.0.0.1:6379> rpush list wangwu xiaoming
(integer) 2
127.0.0.1:6379> lrange list 0 -1
1) "wangwu"
2) "xiaoming"
127.0.0.1:6379> lpush list lisi zhangsan
(integer) 4
127.0.0.1:6379> lrange list 0 -1
1) "zhangsan"
2) "lisi"
3) "wangwu"
4) "xiaoming"
127.0.0.1:6379> lrange list 0 0
1) "zhangsan"
127.0.0.1:6379> lrange list 1 2
1) "lisi"
2) "wangwu"
127.0.0.1:6379> rpop list
"xiaoming"
127.0.0.1:6379> lpop list
"zhangsan"
127.0.0.1:6379> lrange list 0 -1
1) "lisi"
2) "wangwu"
127.0.0.1:6379> llen list
(integer) 2
```

使用场景

任务队列
排行榜
分页查询

<hr>

#### Hash

相当于go里面的map，python中的字典。

##### 基本操作
```shell
127.0.0.1:6379> hset info name xxx
(integer) 1
127.0.0.1:6379> hset info age 18
(integer) 1
127.0.0.1:6379> hget info name
"xxx"
127.0.0.1:6379> hlen info     # 查询哈希表 info 中字段的数量。
(integer) 2
127.0.0.1:6379> hkeys info    # 获取 info 哈希中的所有字段名。
1) "name"
2) "age"
127.0.0.1:6379> hexists info name
(integer) 1
127.0.0.1:6379> hexists info age
(integer) 1
127.0.0.1:6379> hgetall info  # 获取哈希表中所有的字段及其值。
1) "name"
2) "xxx"
3) "age"
4) "18"
127.0.0.1:6379> hdel info age
(integer) 1
127.0.0.1:6379> hkeys info
1) "name"
```

使用场景

记录网站某篇文章的浏览量
存储配置信息

<hr>

#### Set

集合中的元素没有先后顺序。

HashSet就是基于HashMap来实现的，HashSet，他其实就是说一个集合，里面的元素是无序的，他里面的元素不能重复的。

##### 基本操作
```shell
127.0.0.1:6379> sadd set a b c d
(integer) 4
127.0.0.1:6379> smembers set
1) "c"
2) "a"
3) "d"
4) "b"
127.0.0.1:6379> sadd set a e
(integer) 1
127.0.0.1:6379> smembers set
1) "a"
2) "d"
3) "b"
4) "c"
5) "e"
127.0.0.1:6379> scard set
(integer) 5
127.0.0.1:6379> sismember set a
(integer) 1
127.0.0.1:6379> sismember set f
(integer) 0
127.0.0.1:6379> srem set e
(integer) 1
127.0.0.1:6379> smembers set
1) "d"
2) "c"
3) "a"
4) "b"
```

##### 交集、并集和差集

sdiff 差集
sinter 交集
sunion 并集

```shell
127.0.0.1:6379> sadd set1 1 2 3
(integer) 3
127.0.0.1:6379> sadd set2 2 3 4
(integer) 3
127.0.0.1:6379> sdiff set1 set2
1) "1"
127.0.0.1:6379> sdiff set2 set1
1) "4"
127.0.0.1:6379> sinter set1 set2
1) "2"
2) "3"
127.0.0.1:6379> sunion set1 set2
1) "1"
2) "2"
3) "3"
4) "4"
```

##### 随机抽奖

使用spop命令用于移除集合中的指定 key 的一个或多个随机元素，移除后会返回移除的元素

```shell
127.0.0.1:6379> sadd order 1 2 3 4 5 6 7 8
(integer) 8
127.0.0.1:6379> spop order
"3"
127.0.0.1:6379> spop order
"1"
127.0.0.1:6379> spop order 2
1) "2"
2) "7"
127.0.0.1:6379> smembers order
1) "4"
2) "5"
3) "6"
4) "8"
```

使用场景

共同好友
统计网站的独立ip
标签

<hr>

#### sorted set

有序集合

sorted set 增加了一个权重参数 score，使得集合中的元素能够按 score 进行有序排列

也是不能重复的

##### 基本操作
```shell
127.0.0.1:6379> zadd class 88 fengfeng 46 zhangsan 76 wangwu 36 lisi
(integer) 4
127.0.0.1:6379> zrange class 0 -1    # 通过索引，按照分数从低到高返回
1) "lisi"
2) "zhangsan"
3) "wangwu"
4) "fengfeng"
127.0.0.1:6379> zrevrange class 0 -1    # 通过索引，按照分数从高到低返回
1) "fengfeng"
2) "wangwu"
3) "zhangsan"
4) "lisi"
127.0.0.1:6379> zcard class   # 获取有序集合的成员数
(integer) 4
127.0.0.1:6379> zscore class zhangsan   # 查看某个成员的分数
"46"
127.0.0.1:6379> zrank class zhangsan    # 查看成员的排名 从小到大排序
(integer) 1
127.0.0.1:6379> zrevrank class zhangsan      # 查看成员的排名 从大到小排序
(integer) 2
127.0.0.1:6379> zcount class 40 70      # 计算在有序集合中指定区间分数的成员数
(integer) 1
127.0.0.1:6379> zrangebyscore class 40 70     # 通过分数返回有序集合指定区间内的成员
1) "zhangsan"
127.0.0.1:6379> zrevrangebyscore class 80 40      # 返回有序集中指定分数区间内的成员，分数从高到低排序
1) "wangwu"
2) "zhangsan"
127.0.0.1:6379> zrange class 0 -1 withscores      # 把成员和分数一起显示出来
1) "lisi"
2) "36"
3) "zhangsan"
4) "46"
5) "wangwu"
6) "76"
7) "fengfeng"
8) "88"
127.0.0.1:6379> zrem class zhangsan        # 移除一个成员
(integer) 1
127.0.0.1:6379> zrange class 0 -1
1) "lisi"
2) "wangwu"
3) "fengfeng"
127.0.0.1:6379> zremrangebyrank class 0 0     # 移除有序集合中给定的排名区间的所有成员(第一名是0)(低到高排序) 现在是把lisi移除了
(integer) 1
127.0.0.1:6379> zrange class 0 -1
1) "wangwu"
2) "fengfeng"
127.0.0.1:6379> zscore class wangwu
"76"
127.0.0.1:6379> zincrby class 2 wangwu      # 给王五+2
"78"
127.0.0.1:6379> zincrby class -2 wangwu     # 给王五-2
"76"
127.0.0.1:6379> zscore class wangwu
"76"
```

使用场景

排行榜
订单支付超时（下单时插入，member为订单号，score为订单超时时间戳，然后写个定时任务每隔一段时间执行zrange）

<hr>

### 通过go操作五大数据类型

#### String
```shell
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
```

<hr>

#### 