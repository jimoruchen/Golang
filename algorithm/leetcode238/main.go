package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	prefix := make([]int, len(nums))
	suffix := make([]int, len(nums))
	ans := make([]int, len(nums))
	tmp := 1
	for i := 0; i < len(nums); i++ {
		prefix[i] = tmp
		tmp = tmp * nums[i]
	}
	tmp = 1
	for i := len(nums) - 1; i >= 0; i-- {
		suffix[i] = tmp
		tmp = tmp * nums[i]
	}
	for i := 0; i < len(nums); i++ {
		ans[i] = prefix[i] * suffix[i]
	}
	return ans
}

func main() {
	nums := []int{1, 2, 3, 4}
	ans := productExceptSelf(nums)
	fmt.Println(ans)
}
