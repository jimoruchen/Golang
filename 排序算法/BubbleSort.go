package main

import "fmt"

func bubbleSort(nums []int) {
	length := len(nums)
	for i := 0; i < length; i++ {
		swapped := false
		for j := 0; j < length-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func main() {
	nums := []int{2, 1, 4, 3, 5}
	bubbleSort(nums)
	fmt.Println(nums)
}
