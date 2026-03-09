package main

import (
	"fmt"
	"sort"
)

func searchMatrix(matrix [][]int, target int) bool {
	var nums []int
	for _, ma := range matrix {
		nums = append(nums, ma...)
	}
	left, right := 0, len(nums)
	for left < right {
		mid := (right-left)/2 + left
		if target <= nums[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if left < len(nums) && nums[left] == target {
		return true
	}
	return false
}

func searchMatrix1(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	left, right := 0, m*n
	for left < right {
		mid := (right-left)/2 + left
		row := mid / n
		col := mid % n
		value := matrix[row][col]
		if target <= value {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if left < m*n && matrix[left/n][left%n] == target {
		return true
	}
	return false
}

func searchMatrix2(matrix [][]int, target int) bool {
	row := sort.Search(len(matrix), func(i int) bool {
		return matrix[i][0] > target
	}) - 1
	if row < 0 {
		return false
	}
	// col := sort.Search(len(matrix[0]), func(i int) bool {
	//     return matrix[row][i] >= target
	// })
	col := sort.SearchInts(matrix[row], target)
	return col < len(matrix[row]) && target == matrix[row][col]
}

func main() {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}
	fmt.Println(searchMatrix(matrix, 3))
}
