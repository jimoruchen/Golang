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

	println(strings.Count("111000111000", "1"))       //统计字符出现个数		6
	println(strings.TrimSpace("   jimo ruchen   "))   //去除空格，中间的不行		jimo ruchen
	println(strings.HasSuffix("xxx.jpg", ".jpg"))     //判断后缀		true
	println(strings.HasPrefix("user-xxx", "user"))    //判断前缀		true
	println(strings.Contains("user-xxx", "x"))        //判断包含		true
	println(strings.Replace("user-xxx", "x", "y", 3)) //替换3个	user-yyy
	println(strings.ReplaceAll("user-xxx", "x", "y")) //全部替换		user-yyy
	println(strings.ToLower("AbC"))                   //abc
	println(strings.ToUpper("AbC"))                   //ABC
	println(strings.ToTitle("hello world!"))          //HELLO WORLD!

	text := "Hello, World"
	if strings.Contains(text, "World") {
		fmt.Println("Found 'World'") //Found 'World'
	}
	parts := strings.Split(text, ", ")
	fmt.Println(parts) // 	[Hello World]
}
