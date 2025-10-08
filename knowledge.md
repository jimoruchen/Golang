# Golang

## golang基础

### 变量定义与输入输出

#### 1.变量
```go
package main

import "fmt"

var e = 100 //全局可以定义不使用

func main() {
	var a int
	a = 100
	fmt.Println(a)

	var b int = 100
	fmt.Println(b)

	var c = 100
	fmt.Println(c)

	d := 100
	fmt.Println(d)

	fmt.Println(e)

	var a1, a2, a3 = "1", "2", "3"
	fmt.Println(a1, a2, a3)

	var (
		x1 = 100
		x2 = 100
	)
	fmt.Println(x1, x2)

	const version = "v3"
	fmt.Println(version)
}
```

#### 2.输出
```go
package main

import "fmt"

func main() {
	var name = "xx"
	fmt.Println("xxx")
	fmt.Println(name)

	fmt.Print("1", "xxx")
	fmt.Print("2", "xx\n")

	fmt.Printf("\"xx\"\n")

	name = "xx"
	fmt.Printf("%q %T\n", name, name)

	var age = 18
	fmt.Printf("%d %T\n", age, age)

	s := fmt.Sprintf("我的年龄1是%d", age)
	fmt.Println(s)
}
```

#### 3.输入
```go
package main

import "fmt"

func main() {
	var name string
	var age int
	n, err := fmt.Scan(&name, &age)
	//n, err := fmt.Scanln(&name, &age)
	fmt.Println(n, err, name, age)
}
```

<hr>

### 基本数据类型

#### 整数类型
```go
package main

import "fmt"

func main() {
	var a uint8 = 255
	fmt.Printf("%0b %d\n", a, a) //11111111 255

	var b int8 = 127
	fmt.Printf("%0b %d\n", b, b) // 1111111 127
	//0    0000000    0
	//0    1111111    127
	//1    0000000    -128
	b = -128
	fmt.Printf("%0b %d\n", b, b)

	var c = 9223372036854775807
	fmt.Printf("%0b %d\n", c, c)  //64位操作系统，int就为int64，2^63 - 1
	var d = 9223372036854775808.0 //超过64位
	fmt.Println(d)
}
```

#### 浮点型
```go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.MaxFloat32) //3.4028234663852886e+38
	fmt.Println(math.MaxFloat64) //1.7976931348623157e+308

	a := 1.0 //float64类型
	fmt.Println(a)
	b := float64(1)
	fmt.Println(b)

	println(math.Inf(1) > 100000000000000000)
	println(math.Inf(-1) < -10000000000000000)
}

```

#### 字符型
```go
package main

import "fmt"

func main() {
	var c1 byte //int8  0-255
	c1 = 98
	var c2 = 'b'
	fmt.Println(c1, c2)
	fmt.Printf("%c, %c\n", c1, c2)

	var c3 rune
	c3 = '上'
	fmt.Println(c3)
	fmt.Printf("%c\n", c3)
}
```

#### 字符串类型
```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("你好")
	fmt.Println("'你'好")   //单引号
	fmt.Println("\"你\"好") //双引号要转置
	fmt.Println("\\你\\好") //输出斜杠要两个
	fmt.Println("你\n好")   //换行
	fmt.Println("你\t好")   //中间空几格
	fmt.Println("你\r好")   //回到行首
	var a = `你
\
好`
	fmt.Println(a) //多行字符串，``中\反斜杠失效

	println(strings.Count("111000111000", "1"))       //统计字符出现个数
	println(strings.TrimSpace("   jimo ruchen   "))   //去除空格，中间的不行
	println(strings.HasSuffix("xxx.jpg", ".jpg"))     //判断后缀
	println(strings.HasPrefix("user-xxx", "user"))    //判断前缀
	println(strings.Contains("user-xxx", "x"))        //判断包含
	println(strings.Replace("user-xxx", "x", "y", 3)) //替换3个
	println(strings.ReplaceAll("user-xxx", "x", "y")) //全部替换
	println(strings.ToLower("AbC"))
	println(strings.ToUpper("AbC"))
	println(strings.ToTitle("hello world!"))
}
```

#### 布尔类型
```go
package main

import "fmt"

func main() {
	a := true
	b := false
	fmt.Println(a, b)

	var c bool
	fmt.Println(c) //默认为false
}
```

#### 零值问题
```go
package main

import "fmt"

func main() {
	var a int
	var b uint
	var c byte
	var d bool
	var e string
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Printf("%q\n", e)
}
```

<hr>

### 数组、切片、map

#### 数组
```go
package main

import "fmt"

func main() {
	var names1 [3]string
	names1 = [3]string{"zhangsan", "lisi", "wangwu"}
	fmt.Println(names1[0])

	var names2 = [3]string{"zhangsan", "lisi", "wangwu"}
	fmt.Println(names2[0])

	for i, s := range names2 {
		fmt.Println(i, s)
	}

	fmt.Println(len(names1))

	for i := 0; i < len(names1); i++ {
		fmt.Println(names1[i])
	}
}
```
```go
package main

func finalPrices(prices []int) []int {
	length := len(prices)
	//var ans = make([]int, length)
	var ans []int
	for i := 0; i < length; i++ {
		var discount int = 0
		for j := i + 1; j < length; j++ {
			if (prices[j] <= prices[i]) {
				discount = prices[j]
				break;
			}
		}
		ans = append(ans, prices[i] - discount)     //如果未定义长度必须使用append添加元素，否则不能使用索引
	}
	return ans
}
```

#### 切片
````go
package main

import (
	"fmt"
	"sort"
)

func main() {
	var a = [5]int{1, 2, 3, 4, 5}
	// 0  1  2  3  4
	fmt.Println(a)
	var s = a[1:3] // [start: end)
	fmt.Println(s)

	var b = []int{1, 2, 3, 4, 5}
	fmt.Println(b)
	b = append(b, 6)
	fmt.Println(b)
	s1 := append(b[0:3], b[4:]...) //func append(slice []T, elems ...T)，...表示解包将b[4:]的每个元素拿出来
	fmt.Println(s1)
	s1[0] = 100
	fmt.Println(s1)

	var c = make([]int, 3)
	fmt.Println(c)
	var d []int
	fmt.Println(d == nil, d)
	var e = make([]int, 0)
	fmt.Println(e == nil, e)
	f := make([]int, 0)
	fmt.Println(f == nil, f)

	g1 := []int{2, 4, 1, 5, 3}
	sort.Ints(g1)
	fmt.Println(g1)

	g2 := []int{2, 4, 1, 5, 3}
	sort.Slice(g2, func(i, j int) bool {
		return g2[i] < g2[j]
	})
	fmt.Println(g2)

	g3 := []int{2, 4, 1, 5, 3}
	sort.Slice(g3, func(i, j int) bool {
		return g3[i] > g3[j]
	})
	fmt.Println(g3)
}
````

