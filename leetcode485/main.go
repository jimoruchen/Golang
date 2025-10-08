package main

import "fmt"

func main() {
	var nums = []int{1, 1, 1, 2, 3, 1, 1}
	var n int
	n = findMaxConsecutiveOnes(nums)
	fmt.Println(n)
}
