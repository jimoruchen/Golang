package main

func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	if n == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	l, r := 0, n-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] >= nums[0] {
			if target >= nums[0] && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[n-1] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}

func search1(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if target == nums[mid] {
			return mid
		}
		if nums[mid] >= nums[left] {
			if target < nums[mid] && target >= nums[left] {
				right = mid
			} else {
				left = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[right-1] {
				left = mid + 1
			} else {
				right = mid
			}
		}
	}
	return -1
}