#### map
```go
package main

import "fmt"

func main() {
	var map1 map[int]string
	//map1 = make(map[int]string)
	map1 = map[int]string{}
	map1[1] = "aaa"
	map1[2] = "bbb"
	fmt.Println(map1)

	var map2 = make(map[int]string)		//初始化
	map2[1] = "aaa"
	fmt.Println(map2)

	var map3 = map[int]string{}			//初始化
	map3[1] = "aaa"
	fmt.Println(map3)

	for k, v := range map1 {			//遍历
		fmt.Println(k, v)
	}

	value := map1[1]
	fmt.Println(value)
	value1, ok := map1[1]				//查询
	fmt.Println(value1, ok)

	delete(map1, 2)				//删除
	fmt.Println(map1)
}
```

<hr>

### if语句
```go
package main

import "fmt"

func main() {
	fmt.Println("请输入你的年龄：")
	var age int
	fmt.Scan(&age)
	if age <= 0 {
		fmt.Println("未出生")
	} else if age > 0 && age <= 18 {
		fmt.Println("未成年")
	} else if age > 18 && age <= 35 {
		fmt.Println("青年")
	} else {
		fmt.Println("中年")
	}
}
```

### switch语句
```go
package main

import "fmt"

func main() {
	fmt.Println("请输入你的年龄：")
	var age int
	fmt.Scan(&age)
	switch {
	case age <= 0:
		fmt.Println("未出生")
	case age <= 18:
		fmt.Println("未成年")
		fallthrough             //判断成功后继续往下走
	case age <= 35:
		fmt.Println("青年")
	case age > 35:
		fmt.Println("中年")
	}
}
```
输出未成年和青年

```go
package main

import "fmt"

func main() {
	fmt.Println("请输入星期数字：")
	var week int
	fmt.Scan(&week)
	switch week {
	case 1:
		fmt.Println("星期一")
	case 2:
		fmt.Println("星期二")
	case 3:
		fmt.Println("星期三")
	case 4:
		fmt.Println("星期四")
	case 5:
		fmt.Println("星期五")
	case 6, 7:
		fmt.Println("周末")
	default:
		fmt.Println("错误")
	}
}
```

<hr>

### for循环

#### 死循环
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println(time.Now())
		time.Sleep(time.Second)
	}
}
```

#### 传统for循环
```go
package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}
```

#### 类似while
```go
package main

import "fmt"

func main() {
	var i, sum int
	for i <= 100 {
		sum += i
		i++
	}
	fmt.Println(sum)
}
```

#### 类似do-while
```go
package main

import "fmt"

func main() {
	i := 0
	sum := 0
	for {
		sum += i
		i++
		if i == 101 {
			break
		}
	}
	fmt.Println(sum)
}
```
```go
package main

import "fmt"

func main() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}
}
```
打印九九乘法表

<hr>

### 函数（指针）

#### 函数定义和参数
```go
package main

import "fmt"

func Print(s string) {
	fmt.Println(s)
}

func main() {
	s := "hello world"
	Print(s)
}
```
```go
package main

import "fmt"

func f1() {
	fmt.Println("f1")
}

func f2(s string) {
	fmt.Println("f2")
}

func f3(a int, b string) {
	fmt.Println("f3")
}

func f4(a, b int) {
	fmt.Println("f4")
}

func f5(list ...int) {
	fmt.Println(list)
}

func f6(name string, list ...int) {
	fmt.Println(name)
	fmt.Println(list)
}

func main() {
	f1()
	f2("hello")
	f3(1, "hello")
	f4(1, 2)
	f5(7, 8, 9)

	var list1 = []int{1, 2, 3}
	f5(list1...)

	f6("hello", list1...)
}
```
```go
package main

import "fmt"

func f1() {
	return
}

func f2() int {
	var a int = 0
	return a
}

func f3() (int, string) {
	var a int
	var b string
	return a, b
}

func f4() (res string) {
	res = "hello"
	return //裸返回
}

func f5() (res string) {
	return "hello"
}

func divide(a int, b int) (result int, err error) {
	if b == 0 {
		err = fmt.Errorf("divide by zero")
		return
	}
	result = a / b
	return //return（裸返回）的意思是：返回当前所有命名返回值变量的值。
}

func main() {
	f1()
	f2()
	fmt.Println(f3())
	fmt.Println(f4())
	fmt.Println(f5())
	result, err := divide(40, 20)
	fmt.Printf("%d, %v\n", result, err)
}
```

#### 匿名函数
```go
package main

import "fmt"

func main() {
	var add = func(a, b int) int {
		return a + b
	}
	fmt.Println(add(1, 2))

	var divide func(a, b int) int
	divide = func(a, b int) int {
		return a / b
	}
	fmt.Println(divide(1, 1))

	var divide1 = func(a, b int) int {
		return a / b
	}
	fmt.Println(divide1(1, 1))

	divide2 := func(a, b int) int {
		return a / b
	}
	fmt.Println(divide2(1, 1))
}
```

#### 高阶函数
```go
package main

import "fmt"

func main() {
	var nums int
	fmt.Scan(&nums)
	var funcMap = map[int]func(){
		1: func() {
			fmt.Printf("登录")
		},
		2: func() {
			fmt.Printf("个人中心")
		},
		3: func() {
			fmt.Println("注销")
		},
	}
	//funcMap[nums]()		//()是函数调用操作符，执行前面那个函数。
	f := funcMap[nums]
	fmt.Printf("%T\n", f)
	f()
}
```
```go
package main

import "fmt"

func login() {
	fmt.Println("login")
}

func userCenter() {
	fmt.Println("userCenter")
}

func logout() {
	fmt.Println("logout")
}

func main() {
	var nums int
	fmt.Scan(&nums)
	var funcMap = map[int]func(){
		1: login,
		2: userCenter,
		3: logout,
	}
	funcMap[nums]()
}
```
```go
package main

import "fmt"

func login(s string) {
	fmt.Println(s + " " + "login")
}

func userCenter(s string) {
	fmt.Println(s + "userCenter")
}

func logout(s string) {
	fmt.Println(s + "logout")
}

func main() {
	var nums int
	fmt.Scan(&nums)
	var funcMap = map[int]func(string){
		1: login,
		2: userCenter,
		3: logout,
	}
	funcMap[nums]("jimoruchen")
}
```

#### 闭包
```go
package main

