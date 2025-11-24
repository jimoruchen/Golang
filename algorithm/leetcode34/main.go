package main

import (
	"fmt"
	"sort"
)

func searchRange(nums []int, target int) []int {
	left := findFirst(nums, target)
	if left == -1 {
		return []int{-1, -1}
	}
	right := findLast(nums, target)
	return []int{left, right}
}

// 找第一个等于 target 的索引
func findFirst(nums []int, target int) int {
	i, j := 0, len(nums)-1
	res := -1
	for i <= j {
		mid := i + (j-i)/2
		if nums[mid] == target {
			res = mid   // 记录可能的左边界
			j = mid - 1 // 继续向左找
		} else if nums[mid] < target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return res
}

// 找最后一个等于 target 的索引
func findLast(nums []int, target int) int {
	i, j := 0, len(nums)-1
	res := -1
	for i <= j {
		mid := i + (j-i)/2
		if nums[mid] == target {
			res = mid   // 记录可能的右边界
			i = mid + 1 // 继续向右找
		} else if nums[mid] < target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return res
}

func searchRange1(nums []int, target int) []int {
	leftmost := sort.SearchInts(nums, target)
	if leftmost == len(nums) || nums[leftmost] != target {
		return []int{-1, -1}
	}
	rightmost := sort.SearchInts(nums, target+1) - 1
	return []int{leftmost, rightmost}
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Println(searchRange1(nums, target))
}
