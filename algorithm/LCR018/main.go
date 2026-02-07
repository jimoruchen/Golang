package main

import "strings"

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	s = strings.ToLower(s)
	for left < right {
		for left < right && !isVaild(s[left]) {
			left++
		}
		for left < right && !isVaild(s[right]) {
			right--
		}
		if left < right && s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isVaild(b byte) bool {
	return (b >= '0' && b <= '9') || (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}
