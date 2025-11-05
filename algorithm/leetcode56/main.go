package main

import (
	"fmt"
	"slices"
)

func merge(intervals [][]int) [][]int {
	ans := make([][]int, 0)
	// sort.Slice(intervals, func(i, j int) bool {
	//     return intervals[i][0] < intervals[j][0]
	// })
	slices.SortFunc(intervals, func(p, q []int) int {
		return p[0] - q[0]
	})
	for _, nums := range intervals {
		length := len(ans)
		if length > 0 && nums[0] <= ans[length-1][1] {
			ans[length-1][1] = max(nums[1], ans[length-1][1])
		} else {
			ans = append(ans, nums)
		}
	}
	return ans
}

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	intervals = merge(intervals)
	for _, interval := range intervals {
		fmt.Println(interval)
	}
}
