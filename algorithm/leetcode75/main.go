package main

import "fmt"

func sortColors(nums []int) {
	maps := make(map[int]int)
	for _, num := range nums {
		maps[num]++
	}
	for i := 0; i < maps[0]; i++ {
		nums[i] = 0
	}
	for i := maps[0]; i < maps[0]+maps[1]; i++ {
		nums[i] = 1
	}
	for i := maps[0] + maps[1]; i < len(nums); i++ {
		nums[i] = 2
	}
}

func sortColors1(nums []int) {
	p0, p1 := 0, 0
	for i, num := range nums {
		nums[i] = 2
		if num <= 1 {
			nums[p1] = 1
			p1++
		}
		if num == 0 {
			nums[p0] = 0
			p0++
		}
	}
}

func main() {
	nums := []int{1, 1, 2, 0, 2, 0}
	sortColors(nums)
	fmt.Println(nums)
}
