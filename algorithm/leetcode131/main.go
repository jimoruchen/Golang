package main

func partition(s string) [][]string {
	var path []string
	var res [][]string
	var backtracking func(startIndex int)
	backtracking = func(startIndex int) {
		if startIndex == len(s) {
			tmp := make([]string, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}
		for i := startIndex; i < len(s); i++ {
			str := s[startIndex : i+1]
			if isPalindrome(str) {
				path = append(path, str)
				backtracking(i + 1)
				path = path[:len(path)-1]
			}
		}
	}
	backtracking(0)
	return res
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
