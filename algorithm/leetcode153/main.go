package main

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right { //当 left == right 时已找到唯一候选，无需继续；避免死循环
		mid := left + (right-left)/2
		if nums[mid] > nums[len(nums)-1] {
			left = mid + 1
		} else {
			right = mid //mid 可能是最小值，必须保留
		}
	}
	return nums[left]
}
