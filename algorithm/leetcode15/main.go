package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var ans [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left+1] == nums[left] {
					left++
				}
				for left < right && nums[right-1] == nums[right] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return ans
}

func threeSum1(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var ans [][]int
	for left := 0; left < n-2; left++ {
		if left > 0 && nums[left] == nums[left-1] {
			continue
		}
		j := left + 1
		right := n - 1
		for j < right {
			if nums[left]+nums[j]+nums[right] == 0 {
				ans = append(ans, []int{nums[left], nums[j], nums[right]})
				for j < right && nums[j] == nums[j+1] {
					j++
				}
				for j < right && nums[right] == nums[right-1] {
					right--
				}
				j++
				right--
			} else if nums[left]+nums[j]+nums[right] > 0 {
				right--
			} else {
				j++
			}
		}
	}
	return ans
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	ans := threeSum(nums)
	fmt.Println(ans)
}
