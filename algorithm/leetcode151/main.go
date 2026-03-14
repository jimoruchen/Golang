package main

import "strings"

func reverseWords(s string) string {
	str := strings.Fields(s)
	left := 0
	right := len(str) - 1
	for left < right {
		str[left], str[right] = str[right], str[left]
		left++
		right--
	}
	var sb strings.Builder
	for i, st := range str {
		sb.WriteString(st)
		if i < len(str)-1 {
			sb.WriteString(" ")
		}
	}
	return sb.String()
}
