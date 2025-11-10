package main

import "fmt"

func subarraySum(nums []int, k int) int {
	count, sum := 0, 0
	prefixMap := map[int]int{0: 1}
	for _, num := range nums {
		sum += num
		if tmp, ok := prefixMap[sum-k]; ok {
			count += tmp
		}
		prefixMap[sum]++
	}
	return count
}

func main() {
	var nums = []int{1, 2, 3}
	k := 3
	fmt.Println(subarraySum(nums, k))
}
