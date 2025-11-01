# Golang

## g0-redis

## docker安装Redis

### 拉取 Redis 镜像
`docker pull redis`

### 运行 Redis 容器
`docker run --name my-redis -d -p 6379:6379 ^
-v D:/my-redis:/data ^
redis redis-server --requirepass 123456 --appendonly yes`

### 验证是否运行成功
`docker ps`

### 使用 redis-cli 连接测试
`docker exec -it my-redis redis-cli`
`auth 123456`

<hr>

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

func main() {
	RedisClient()
	//RedisString()
	//RedisString1()
	//RedisList()
	//RedisHash()
	//RedisSet()
	RedisZset()
}
```

<hr>

### redis常用命令

```shell
C:\Users\10078>docker exec -it my-redis redis-cli
127.0.0.1:6379> auth 123456
OK
127.0.0.1:6379> select 1
OK
127.0.0.1:6379[1]> select 0
OK
127.0.0.1:6379> keys *
1) "class"
2) "set1"
3) "set2"
4) "info"
5) "set"
6) "list"
127.0.0.1:6379> keys s*
1) "set1"
2) "set2"
3) "set"
127.0.0.1:6379> ping
PONG
127.0.0.1:6379> move class 1
(integer) 1
127.0.0.1:6379> select 1
OK
127.0.0.1:6379[1]> keys *
1) "class"
127.0.0.1:6379[1]> type class
zset
127.0.0.1:6379[1]> select 0
OK
127.0.0.1:6379> setex name 20 xxx
OK
127.0.0.1:6379> ttl name
(integer) 17
127.0.0.1:6379> persist name    # 移除key的过期时间，key永不过期
(integer) 1
127.0.0.1:6379> ttl name
(integer) -1
127.0.0.1:6379> client list     # 获取连接到服务器的客户端连接列表
id=10 addr=127.0.0.1:42738 laddr=127.0.0.1:6379 fd=25 name= age=7097 
idle=7072 flags=N db=0 sub=0 psub=0 ssub=0 multi=-1 watch=0 qbuf=0 
qbuf-free=0 argv-mem=0 multi-mem=0 rbs=1024 rbp=0 obl=0 oll=0 omem=0 
tot-mem=1920 events=r cmd=get user=default redir=-1 resp=2 lib-name= 
lib-ver= io-thread=0 tot-net-in=142 tot-net-out=146 tot-cmds=2
127.0.0.1:6379> dbsize          # 获取当前数据库key的数量
(integer) 6
127.0.0.1:6379> select 1  
OK
127.0.0.1:6379[1]> dbsize       # 获取当前数据库key的数量
(integer) 1
127.0.0.1:6379[1]> flushdb      # 删除当前数据库的所有key
OK
127.0.0.1:6379[1]> dbsize
(integer) 0
127.0.0.1:6379[1]> FLUSHALL     # 删除全部数据库的所有key
OK
```

<hr>

### 事务

Redis 事务的本质是一组命令的集合。事务支持一次执行多个命令，一个事务中所有命令都会被序列化。
在事务执行过程，会按照顺序串行化执行队列中的命令，其他客户端提交的命令请求不会插入到事务执行命令序列中。
因为我们的程序是并发的，你在一个程序里面设置值，然后取值，这很正常
但是如果并发存在，那么肯定就会存在，取值的时候不是我自己设置的那个值
基于上面的问题，那我在一个客户端操作的时候，把所有的指令一次性按照顺序排他的放在一个队列中，执行完了之后再让其他的客户端操作

实际上在Redis中也会出现多个命令同时竞争同一个数据的情况，比如现在有两条命令同时执行，他们都要去修改a的值，那么这个时候就只能动用锁机制来保证同一时间只能有一个命令操作。

虽然Redis中也有锁机制，但是它是一种乐观锁，不同于MySQL，我们在MySQL中认识的锁是悲观锁，那么什么是乐观锁什么是悲观锁呢？

悲观锁：时刻认为别人会来抢占资源，禁止一切外来访问，直到释放锁，具有强烈的排他性质。
乐观锁：并不认为会有人来抢占资源，所以会直接对数据进行操作，在操作时再去验证是否有其他人抢占资源。
Redis中可以使用watch来监视一个目标，如果执行事务之前被监视目标发生了修改，则取消本次事务：

1. MULTI —— 开启事务
标记一个事务块的开始。
2. EXEC —— 执行事务
   执行所有在 MULTI 和 EXEC 之间的命令。
3. DISCARD —— 取消事务
   放弃事务，不执行任何命令。
4. WATCH —— 监视键
   乐观锁机制：如果被监视的键在事务执行前被修改，则整个事务不会执行。

A
```shell
127.0.0.1:6379> multi
OK
127.0.0.1:6379(TX)> set a 100
QUEUED
127.0.0.1:6379(TX)> get a
QUEUED
127.0.0.1:6379(TX)> exec
1) OK
2) "100"
127.0.0.1:6379> get a
"100"
```
B
```shell
127.0.0.1:6379> set a 10000
OK
127.0.0.1:6379> get a
"100"
```

A
```shell
127.0.0.1:6379> watch a
OK
127.0.0.1:6379> multi
OK
127.0.0.1:6379(TX)> set a 200
QUEUED
127.0.0.1:6379(TX)> exec
(nil)
```
B
```shell
127.0.0.1:6379> get a
"100"
127.0.0.1:6379> set a 300
OK
```

A
```shell
127.0.0.1:6379> watch a
OK
127.0.0.1:6379> multi
OK
127.0.0.1:6379(TX)> set a 100
QUEUED
127.0.0.1:6379(TX)> exec
(nil)
127.0.0.1:6379> unwatch
```
B
```shell
127.0.0.1:6379> set a 200
OK
127.0.0.1:6379> set a 100
OK
```
ABA问题

```shell
127.0.0.1:6379> set age 20
OK
```
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
```

<hr>

### 持久化

#### RDB持久化
默认是开启的，默认的存储文件是 dump.rdb

1.配置自动备份:
默认是 一分钟内修改了一万次，5分钟内修改了10次，30分钟内修改了1次
```shell
save 3600 1
save 300 100
save 60 10000
```
2.手动命令备份:
save：save时只管保存，其他不管，全部阻塞，手动保存，不建议使用。
bgsave：redis会在后台异步进行快照操作，快照同时还可以响应客户端情况。
可以使用lastsave命令获取最后一次成功生成快照的时间（时间戳）

<hr>

#### AOF持久化
虽然RDB能够很好地解决数据持久化问题，但是它的缺点也很明显：每次都需要去完整地保存整个数据库中的数据，
同时后台保存过程中也会产生额外的内存开销， 最严重的是它并不是实时保存的，如果在自动保存触发之前服务器崩溃
，那么依然会导致少量数据的丢失。 而AOF就是另一种方式，它会以日志的形式将我们每次执行的命令都进行保存，
服务器重启时会将所有命令依次执行， 通过这种重演的方式将数据恢复，这样就能很好解决实时性存储问题。

```shell
appendonly yes
appendfilename "appendonly.aof"
appendfsync always
dir ./
```
appendfsync always：每次写入立即同步
appendfsync everysec：每秒同步
appendfsync no：不主动同步