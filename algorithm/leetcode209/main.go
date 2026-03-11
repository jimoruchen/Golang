package main

func minSubArrayLen(target int, nums []int) int {
	left := 0
	sum := 0
	n := len(nums) + 1
	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		for sum >= target {
			n = min(n, right-left+1)
			sum -= nums[left]
			left++
		}
	}
	if n == len(nums)+1 {
		n = 0
	}
	return n
}
