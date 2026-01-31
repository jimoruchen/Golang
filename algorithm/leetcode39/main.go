package main

func combinationSum(candidates []int, target int) [][]int {
	var path []int
	var result [][]int
	var backtracking func(int, int)
	backtracking = func(startIndex, currentSum int) {
		if currentSum > target {
			return
		}
		if currentSum == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			result = append(result, tmp)
			return
		}
		for i := startIndex; i < len(candidates); i++ {
			path = append(path, candidates[i])
			backtracking(i, currentSum+candidates[i])
			path = path[:len(path)-1]
		}
	}
	backtracking(0, 0)
	return result
}
