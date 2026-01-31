package main

func generateParenthesis(n int) []string {
	var path []byte
	var result []string
	var backtracking func(int, int)
	backtracking = func(left, right int) {
		if len(path) == 2*n {
			result = append(result, string(path))
			return
		}
		if left < n {
			path = append(path, '(')
			backtracking(left+1, right)
			path = path[:len(path)-1]
		}
		if right < left {
			path = append(path, ')')
			backtracking(left, right+1)
			path = path[:len(path)-1]
		}
	}
	backtracking(0, 0)
	return result
}
