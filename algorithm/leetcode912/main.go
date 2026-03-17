package main

import (
	"math/rand"
	"slices"
)

func sortArray(nums []int) []int {
	if slices.IsSorted(nums) {
		return nums
	}
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	pivot := partition(nums, left, right)
	quickSort(nums, pivot+1, right)
	quickSort(nums, left, pivot-1)
}

func partition(nums []int, left, right int) int {
	index := rand.Intn(right-left+1) + left
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
	nums[left], nums[i] = nums[i], nums[left]
	return i
}