import "fmt"

func f() func() int {
	var i int
	return func() int {     // 返回的函数是闭包
		i++                 //引用并修改外部变量 i
		return i
	}
}

func main() {
	ff := f()               //独立的闭包
	fmt.Println(ff())
	fmt.Println(ff())
	fmt.Println(ff())
}
```
闭包是一个函数，它引用了其外部作用域中的变量，即使外部函数已经执行完毕，这些变量仍然可以被内部函数访问和修改。
闭包 = 函数 + 引用的外部环境（变量）
它“封闭”了外部作用域的变量，形成一个独立的执行环境。
外部变量 i 被返回的匿名函数捕获。
即使 f 执行完毕，i 仍存活。

```go
package main

import "fmt"

func sharedCounter() (func() int, func() int) {
	count := 0
	inc := func() int {
		count++
		return count
	}
	dec := func() int {
		count--
		return count
	}
	return inc, dec
}

func main() {
	inc, dec := sharedCounter()
	fmt.Println(inc()) // 1
	fmt.Println(inc()) // 2
	fmt.Println(dec()) // 1
}
```
多个闭包共享变量，inc 和 dec 共享同一个 count 变量。

```go
package main

import (
	"fmt"
	"time"
)

func timeSleep(sleepTime int) func(...int) int {
	return func(numList ...int) int {
		time.Sleep(time.Duration(sleepTime) * time.Second)      //sleepTime是变量必须转换
		var sum int
		for _, num := range numList {
			sum += num
		}
		return sum
	}
}

func main() {
	//fmt.Println(timeSleep(2)(1, 2, 3))
	fmt.Println("开始计时")
	f := timeSleep(2)
	fmt.Println(f(1, 2, 3))
}
```
```go
package main

import "fmt"

func main() {
	var funcs []func()
	for i := 0; i < 3; i++ {
		funcs = append(funcs, func() {
			fmt.Println(i)
		})
	}
	for _, f := range funcs {
		f()
	}
}
```
从 Go 1.22 开始，这段代码会输出 0, 1, 2，
Go 1.22 修改了 for循环的语义：
每次迭代创建新的循环变量：现在每次循环迭代都会创建新的变量实例
闭包捕获的是迭代变量：闭包捕获的是当前迭代的变量副本

#### 递归
```go
package main

import "fmt"

func fibonaci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonaci(n-1) + fibonaci(n-2)
}

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", fibonaci(i))
	}
}
```

#### 值传递和引用传递
```go
package main

import "fmt"

func swap(a, b *int) (x, y *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
	return a, b
}

func main() {
	var a int
	var b int
	a, b = 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
}
```

#### 指针
```go
package main

import "fmt"

func main() {
	var a int = 100
	var pa *int = &a
	fmt.Println(pa, *pa)
	fmt.Println(&a, a)
}
```
var pa *int = &a这行代码可以分解为：
*int- 声明一个指向int类型的指针变量
pa- 变量名
&a- 获取变量a的内存地址

<hr>

### init函数和defer函数

#### init函数
```go
package main

import "fmt"

var file string

func init() {
	file = "读取配置文件"
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}

func init() {
	fmt.Println("init3")
}

func main() {
	fmt.Println("main")
	fmt.Println(file)
}
```
```go
package main

import "fmt"

var file string

func init() {
	readFile()
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}

func init() {
	fmt.Println("init3")
}

func readFile() {
	file = "读取配置文件"
}

func main() {
	fmt.Println("main")
	fmt.Println(file)
}

```
init()函数是一个特殊的函数，存在以下特性：
不能被其他函数调用，而是在main函数执行之前，自动被调用
init函数不能作为参数传入
不能有传入参数和返回值
一个go文件可以有多个init函数，谁在前面谁就先执行

#### defer函数
```go
package main

import "fmt"

func m() (x int) {
	fmt.Println("m before")
	defer func() {
		fmt.Println("before", x)
		x = 2
	}()
	fmt.Println("m after")
	x = 1
	return
}

func main() {
	fmt.Println("main before")
	x := m()
	fmt.Println(x)          //只能该命名返回值
}
```
关键字 defer 用于注册延迟调用
这些调用直到 return 前才被执。因此，可以用来做资源清理
多个defer语句，按先进后出的方式执行，谁离return近谁先执行
defer语句中的变量，在defer声明时就决定了

<hr>

### 结构体

#### 定义
```go
package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func (s Student) PrintInfo() {
	fmt.Printf("Student name: %s, Age: %d\n", s.Name, s.Age)
}

func main() {
	s1 := Student{
		Name: "zhangsan",
		Age:  18,
	}
	s1.PrintInfo()

	var s2 Student
	s2 = Student{
		Name: "lisi",
		Age:  18,
	}
	s2.PrintInfo()
}
```
```go
package main

import "fmt"

type Address struct {
	City    string
	Country string
}

type Student struct {
	Name    string
	Age     int
	Address Address
}

func (s Student) PrintInfo() {
	fmt.Printf("Student name: %s, Age: %d, City: %s, Country: %s\n", s.Name, s.Age, s.Address.City, s.Address.Country)
}

func main() {
	s1 := Student{
		Name: "zhangsan",
		Age:  18,
		Address: Address{
			City:    "chongqing",
			Country: "China",
		},
	}
	s1.PrintInfo()
}
```
```go
package main

import "fmt"

type Address struct {
	City    string
	Country string
}

type Student struct {
	Name string
	Age  int
	Address
}
//方法接收者
func (s Student) PrintInfo() {
	fmt.Printf("Student name: %s, Age: %d, City: %s, Country: %s\n", s.Name, s.Age, s.City, s.Country)
}

func main() {
	s1 := Student{
		Name: "zhangsan",
		Age:  18,
		Address: Address{
			City:    "chongqing",
			Country: "China",
		},
	}
	s1.PrintInfo()
}
```
匿名嵌套可直接访问内部结构体的字段

#### 继承
```go
package main

import "fmt"

type Father struct {
	Name  string
	House string
}

func (f Father) Work() {
	fmt.Println("Father Work")
}

type Mather struct {
	Friends int
}

type Son struct {
	Name   string
	Money  int
	Father        // 匿名组合 会将组合的属性直接挂载到对象上，前提是不重名
	Mama   Mather // 具名组合
}

