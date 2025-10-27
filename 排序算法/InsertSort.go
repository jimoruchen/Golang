package main

import "fmt"

func insertSort(nums []int) {
	n := len(nums)
	for i := 1; i < n; i++ {
		key := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > key {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = key
	}
}

func main() {
	nums := []int{2, 1, 4, 3, 5}
	insertSort(nums)
	fmt.Println(nums)
}
