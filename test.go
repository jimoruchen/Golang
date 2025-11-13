package main

import "fmt"

func main() {
	type Person struct {
		Name string
	}
	p := &Person{
		Name: "John",
	}
	fmt.Println(p.Name)
	fmt.Println((*p).Name)
}