func main() {
	father := Father{
		Name:  "James",
		House: "House",
	}

	mama := Mather{
		Friends: 10,
	}

	son := Son{
		Name:   "aaa",
		Money:  10,
		Father: father,
		Mama:   mama,
	}

	fmt.Println(father.House)
	fmt.Println(son.House)
	fmt.Println(son.Father.Name)
	fmt.Println(son.Name)
	fmt.Println(son.Mama.Friends)
	son.Work()
	son.Father.Work()
}
```

#### 结构体指针
```go
package main

import "fmt"

type Info struct {
	Name string
}

func setName(i *Info, name string) {
	i.Name = name
}

func main() {
	var info = Info{
		Name: "aaa",
	}
	setName(&info, "bbb")
	fmt.Println(info)
}
```
```go
package main

import "fmt"

type Info struct {
	Name string
}

func (i *Info) setName(name string) {
	i.Name = name
}

func main() {
	var info = Info{
		Name: "aaa",
	}
	info.setName("bbb") //Go 语言允许直接用结构体变量调用指针接收者方法，编译器会自动取地址,等价于 (&info).setName("bbb")
	fmt.Println(info)
}
```

#### 结构体Tag
```go
package main

import (
	"encoding/json"
	"fmt"
)

type Info struct {
	Name     string `json:"name,omitempty"`
	Age      int    `json:"age,omitempty"` // omitempty 表示如果字段是零值（""、0、false、nil），则不会出现在 JSON 输出中。
	Password string `json:"-"`             // - 表示 Password不会 被序列化或反序列化（即 JSON 中不会出现该字段）。
}

func main() {
	var info = Info{
		Name:     "sss",
		Age:      10,
		Password: "123456",
	}
	//序列化
	byteData, _ := json.Marshal(info)
	fmt.Println(string(byteData))
	//反序列化
	var jsonStr = "{\"name\":\"sss\",\"age\":22}"
	var i1 Info
	err := json.Unmarshal([]byte(jsonStr), &i1) //将json字符串转换为字节数组
	fmt.Println(err)
	fmt.Println(i1)
}
```

### 自定义数据类型

#### 自定义类型
```go
package main

import "fmt"

type codeType int
type payType int

func (c codeType) String() string {
	return fmt.Sprintf("this is %d", c)
}

var successCode codeType = 0
var payCode payType = 0

func print(code codeType) {
	fmt.Println(code)
}

func main() {
	print(successCode)
}

```
```go
package main

import "fmt"

type Code int

const (
	SuccessCode    Code = 0
	ValidCode      Code = 7 // 校验失败的错误
	ServiceErrCode Code = 8 // 服务错误
)

func (c Code) GetMsg() string {
	return "成功"
}

func (c Code) String() string {
	fmt.Printf("code:%d", c)
	return "success"
}

func main() {
	fmt.Println(SuccessCode.GetMsg())
	var i int
	fmt.Println(int(SuccessCode) == i) // 必须要转成原始类型才能判断
	fmt.Println(SuccessCode)		   // 实现了String()方法，自动调用
}
```

#### 类型别名
```go
package main

import "fmt"

type AliasCode = int
type MyCode int

const (
	SuccessCode      MyCode    = 0
	SuccessAliasCode AliasCode = 0
)

// MyCodeMethod 自定义类型可以绑定自定义方法
func (m MyCode) MyCodeMethod() {

}

// MyAliasCodeMethod 类型别名 不可以绑定方法
//func (m AliasCode) MyAliasCodeMethod() {
//
//}

func main() {
	// 类型别名，打印它的类型还是原始类型
	fmt.Printf("%T %T \n", SuccessCode, SuccessAliasCode) // main.MyCode int
	// 可以直接和原始类型比较
	var i int
	fmt.Println(SuccessAliasCode == i)
	fmt.Println(int(SuccessCode) == i) // 必须转换之后才能和原始类型比较
}
```

<hr>

### 接口

```go
package main

import "fmt"

// Animal 定义一个animal的接口，它有唱，跳，rap的方法
type Animal interface {
	sing()
	jump()
	rap()
}

// Chicken 需要全部实现这些接口
type Chicken struct {
	Name string
}

func (c Chicken) sing() {
	fmt.Println("chicken 唱")
}

func (c Chicken) jump() {
	fmt.Println("chicken 跳")
}
func (c Chicken) rap() {
	fmt.Println("chicken rap")
}

// 全部实现完之后，chicken就不再是一只普通的鸡了

func main() {
	var animal Animal

	animal = Chicken{"ik"}

	animal.sing()
	animal.jump()
	animal.rap()

}
```
```go
package main

import "fmt"

// Animal 定义一个animal的接口，它有唱，跳，rap的方法
type Animal interface {
	sing()
	jump()
	rap()
}

// Chicken 需要全部实现这些接口
type Chicken struct {
	Name string
}

func (c Chicken) sing() {
	fmt.Println("chicken 唱")
}

func (c Chicken) jump() {
	fmt.Println("chicken 跳")
}
func (c Chicken) rap() {
	fmt.Println("chicken rap")
}

// Cat 需要全部实现这些接口
type Cat struct {
	Name string
}

func (c Cat) sing() {
	fmt.Println("cat 唱")
}

func (c Cat) jump() {
	fmt.Println("cat 跳")
}
func (c Cat) rap() {
	fmt.Println("cat rap")
}

func sing(obj Animal) {
	obj.sing()
}

// 全部实现完之后，chicken就不再是一只普通的鸡了

func main() {
	chicken := Chicken{"ik"}
	cat := Cat{"阿狸"}
	sing(chicken)
	sing(cat)
}
``` 

#### 类型断言
```go
package main

import "fmt"

// Animal 定义一个animal的接口，它有唱，跳，rap的方法
type Animal interface {
	sing()
	jump()
	rap()
}

// Chicken 需要全部实现这些接口
type Chicken struct {
	Name string
}

func (c Chicken) sing() {
	fmt.Println("chicken 唱")
}

func (c Chicken) jump() {
	fmt.Println("chicken 跳")
}
func (c Chicken) rap() {
	fmt.Println("chicken rap")
}

// Cat 需要全部实现这些接口
type Cat struct {
	Name string
}

func (c Cat) sing() {
	fmt.Println("cat 唱")
}

func (c Cat) jump() {
	fmt.Println("cat 跳")
}
func (c Cat) rap() {
	fmt.Println("cat rap")
}

func sing(obj Animal) {
	switch obj.(type) {
	case Chicken:
		chicken, ok := obj.(Chicken) //将接口类型的值 obj转换为具体的类型 Chicken。
		fmt.Println("Chicken值：", chicken.Name, "是否是对应类型：", ok)
	case Cat:
		cat := obj.(Cat)
		fmt.Println("Cat值：", cat.Name)
	}
	obj.sing()
}

