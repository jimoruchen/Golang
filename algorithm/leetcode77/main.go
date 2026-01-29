package main

import "fmt"

func combine(n int, k int) [][]int {
	var path []int
	var result [][]int
	backtracking(n, k, 1, path, &result)
	return result
}

func backtracking(n, k, startIndex int, path []int, result *[][]int) {
	if len(path) == k {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*result = append(*result, tmp)
		return
	}
	for i := startIndex; i <= n; i++ {
		path = append(path, i)
		backtracking(n, k, i+1, path, result)
		path = path[:len(path)-1]
	}
}

func combine1(n int, k int) [][]int {
	var result [][]int
	var path []int
	var backtracking func(int)
	backtracking = func(startIndex int) {
		if len(path) == k {
			tmp := make([]int, len(path))
			copy(tmp, path)
			result = append(result, tmp)
			return
		}
		for i := startIndex; i <= n; i++ {
			path = append(path, i)
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtracking(1)
	return result
}

func main() {
	n := 4
	k := 2
	fmt.Println(combine(n, k))
}
