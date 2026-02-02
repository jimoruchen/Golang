package main

func removeInvalidParentheses(s string) []string {
	left, right := 0, 0 //要去除的 ( 和 )
	for _, char := range s {
		if char == '(' {
			left++
		} else if char == ')' {
			if left > 0 {
				left--
			} else {
				right++
			}
		}
	}
	var path []byte
	var result []string
	visited := make(map[string]bool)
	var backtracking func(index, leftCount, rightCount, leftRem, rightRem int)
	backtracking = func(index, leftCount, rightCount, leftRem, rightRem int) {
		if index == len(s) {
			if leftRem == 0 && rightRem == 0 {
				str := string(path)
				if !visited[str] {
					visited[str] = true
					result = append(result, str)
				}
			}
			return
		}
		char := s[index]
		if char == '(' && leftRem > 0 {
			backtracking(index+1, leftCount, rightCount, leftRem-1, rightRem)
		}
		if char == ')' && rightRem > 0 {
			backtracking(index+1, leftCount, rightCount, leftRem, rightRem-1)
		}
		path = append(path, char)
		if char == '(' {
			backtracking(index+1, leftCount+1, rightCount, leftRem, rightRem)
		} else if char == ')' {
			if rightCount < leftCount {
				backtracking(index+1, leftCount, rightCount+1, leftRem, rightRem)
			}
		} else {
			backtracking(index+1, leftCount, rightCount, leftRem, rightRem)
		}
		path = path[:len(path)-1]
	}
	backtracking(0, 0, 0, left, right)
	return result
}