func main() {
	chicken := Chicken{"ik"}
	cat := Cat{"阿狸"}
	sing(chicken)
	sing(cat)
}
```

### 协程和channel

#### 协程
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var wait sync.WaitGroup

func shop(name string) {
	fmt.Printf("%s开始购物\n", name)
	time.Sleep(1 * time.Second)
	fmt.Printf("%s结束购物\n", name)
	wait.Done()
}

func main() {
	now := time.Now()
	//购物接力
	//shop("张三")
	//shop("李四")
	//shop("王五")

	wait.Add(3)
	//主线程结束，协程跟着结束
	go shop("张三")
	go shop("李四")
	go shop("王五")

	wait.Wait()

	//time.Sleep(1 * time.Second)

	fmt.Println("购买完成", time.Since(now))
}
```

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func sing(wait *sync.WaitGroup) {
	fmt.Println("唱歌")
	time.Sleep(1 * time.Second)
	fmt.Println("唱歌结束")
	wait.Done()
}

func main() {
	var wait sync.WaitGroup
	now := time.Now()

	wait.Add(2)
	go sing(&wait)
	go sing(&wait)
	wait.Wait()

	fmt.Println("购买完成", time.Since(now))
}
```
WaitGroup 让主线程等待协程

#### Channel
```go
package main

import "fmt"

func main() {
	ch := make(chan int)

	// 发送数据并关闭 channel
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
		close(ch) // 必须关闭！
	}()

	// 使用 range 遍历 channel
	for val := range ch {
		fmt.Println(val) // 输出 0, 1, 2
	}

	//for {
	//	val, ok := <-ch // 手动接收并检查 channel 状态
	//	if !ok {
	//		break
	//	}
	//	fmt.Println(val)
	//}
}
```
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var moneyChan = make(chan int, 3)

func pay(name string, money int, wait *sync.WaitGroup) {
	fmt.Printf("%s开始购物\n", name)
	time.Sleep(1 * time.Second)
	fmt.Printf("%s结束购物\n", name)
	moneyChan <- money
	wait.Done()
}

func main() {
	now := time.Now()
	var wait sync.WaitGroup

	wait.Add(3)

	go pay("张三", 2, &wait)
	go pay("李四", 3, &wait)
	go pay("王五", 5, &wait)

	go func() {
		wait.Wait()         // 阻塞当前协程，直到所有协程完成
		close(moneyChan)
	}()

	//for {
	//	money, ok := <-moneyChan
	//	fmt.Println(money, ok)
	//	if !ok {        // ok == true：表示成功从 channel 读取到数据。ok == false：表示 channel 已关闭且无剩余数据。
	//		break
	//	}
	//}

	//for val := range moneyChan {
	//	fmt.Println(val)
	//}

	var moneyList []int
	for money := range moneyChan {
		moneyList = append(moneyList, money)
	}

	fmt.Println("购买完成", time.Since(now))
	fmt.Println(moneyList)
}
```
全局moneyChan

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func pay(name string, money int, moneyChan chan<- int, wait *sync.WaitGroup) {
	fmt.Println(name, "唱歌")
	time.Sleep(1 * time.Second)
	fmt.Println(name, "唱歌结束")
	moneyChan <- money
	wait.Done()
}

func main() {
	var moneyChan = make(chan int)
	var wait sync.WaitGroup
	now := time.Now()

	wait.Add(2)
	go pay("张三", 10, moneyChan, &wait)
	go pay("王五", 20, moneyChan, &wait)

	go func() {
		defer close(moneyChan)
		wait.Wait()
	}()

	var moneyList []int
	for money := range moneyChan {
		moneyList = append(moneyList, money)
	}

	fmt.Println("购买完成", time.Since(now))
}
```
局部moneyChan

chan int 双向 channel 读写均可

chan<- int 只写 channel 函数仅发送数据

<-chan int 只读 channel 函数仅接收数据

结合 doneChan + WaitGroup的情况
如果既要等待多个协程，又要支持超时，可以结合使用
```go
func main() {
var wg sync.WaitGroup
done := make(chan struct{})

wg.Add(3)
go worker(&wg, 1)
go worker(&wg, 2)
go worker(&wg, 3)

// 启动一个协程等待所有任务完成
go func() {
wg.Wait()
close(done) // 全部完成后关闭 channel
}()

select {
case <-done:
fmt.Println("所有任务完成")
case <-time.After(2 * time.Second):
fmt.Println("超时")
}
}
```

#### select
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var moneyChan = make(chan int, 3)
var nameChan = make(chan string, 3)

func pay(name string, money int, wait *sync.WaitGroup) {
	fmt.Printf("%s开始购物\n", name)
	time.Sleep(1 * time.Second)
	fmt.Printf("%s结束购物\n", name)
	moneyChan <- money
	nameChan <- name
	wait.Done()
}

func main() {
	now := time.Now()
	var wait sync.WaitGroup

	wait.Add(3)

	go pay("张三", 2, &wait)
	go pay("李四", 3, &wait)
	go pay("王五", 5, &wait)

	go func() {
		wait.Wait() // 阻塞当前协程，直到所有协程完成
		close(moneyChan)
		close(nameChan)
	}()

	var moneyList []int
	var nameList []string

	for {
		select {
		case money, ok := <-moneyChan:
			if !ok {
				moneyChan = nil // 关闭后设为 nil
			} else {
				moneyList = append(moneyList, money)
			}
		case name, ok := <-nameChan:
			if !ok {
				nameChan = nil // 关闭后设为 nil
			} else {
				nameList = append(nameList, name)
			}
		}
		// 退出条件：两个 channel 均已关闭
		if moneyChan == nil && nameChan == nil {
			break
		}
	}

	fmt.Println("购买完成", time.Since(now))
	fmt.Println(moneyList)
	fmt.Println(nameList)
}
```
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var moneyChan = make(chan int, 3)
var nameChan = make(chan string, 3)
var doneChan = make(chan struct{})

func pay(name string, money int, wait *sync.WaitGroup) {
	fmt.Printf("%s开始购物\n", name)
	time.Sleep(1 * time.Second)
	fmt.Printf("%s结束购物\n", name)
	moneyChan <- money
	nameChan <- name
	wait.Done()
}

