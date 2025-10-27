package main

import "fmt"

func insertSort(nums []int) {
	length := len(nums)
	for i := 1; i < length; i++ {
		key := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > key {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = key
	}
}

func shellSort(nums []int) {
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
	var nums = []int{2, 1, 4, 3, 5}
	insertSort(nums)
	fmt.Println(nums)
	shellSort(nums)
	fmt.Println(nums)
}
