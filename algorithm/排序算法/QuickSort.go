package main

import (
	"fmt"
	"math/rand"
)

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	pivot := partition1(nums, left, right)
	quickSort(nums, left, pivot-1)
	quickSort(nums, pivot+1, right)
}

func partition(nums []int, left, right int) int {
	i, j := left, right
	for i < j {
		for i < j && nums[j] >= nums[left] {
			j--
		}
		for i < j && nums[i] <= nums[left] {
			i++
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i], nums[left] = nums[left], nums[i]
	return i
}

func partition1(nums []int, left, right int) int {
	index := left + rand.Intn(right-left+1)
	nums[left], nums[index] = nums[index], nums[left]
	i, j := left, right
	for i < j {
		for i < j && nums[j] >= nums[left] {
			j--
		}
		for i < j && nums[i] <= nums[left] {
			i++
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i], nums[left] = nums[left], nums[i]
	return i
}

func main() {
	nums := []int{2, 1, 4, 3, 5}
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
