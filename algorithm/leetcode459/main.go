package main

import "strings"

func repeatedSubstringPattern(s string) bool {
	for i := 1; i <= len(s)/2; i++ {
		if len(s)%i != 0 {
			continue
		}
		flag := true
		sub := s[:i]
		for j := i; j < len(s); j += i {
			if sub != s[j:j+i] {
				flag = false
				break
			}
		}
		if flag {
			return true
		}
	}
	return false
}

func repeatedSubstringPattern1(s string) bool {
	ss := s + s
	// 去掉首尾字符
	ss = ss[1 : len(ss)-1]
	// 判断 s 是否在 ss 中
	return strings.Contains(ss, s)
}