func main() {
	now := time.Now()
	var wait sync.WaitGroup

	wait.Add(3)

	go pay("张三", 2, &wait)
	go pay("李四", 3, &wait)
	go pay("王五", 5, &wait)

	go func() {
		defer close(moneyChan)
		defer close(nameChan)
		defer close(doneChan)
		wait.Wait() // 阻塞当前协程，直到所有协程完成
	}()

	var moneyList []int
	var nameList []string

	var event = func() {
		for {
			select {
			case money := <-moneyChan:
				moneyList = append(moneyList, money)
			case name := <-nameChan:
				nameList = append(nameList, name)
			case <-doneChan:
				return
			}
		}
	}
	event()
	fmt.Println("购买完成", time.Since(now))
	fmt.Println(moneyList)
	fmt.Println(nameList)
}
```

#### 超时处理
```go
package main

import (
	"fmt"
	"time"
)

var doneChan = make(chan struct{})

func event() {
	fmt.Println("执行开始")
	time.Sleep(time.Second)
	fmt.Println("执行完毕")
	close(doneChan)
}

func main() {
	go event()
	select {
	case <-doneChan:
		fmt.Println("协程执行完毕")
	case <-time.After(2 * time.Second):
		fmt.Println("超时")
		return
	}
}
```

<hr>

### 线程安全与sync.Map

#### 线程安全
```go
package main

import (
	"fmt"
	"sync"
)

func add(num *int, wait *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		*num++
	}
	wait.Done()
}

func sub(num *int, wait *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		*num--
	}
	wait.Done()
}

func main() {
	var wait sync.WaitGroup
	wait.Add(2)
	var num = 2

	go add(&num, &wait)
	go sub(&num, &wait)

	wait.Wait()
	fmt.Println(num)
}
```
CPU的调度方法为抢占式执行，随机调度，结果完全无法预测

#### 同步锁
```go
package main

import (
	"fmt"
	"sync"
)

func add(num *int, wait *sync.WaitGroup, lock *sync.Mutex) {
	lock.Lock()
	for i := 0; i < 10; i++ {
		*num++
	}
	lock.Unlock()
	wait.Done()
}

func sub(num *int, wait *sync.WaitGroup, lock *sync.Mutex) {
	lock.Lock()
	for i := 0; i < 10; i++ {
		*num--
	}
	lock.Unlock()
	wait.Done()
}

func main() {
	var wait sync.WaitGroup
	var num = 2
	var lock sync.Mutex
	wait.Add(2)

	go add(&num, &wait, &lock)
	go sub(&num, &wait, &lock)

	wait.Wait()
	fmt.Println(num)
}
```

#### 线程安全下的map
```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var maps = map[int]string{}
	var wait sync.WaitGroup
	var lock sync.Mutex

	wait.Add(2)
	go func() {
		for {
			lock.Lock()
			maps[1] = "sss"
			lock.Unlock()
		}
		wait.Done()
	}()
	go func() {
		for {
			lock.Lock()
			fmt.Println(maps[1])
			lock.Unlock()
		}
		wait.Done()
	}()

	wait.Wait()
	select {}
}
```
```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var maps = sync.Map{}

	go func() {
		for {
			maps.Store(1, "sss")
		}
	}()
	go func() {
		for {
			val, ok := maps.Load(1)
			fmt.Println(val, ok)
		}
	}()
	select {}

}
```

### 异常处理

#### 常见的异常处理

##### 向上抛
```go
package main

import (
	"errors"
	"fmt"
)

func div(a, b int) (res int, err error) {
	if b == 0 {
		err = errors.New("division by zero")
		return
	}
	res = a / b
	return
}

func server(a, b int) (res int, err error) {
	res, err = div(a, b)
	if err != nil {
		return      //把 err 原样返回，这就是“向上抛”
	}
	res++
	return
}

func main() {
	var a, b = 2, 1
	res, err := server(a, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
```

##### 中断程序
```go
package main

import (
	"fmt"
	"os"
)

func init() {
	_, err := os.ReadFile("111")
	if err != nil {
		//log.Fatalln(err)
		panic(err)
	}
}

func main() {
	fmt.Println("main")
}
```

##### 恢复程序
```go
package main

import (
	"fmt"
	"runtime/debug"
)

func read() {
	var list = []int{1, 2}
	fmt.Println(list[2])
}

func main1() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	read()
}

func main() {
	main1()
	fmt.Println("正常逻辑")
}
```

### 泛型

#### 泛型函数
```go
package main

import "fmt"

type Number interface {		// 泛型约束（类型集合）
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

func add[T int | uint](a, b T) T {
	return a + b
}

func sub[T Number](a, b T) T {
	return a - b
}

func main() {
	var a, b int = 1, 2
	fmt.Println(add(a, b))
	var c, d uint = 1, 2
	fmt.Println(add(c, d))
	var e, f int = 1, 2
	fmt.Println(sub(e, f))
}
```

#### 泛型结构体
```go
package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

type User struct {
	Name string
	Age  int
}

func main() {
	res := Response{
		Code: 200,
		Msg:  "ok",
		Data: User{
			Name: "sss",
			Age:  22,
		},
	}
	byteData, _ := json.Marshal(res)
	fmt.Println(string(byteData))       //{"code":200,"msg":"ok","data":{"Name":"sss","Age":22}}
}
```
```go
package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var jsonStr = "{\"code\":200,\"msg\":\"ok\",\"data\":{\"Name\":\"sss\",\"Age\":22}}"
	var res1 Response
	json.Unmarshal([]byte(jsonStr), &res1)
	fmt.Println(res1) // {200 ok map[Age:22 Name:sss]}
}
```
```go
package main

import (
	"encoding/json"
	"fmt"
)

type Response[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var jsonStr = "{\"code\":200,\"msg\":\"ok\",\"data\":{\"Name\":\"sss\",\"Age\":22}}"
	var res1 Response[User]
	err := json.Unmarshal([]byte(jsonStr), &res1)
	if err != nil {
		return
	}
	fmt.Println(res1)
	fmt.Println(res1.Data)
	fmt.Println(res1.Data.Name)
	//{200 ok {sss 22}}
	//{sss 22}
	//sss
}
```

#### 泛型切片
```go
package main

import "fmt"

type mySlice[T any] []T

func main() {
	var slice1 mySlice[string]
	slice1 = append(slice1, "hello")
	fmt.Println(slice1)

	var slice2 mySlice[int]
	slice2 = append(slice2, 123)
	fmt.Println(slice2)
}
```

#### 泛型map
```go
package main

import "fmt"

type myMap[K int | string, V any] map[K]V

func main() {
	var map1 = make(myMap[int, string])
	map1[1] = "sss"
	fmt.Println(map1)
}
```

### 文件操作

#### 文件读取

##### 一次性读取
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	byteData, err := os.ReadFile("D:\\Golang_Code\\Golang\\hello.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(byteData))
}
```

##### 分片读
```go
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("hello.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	for {
		var byteData = make([]byte, 12)
		n, err := file.Read(byteData)
		if err == io.EOF {
			break
		}
		fmt.Println(string(byteData), n)
	}
}
```

##### 按行读
```go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("hello.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf := bufio.NewReader(file)
	for {
		line, _, err := buf.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(string(line))
	}
}
```

#### 指定分隔符
```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("hello.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf := bufio.NewScanner(file)
	buf.Split(bufio.ScanWords)
	var index int
	for buf.Scan() {
		index++
		fmt.Println(index, buf.Text())
	}
}
```

#### 文件写入

##### 手动写
```go
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.OpenFile("x.txt", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 写入数据
	_, err = file.Write([]byte("Hello World"))
	if err != nil {
		fmt.Println("写入失败:", err)
		return
	}

	// ✅ 关键：将文件指针移动到开头
	_, err = file.Seek(0, 0)
	if err != nil {
		fmt.Println("定位失败:", err)
		return
	}

	// 读取全部内容
	byteData, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(byteData)) // 输出: Hello World
}
```

##### 一次性写
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.WriteFile("xx.txt", []byte("hello world"), 0666)
	fmt.Println(err)
}
```

##### 文件复制
```go
package main

