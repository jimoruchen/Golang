# Golang

## for循环
<img src="https://s2.loli.net/2025/11/12/1DpRvjgumexFtaJ.png"  alt="">

<hr>

## new() 与 make() 的区别
new(T) 和 make(T,args) 是 Go 语言内建函数，用来分配内存，但适用的类型不同。
new(T) 会为 T 类型的新值分配已置零的内存空间，并返回地址（指针），即类型为 *T的值。换句话说就是，返回一个指针，该指针指向新分配的、类型为 T 的零值。适用于值类型，如数组、结构体等。
make(T,args) 返回初始化之后的 T 类型的值，这个值并不是 T 类型的零值，也不是指针 *T，是经过初始化之后的 T 的引用。make() 只适用于 slice、map 和 channel。

<hr>

## 将一个切片追加到另一个切片上
```go
func main() {
	list1 := []int{1, 2, 3}
	list2 := []int{4, 5, 6}
	list1 = append(list1, list2...)
	fmt.Println(list1)
}
```

<hr>

## var变量声明
```go
var(
    size := 1024
    max_size = size*2
)
```
变量声明的简短模式(:=)
1.必须使用显示初始化；
2.不能提供数据类型，编译器会自动推导；
3.只能在函数内部使用简短模式；

<hr>

## 结构体比较
```go
func main() {
    sn1 := struct {
        age  int
        name string
    }{age: 11, name: "qq"}
    sn2 := struct {
        age  int
        name string
    }{age: 11, name: "qq"}

    if sn1 == sn2 {
        fmt.Println("sn1 == sn2")
    }

    sm1 := struct {
        age int
        m   map[string]string
    }{age: 11, m: map[string]string{"a": "1"}}
    sm2 := struct {
        age int
        m   map[string]string
    }{age: 11, m: map[string]string{"a": "1"}}

    if sm1 == sm2 {
        fmt.Println("sm1 == sm2")
    }
}
```
上面正确，下面错误。
结构体只能比较是否相等，但是不能比较大小。
相同类型的结构体才能够进行比较，结构体是否相同不但与属性类型有关，还与属性顺序相关，sn3 与 sn1 就是不同的结构体；
```go
sn3:= struct {
	name string
	age  int
}{age:11,name:"qq"}
```
如果 struct 的所有成员都可以比较，则该 struct 就可以通过 == 或 != 进行比较是否相等，比较时逐个项进行比较，如果每一项都相等，则两个结构体才相等，否则不相等；
那什么是可比较的呢，常见的有 bool、数值型、字符、指针、数组等，像切片、map、函数等是不能比较的。

<hr>

## 通过指针变量访问成员变量
指针 是一个变量，它存储的是另一个变量的内存地址。
解引用 就是 通过这个地址，去访问或操作那个实际的值。
```go
func main() {
	type Person struct {
		Name string
	}
	p := &Person{
		Name: "John",
	}
	fmt.Println(p.Name)     //编译器会自动解引用指针，等价于 (*p).name
	fmt.Println((*p).Name)
}
```

<hr>

## 类型别名和类型定义
```go
package main

import "fmt"

type MyInt1 int     //新类型
type MyInt2 = int   //类型别名

func main() {
    var i int =0
    var i1 MyInt1 = i   //将 int 类型的变量赋值给 MyInt1 类型的变量，所以不行。
    var i2 MyInt2 = i   // MyInt2 只是 int 的别名，本质上还是 int，可以赋值。
    fmt.Println(i1,i2)
}
```

<hr>

## 拼接字符串
<img src="https://s2.loli.net/2025/11/17/NgxZKqS4ksdzbRv.png"  alt="">

B,D正确，A,C中单引号 ' 只能用于定义单个字符（rune），不能用于字符串。

## strings包
`strings.EqualFold(s, t string) bool`
忽略大小写比较两个字符串（支持 Unicode）。
`strings.EqualFold("Hello", "HELLO") // true
`

`strings.Contains(s, substr string) bool`
判断是否包含子串：
`strings.Contains("hello", "ll") // true`

`strings.HasPrefix(s, prefix string) bool`
是否以某前缀开头：
`strings.HasPrefix("https://example.com", "https") // true`

`strings.HasSuffix(s, suffix=str) bool`
是否以某后缀结尾：
`strings.HasSuffix("test.go", ".go") // true`

