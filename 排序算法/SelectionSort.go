package main

import "fmt"

func SelectionSort(nums []int) {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

func SelectionSort1(nums []int) {
	length := len(nums)
	for i := 0; i < length; i++ {
		minIndex := i
		for j := i + 1; j < length; j++ {
			if nums[minIndex] > nums[j] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
}

func main() {
	nums := []int{2, 1, 4, 3, 5}
	SelectionSort1(nums)
	fmt.Println(nums)
}