import (
	"io"
	"os"
)

func main() {
	rFile, err := os.Open("girl.png")
	if err != nil {
		panic(err)
	}
	defer rFile.Close()

	wFile, err := os.OpenFile("girl1.png", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	defer wFile.Close()
	io.Copy(wFile, rFile)
}
```

##### 目录操作
```go
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	dir, err := os.ReadDir("leetcode01")
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range dir {
		// 打印名称和是否为目录
		fmt.Print(entry.Name(), "\t", entry.IsDir())

		// 获取详细信息（如文件大小）
		info, err := entry.Info()
		if err != nil {
			fmt.Printf("\t无法获取信息: %v\n", err)
			continue
		}

		// 只有非目录才有“大小”概念，目录的 Size() 是其元数据大小，通常不是内容大小
		if !entry.IsDir() {
			fmt.Printf("\t%d bytes\n", info.Size())
		} else {
			fmt.Printf("\t<DIR>\n")
		}
	}
}
```

### 单元测试
```go
package main

func Add(a, b int) int {
	return a + b
}
```
```go
package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 2)
	if result != 4 {
		t.Errorf("Add(2, 2) = %d; want %d", result, 4)
		return
	}
	t.Logf("测试成功")
}
```

#### 子测试
```go
package main

import "testing"

func TestAdd(t *testing.T) {
	t.Run("add1", func(t *testing.T) {
		if Add(1, -1) != 0 {
			t.Error("add1 failed")
			return
		}
	})
	t.Run("add2", func(t *testing.T) {
		if Add(1, -2) != -1 {
			t.Error("add2 failed")
			return
		}
	})
}
```
```go
package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	cases := []struct {
		name           string
		a, b, expected int
	}{
		{"Add1", 1, 2, 3},
		{"Add2", 1, -1, 0},
		{"Add3", 1, -2, -1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := Add(c.a, c.b)
			if result != c.expected {
				t.Errorf("Expected %d, but got %d", c.expected, result)
			}
		})
	}
}
```

#### TestMain函数
```go
package main

import (
	"fmt"
	"os"
	"testing"
)

func setup() {
	fmt.Println("测试前")
}

func teardown() {
	fmt.Println("测试后")
}

func TestAdd(t *testing.T) {
	fmt.Println("测试中")
	cases := []struct {
		name           string
		a, b, expected int
	}{
		{"Add1", 1, 2, 3},
		{"Add2", 1, -1, 0},
		{"Add3", 1, -2, -1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := Add(c.a, c.b)
			if result != c.expected {
				t.Errorf("Expected %d, but got %d", c.expected, result)
			}
		})
	}
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
```

### 反射

#### 类型判断
```go
package main

import (
	"fmt"
	"reflect"
)

func refType(obj any) {
	typeObj := reflect.TypeOf(obj)
	fmt.Println(typeObj, typeObj.Kind())
	switch typeObj.Kind() {
	case reflect.Int:
		fmt.Println("int")
	case reflect.String:
		fmt.Println("string")
	case reflect.Struct:
		fmt.Println("struct")
	}
}

func main() {
	refType(11)
	refType("hello")
	refType(struct {
		Name string
	}{})
}
```
typeObj 是 reflect.Type 类型的对象，它表示一个类型的完整信息。
类型名称：Person
所属包路径（如果有）
字段信息（字段名、类型、tag等）
方法列表
嵌套结构等元信息
typeObj.Kind() 返回的是这个类型的底层基础种类（kind）。

#### 通过反射获取值
```go
package main

import (
	"fmt"
	"reflect"
)

func refValue(obj any) {
	v := reflect.ValueOf(obj)
	fmt.Println(v, v.Type(), v.Kind())
	switch v.Kind() {
	case reflect.Int:
		fmt.Println("int", v.Int())
	case reflect.String:
		fmt.Println("string", v.String())
	case reflect.Struct:
		fmt.Println("struct")
	}
}

func main() {
	refValue(11)
	refValue("hello")
	refValue(struct {
		Name string
	}{
		Name: "world",
	})
}
```
reflect.ValueOf(obj) 返回一个 reflect.Value 类型的对象。它封装了 obj 的实际值，而不仅仅是类型信息。
读取值
修改值（在满足条件时）
调用方法
访问字段（如果是结构体）
构建切片、map 等

v.Type()和typeObj := reflect.TypeOf(obj)等价，完整的类型信息。

<a href="https://sm.ms/image/EgymqTjo6bYJQVc" target="_blank"><img src="https://s2.loli.net/2025/09/22/EgymqTjo6bYJQVc.png" ></a>

#### 通过反射修改值
```go
package main

import (
	"fmt"
	"reflect"
)

func refsetValue(obj any, value any) {
	v1 := reflect.ValueOf(obj)
	v2 := reflect.ValueOf(value)
	if v1.Elem().Kind() != v2.Kind() {
		return
	}
	switch v1.Elem().Kind() {
	//解引用
	case reflect.String:
		v1.Elem().SetString(v2.String()) //v1.Elem().SetString(value.(string)) //断言
	case reflect.Int:
		v1.Elem().SetInt(v2.Int())
	}

}

