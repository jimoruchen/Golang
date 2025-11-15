package main

import "fmt"

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	tmp := val
	if len(this.minStack) > 0 {
		tmp = min(this.minStack[len(this.minStack)-1], val)
	}
	this.minStack = append(this.minStack, tmp)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

func main() {
	MinStack := Constructor()
	MinStack.Push(1)
	MinStack.Push(2)
	MinStack.Push(3)
	fmt.Println(MinStack.Top(), MinStack.GetMin())
}
