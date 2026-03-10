package main

func sortedSquares(nums []int) []int {
	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] * nums[i]
	}
	left, right := 0, len(nums)-1
	for pos := len(nums) - 1; pos >= 0; pos-- {
		if nums[left] > nums[right] {
			ans[pos] = nums[left]
			left++
		} else {
			ans[pos] = nums[right]
			right--
		}
	}
	return ans
}
