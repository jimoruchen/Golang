package main

func spiralOrder(matrix [][]int) []int {
	var ans []int
	m, n := len(matrix), len(matrix[0])
	if m == 0 || n == 0 {
		return nil
	}
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	dir := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	row, col, index := 0, 0, 0
	for i := 0; i < m*n; i++ {
		ans = append(ans, matrix[row][col])
		visited[row][col] = true
		nextRow, nextCol := row+dir[index][0], col+dir[index][1]
		if nextRow < 0 || nextRow >= m || nextCol < 0 || nextCol >= n || visited[nextRow][nextCol] {
			index = (index + 1) % 4
		}
		row += dir[index][0]
		col += dir[index][1]
	}
	return ans
}

func spiralOrder1(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	var ans []int
	m, n := len(matrix), len(matrix[0])
	top, bottom, left, right := 0, m-1, 0, n-1
	for len(ans) < m*n {
		// 1. 从左到右 (Top Row)
		for col := left; col <= right && len(ans) < m*n; col++ {
			ans = append(ans, matrix[top][col])
		}
		top++
		// 2. 从上到下 (Right Col)
		for row := top; row <= bottom && len(ans) < m*n; row++ {
			ans = append(ans, matrix[row][right])
		}
		right--
		// 3. 从右到左 (Bottom Row)
		for col := right; col >= left && len(ans) < m*n; col-- {
			ans = append(ans, matrix[bottom][col])
		}
		bottom--
		// 4. 从下到上 (Left Col)
		for row := bottom; row >= top && len(ans) < m*n; row-- {
			ans = append(ans, matrix[row][left])
		}
		left++
	}
	return ans
}
