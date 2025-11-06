package main

import "fmt"

func maxArea(height []int) int {
	length := len(height)
	ans := 0
	left := 0
	right := length - 1
	for left != right {
		tmp := 0
		if height[left] < height[right] {
			tmp = height[left] * (right - left)
			ans = max(ans, tmp)
			left++
		} else {
			tmp = height[right] * (right - left)
			ans = max(ans, tmp)
			right--
		}
	}
	return ans
}

func main() {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(nums))
}
