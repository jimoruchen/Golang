package main

import "fmt"

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right { //最后一轮没命中刚好left+1就是要插入的位置
		mid := (right-left)/2 + left
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return left
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 5
	fmt.Println(searchInsert(nums, target))
}
