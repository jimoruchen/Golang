package main

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else if nums[mid] == nums[right] {
			right--
		} else {
			right = mid
		}
	}
	return nums[left]
}

func findMin1(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else if nums[mid] == nums[right] {
			right--
		} else {
			right = mid
		}
	}
	return nums[left]
}
