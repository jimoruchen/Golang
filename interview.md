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

## 