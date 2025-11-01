package main

import "fmt"

func longestConsecutive(nums []int) int {
	ans := 0
	maps := make(map[int]bool)
	for _, num := range nums {
		maps[num] = true
	}
	for k, _ := range maps {
		if maps[k-1] {
			continue
		}
		y := k + 1
		for maps[y] {
			y++
		}
		ans = max(ans, y-k)
	}
	return ans
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums))
}
