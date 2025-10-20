package main

import "fmt"

func ShellSort(nums []int) {
	length := len(nums)
	for gap := length / 2; gap > 0; gap /= 2 {
		for i := gap; i < length; i++ {
			key := nums[i]
			j := i
			for j >= gap && nums[j-gap] > key {
				nums[j] = nums[j-gap]
				j -= gap
			}
			nums[j] = key
		}
	}
}

func main() {
	nums := []int{2, 1, 4, 3, 5}
	ShellSort(nums)
	fmt.Println(nums)
}
