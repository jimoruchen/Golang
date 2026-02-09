package main

import "math/rand"

func findKthLargest(nums []int, k int) int {
	return quickSort(nums, 0, len(nums)-1, len(nums)-k)
}

// func quickSort(nums []int, left, right int) {
//     if left >= right {
//         return
//     }
//     pivot := partition(nums, left, right)
//     quickSort(nums, left, pivot - 1)
//     quickSort(nums, pivot + 1, right)
// }

func quickSort(nums []int, left, right, k int) int {
	if left == right {
		return nums[left]
	}
	pivot := partition(nums, left, right)
	if pivot == k {
		return nums[pivot]
	} else if pivot > k {
		return quickSort(nums, left, pivot-1, k)
	} else {
		return quickSort(nums, pivot+1, right, k)
	}
}

func partition(nums []int, left, right int) int {
	i, j := left, right
	index := left + rand.Intn(right-left+1)
	nums[index], nums[i] = nums[i], nums[index]
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