`strings.Index(s, substr string) int`
返回子串第一次出现的位置，未找到返回 -1：
`strings.Index("hello", "l") // 2`

`strings.Split(s, sep string) []string`
按分隔符切分：
`strings.Split("a,b,c", ",") // []string{"a", "b", "c"}`

`strings.Fields(s string) []string`
按空白字符（空格、制表符、换行等）切分，自动忽略多余空白：
`strings.Fields("  a  b\tc\n") // ["a", "b", "c"]`

`strings.SplitN(s, sep string, n int) []string`
最多切分成 n 个部分：
`strings.SplitN("a,b,c,d", ",", 3) // ["a", "b", "c,d"]`

`strings.Join(elems []string, sep string) string`
用分隔符连接字符串切片：
`strings.Join([]string{"a", "b", "c"}, "-") // "a-b-c"`

`strings.Builder`
```go
var sb strings.Builder
sb.WriteString("hello")
sb.WriteString(" ")
sb.WriteString("world")
result := sb.String() // "hello world"
```

`strings.TrimSpace(s string) string`
去除首尾所有空白字符（包括 \t, \n, 空格等）：
`strings.TrimSpace(" \t hello \n ") // "hello"`

`strings.Trim(s, cutset string) string`
去除首尾指定字符集中的任意字符：
`strings.Trim("!!hello!!", "!") // "hello"`

`strings.TrimPrefix(s, prefix string) string`
去除前缀（如果存在）：
`strings.TrimPrefix("https://example.com", "https://") // "example.com"`

`strings.TrimSuffix(s, suffix string) string`
去除后缀（如果存在）：
`strings.TrimSuffix("file.txt.bak", ".bak") // "file.txt"`

`strings.Replace(s, old, new string, n int) string`
替换前 n 个 old 为 new；n = -1 表示全部替换：
`strings.Replace("aabbcc", "b", "x", -1) // "aaxxcc"`
`strings.ReplaceAll(s, old, new) // "aaxxcc"`

`strings.Repeat(s string, count int) string`
重复字符串：
`strings.Repeat("ha", 3) // "hahaha"`

`strings.ToUpper(s string) string / strings.ToLower(s string) string`
转大写/小写（仅处理 ASCII）：
`strings.ToUpper("hello") // "HELLO"`

<hr>

## sort.Slice
```go
nums := []int{5, 2, 6, 3, 1}
sort.Slice(nums, func(i, j int) bool {
    return nums[i] < nums[j] // 升序
})
// 结果: [1, 2, 3, 4, 5]
```
```go
sort.Slice(nums, func(i, j int) bool {
    return nums[i] > nums[j] // 降序
})
// 结果: [5, 4, 3, 2, 1]
```
```go
type Person struct {
    Name string
    Age  int
}

people := []Person{
    {"Alice", 25},
    {"Bob",   30},
    {"John",  25},
}
sort.Slice(people, func(i, j int) bool {
return people[i].Age < people[j].Age
})
// 结果: John(25), Alice(25), Bob(30)
```
```go
sort.Slice(people, func(i, j int) bool {
    if people[i].Age != people[j].Age {
        return people[i].Age < people[j].Age // 年龄升序
    }
    return people[i].Name < people[j].Name // 年龄相同时按名字升序
})
```

<hr>

## slices.SortFunc
```go
package main

import (
    "fmt"
    "slices"
)

func main() {
    words := []string{"banana", "apple", "cherry"}
    slices.SortFunc(words, func(a, b string) int {
        return len(a) - len(b) // 按长度升序
    })
    fmt.Println(words) // 输出: [apple banana cherry]
}
```
```go
package main

import (
    "fmt"
    "sort"
    "slices"
)

func main() {
    // 原始数据：表示区间 [start, end]
    intervals := [][]int{{3, 6}, {1, 4}, {2, 8}, {5, 7}}

    // 方法1: 使用 sort.Slice (Go 1.8+)
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    fmt.Println("sort.Slice:", intervals)

    // 重置数据（因为上面已修改原切片）
    intervals = [][]int{{3, 6}, {1, 4}, {2, 8}, {5, 7}}

    // 方法2: 使用 slices.SortFunc (Go 1.21+)
    slices.SortFunc(intervals, func(a, b []int) int {
        return a[0] - b[0]
    })
    fmt.Println("slices.SortFunc:", intervals)
}
```

<hr>
