package main

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	left, top := 0, 0
	right, bottom := n-1, n-1
	tmp := 1
	for left <= right && top <= bottom {
		for j := left; j <= right; j++ {
			matrix[top][j] = tmp
			tmp++
		}
		top++
		for i := top; i <= bottom; i++ {
			matrix[i][right] = tmp
			tmp++
		}
		right--
		for j := right; j >= left; j-- {
			matrix[bottom][j] = tmp
			tmp++
		}
		bottom--
		for i := bottom; i >= top; i-- {
			matrix[i][left] = tmp
			tmp++
		}
		left++
	}
	return matrix
}
