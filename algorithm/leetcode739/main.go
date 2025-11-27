package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	var stack []int
	for i := 0; i < n; i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[prevIndex] = i - prevIndex
		}
		stack = append(stack, i)
	}
	return ans
}

func main() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	ans := dailyTemperatures(temperatures)
	fmt.Println(ans)
}
