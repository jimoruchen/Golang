package main

import (
	"container/list"
	"fmt"
)

func isValid(s string) bool {
	stack := list.New()
	for _, str := range s {
		if str == '(' || str == '[' || str == '{' {
			stack.PushBack(str)
		} else if str == ')' {
			if stack.Len() == 0 {
				return false
			}
			if stack.Back().Value != '(' {
				return false
			}
			stack.Remove(stack.Back())
		} else if str == ']' {
			if stack.Len() == 0 {
				return false
			}
			if stack.Back().Value != '[' {
				return false
			}
			stack.Remove(stack.Back())
		} else {
			if stack.Len() == 0 {
				return false
			}
			if stack.Back().Value != '{' {
				return false
			}
			stack.Remove(stack.Back())
		}
	}
	if stack.Len() != 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(isValid("()"))
}
