package main

import "fmt"

func permute(nums []int) [][]int {
	var path []int
	var result [][]int
	used := make([]bool, len(nums))
	var backtracking func()
	backtracking = func() {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			result = append(result, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			backtracking()
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	backtracking()
	return result
}

func permute1(nums []int) [][]int {
	var resulit [][]int
	var path []int
	visited := make([]bool, len(nums))
	var backtracking func()
	backtracking = func() {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			resulit = append(resulit, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if visited[i] {
				continue
			}
			path = append(path, nums[i])
			visited[i] = true
			backtracking()
			visited[i] = false
			path = path[:len(path)-1]
		}
	}
	backtracking()
	return resulit
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(permute(nums))
}
