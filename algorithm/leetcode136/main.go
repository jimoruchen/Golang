package main

import "fmt"

func singleNumber(nums []int) int {
	ans := 0
	maps := make(map[int]int)
	for _, num := range nums {
		maps[num]++
	}
	for k, v := range maps {
		if v == 1 {
			ans = k
		}
	}
	return ans
}

func singleNumber1(nums []int) int {
	ans := 0
	for _, num := range nums {
		ans ^= num
	}
	return ans
}

func main() {
	nums := []int{1, 2, 2, 3, 3}
	fmt.Println(singleNumber1(nums))
}
