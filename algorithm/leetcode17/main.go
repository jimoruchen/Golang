package main

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	mapping := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var path []byte
	var result []string
	var backtracking func(int)
	backtracking = func(index int) {
		if len(digits) == index {
			result = append(result, string(path))
			return
		}
		for _, digit := range mapping[digits[index]-'0'] {
			path = append(path, byte(digit))
			backtracking(index + 1)
			path = path[:len(path)-1]
		}
	}
	backtracking(0)
	return result
}
