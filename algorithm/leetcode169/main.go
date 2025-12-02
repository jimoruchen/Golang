package main

import (
	"fmt"
	"sort"
)

func majorityElement(nums []int) int {
	ans := 0
	maps := make(map[int]int)
	for _, num := range nums {
		maps[num]++
	}
	for k, v := range maps {
		if v > len(nums)/2 {
			ans = k
		}
	}
	return ans
}

func majorityElement1(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func main() {
	nums := []int{1, 2, 2}
	fmt.Println(majorityElement(nums))
}
