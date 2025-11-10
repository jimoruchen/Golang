package main

import "fmt"

func rotate(nums []int, k int) {
	var tmp []int
	k = k % len(nums)
	for i := len(nums) - k; i < len(nums); i++ {
		tmp = append(tmp, nums[i])
	}
	fmt.Println(tmp)
	for i := 0; i < len(nums); i++ {
		tmp = append(tmp, nums[i])
	}
	copy(nums, tmp)
}

func rotate1(nums []int, k int) {
	tmp := make([]int, len(nums))
	k = k % len(nums)
	for i, num := range nums {
		tmp[(k+i)%len(nums)] = num
	}
	copy(nums, tmp)
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	fmt.Println(nums)
}
