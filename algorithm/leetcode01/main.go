package main

import "fmt"

func main() {
	var nums = []int{2, 7, 11, 15}
	var target = 9
	ans := twoSum1(nums, target)
	fmt.Println(ans)
}