func main() {
	name := "x"
	age := 18
	refsetValue(&name, "xxx")
	refsetValue(&age, 20)
	fmt.Println(name)
	fmt.Println(age)
}
```

#### 结构体反射
```go
package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string `json:"name"`
	Age   int
	IsMan bool
}

func ParseJson(obj any) {
	v := reflect.ValueOf(obj)
	t := reflect.TypeOf(obj)
	for i := 0; i < v.NumField(); i++ {
		tf := t.Field(i)
		jsonTag := tf.Tag.Get("json")
		//fmt.Println(tf.Name, tf.Tag)
		if jsonTag == "" {
			jsonTag = tf.Name
		}
		fmt.Println(jsonTag, v.Field(i))
	}
}

func main() {
	s := Student{"XXX", 20, true}
	ParseJson(s)
}
```

#### 修改结构体中某些值
```go
package main

import (
	"fmt"
	"reflect"
	"strings"
)

type User struct {
	Name1 string `big:"name1"`
	Name2 string
}

func SetStruct(obj any) {
	v := reflect.ValueOf(obj).Elem()
	t := reflect.TypeOf(obj).Elem()
	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i)
		tf := t.Field(i)
		jsonTag := tf.Tag.Get("big")
		if jsonTag == "" {
			continue
		}
		value.SetString(strings.ToUpper(value.String()))
	}
}

func main() {
	user := User{"name1", "name2"}
	SetStruct(&user)
	fmt.Println(user)
}
```

#### 
```go
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func (User) Hello(name string) {
	fmt.Println("Hello World", name)
}

func Call(obj any) {
	v := reflect.ValueOf(obj).Elem()
	t := reflect.TypeOf(obj).Elem()
	for i := 0; i < t.NumMethod(); i++ {
		methodValue := v.Method(i)
		methodType := t.Method(i)
		if methodType.Name != "Hello" {
			continue
		}
		methodValue.Call([]reflect.Value{reflect.ValueOf("xxxxx")})
	}
}

func main() {
	user := User{"xxx", 21}
	Call(&user)
}
```
reflect.Value 是 reflect 包中的一个结构体类型（实际是 struct），它封装了一个任意类型的值，并提供了一系列方法来：
查看这个值的类型
获取或修改它的字段
调用它的方法
设置它的值（如果可设置）
判断它的种类（是 struct？int？string？）

获取 reflect.Value	reflect.ValueOf(x)
获取类型信息	v.Type(), v.Kind()
获取值内容	v.Int(), v.String(), v.Interface()
修改值	v.SetXxx()（必须可设置）
结构体字段	v.Field(i) 或 v.FieldByName("Name")
方法调用	v.Method(i).Call(args)
指针解引用	v.Elem()

#### orm
```go
package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type ClassModel struct {
	Name string `jimoruchen-orm:"name"`
	Id   int    `jimoruchen-orm:"id"`
}

func Find(obj any, query ...any) (sql string, err error) {
	//判断obj必须是结构体
	t := reflect.TypeOf(obj)
	if t.Kind() != reflect.Struct {
		err = errors.New("obj must be a struct")
		return
	}

	var where string
	//验证问号个数是否是对等的
	if len(query) > 0 {
		q := query[0]
		qs, ok := q.(string)
		if !ok {
			err = errors.New("query must be string")
			return
		}
		if strings.Count(qs, "?") + 1 != len(query) {
			err = errors.New("查询参数不匹配")
		}
		//拼接where
		for _, a := range query[1:] {
			switch s := a.(type) {
			case string:
				qs = strings.Replace(qs, "?", fmt.Sprintf("'%s'", s), 1)
			case int:
				qs = strings.Replace(qs, "?", fmt.Sprintf("%d", s), 1)
			}
		}
		where = "where " + qs
	}
	//拼接所有的有jimoruchen-orm的字段
	var columns []string
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		ormTag := field.Tag.Get("jimoruchen-orm")
		if ormTag == "" {
			continue
		}
		columns = append(columns, ormTag)
	}

	//算表名，小写机构体的名字加下划线s
	name := strings.ToLower(t.Name()) + "s"
	
	sql = fmt.Sprintf("select %s from %s %s", strings.Join(columns, ", "), name, where)
	return
}

func main() {
	sql, err := Find(ClassModel{}, "name = ?", "三年一班")
	fmt.Println(sql, err)
	sql, err = Find(ClassModel{}, "id = ? and name = ?", 1, "三年一班")
	fmt.Println(sql, err)
	sql, err = Find(ClassModel{})
	fmt.Println(sql, err)
}
```

### 网络编程

#### TCP

##### 服务端
```go
package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("listening on", listener.Addr())
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("客户端地址是：" + conn.RemoteAddr().String())
		conn.Write([]byte("hello world"))
		time.Sleep(2 * time.Second)
		conn.Close()
	}
}
```

#### 客户端
````go
package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		var buf = make([]byte, 1024)
		n, err := conn.Read(buf)
		if err == io.EOF {
			break
		}
		fmt.Println(string(buf[:n]))
	}
}
````

#### 服务端
```go
package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("listening on", addr)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("客户端：" + conn.RemoteAddr().String() + "连接...")
		for {
			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			if err == io.EOF {
				fmt.Println("客户端" + conn.RemoteAddr().String() + "断开连接...")
				break
			}
			fmt.Println(string(buf[:n]))
		}
	}
}
```

#### 客户端
````go
package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		var text string
		fmt.Printf("请输入内容：")
		fmt.Scanln(&text)
		if text == "exit" {
			conn.Close()
			break
		}
		conn.Write([]byte(text))
	}
}
````

### HTTP

#### 服务端
```go
package main

import (
	"fmt"
	"net/http"
)

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.URL.Path, request.UserAgent())
	writer.Write([]byte("Hello World!"))
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("web server listening at addr: http://127.0.0.1:8080")
	http.ListenAndServe("127.0.0.1:8080", nil)
}
```

#### 客户端
```go
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	response, err := http.Get("http://127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	byteData, _ := io.ReadAll(response.Body)
	fmt.Println(string(byteData))

}
```

#### 服务端
```go
package main

import (
	"fmt"
	"net/http"
	"os"
)

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.URL.Path, request.UserAgent())
	byteData, err := os.ReadFile("http_study/index.html")
	if err != nil {
		writer.Write([]byte("文件不存在"))
		return
	}
	writer.Write(byteData)
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("web server listening at addr: http://127.0.0.1:8080")
	http.ListenAndServe("127.0.0.1:8080", nil)
}
```

#### 客户端
```go
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	response, err := http.Get("http://127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	byteData, _ := io.ReadAll(response.Body)
	fmt.Println(string(byteData))

}
``






