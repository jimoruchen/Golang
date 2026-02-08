package main

func jump(nums []int) int {
	maxLength := 0
	right := 0
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		maxLength = max(maxLength, i+nums[i])
		if i == right {
			right = maxLength
			count++
		}
	}
	return count
}
