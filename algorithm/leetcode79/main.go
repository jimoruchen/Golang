package main

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	if m == 0 || n == 0 || len(word) == 0 {
		return false
	}
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	var backtracking func(index, i, j int) bool
	backtracking = func(index, i, j int) bool {
		if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] || board[i][j] != word[index] {
			return false
		}
		if index == len(word)-1 {
			return true
		}
		visited[i][j] = true
		found := backtracking(index+1, i-1, j) || backtracking(index+1, i, j-1) || backtracking(index+1, i+1, j) || backtracking(index+1, i, j+1)
		if !found {
			visited[i][j] = false
		}
		return found
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if backtracking(0, i, j) {
				return true
			}
		}
	}
	return false
}
