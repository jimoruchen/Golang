package main

import (
	"fmt"
	"sort"
)

func searchMatrix(matrix [][]int, target int) bool {
	var tmp []int
	for _, row := range matrix {
		tmp = append(tmp, row...)
	}
	left := 0
	right := len(tmp) - 1
	for left <= right {
		mid := (right-left)/2 + left
		if tmp[mid] == target {
			return true
		} else if tmp[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

func searchMatrix1(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	left, right := 0, m*n-1

	for left <= right {
		mid := (left + right) / 2
		row, col := mid/n, mid%n
		val := matrix[row][col]
		if val == target {
			return true
		} else if val > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
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
