package main

func canJump(nums []int) bool {
	maxLength := 0
	for i := 0; i < len(nums); i++ {
		if maxLength >= len(nums)-1 {
			return true
		}
		if i > maxLength {
			return false
		}
		maxLength = max(maxLength, i+nums[i])
	}
	return true
}
