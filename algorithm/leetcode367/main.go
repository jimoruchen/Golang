package main

func isPerfectSquare(num int) bool {
	left, right := 0, num
	for left < right {
		mid := (right-left)/2 + left
		if num <= mid*mid {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if left*left == num {
		return true
	}
	return false
}
