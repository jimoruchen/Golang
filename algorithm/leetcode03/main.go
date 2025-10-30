package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	maps := make(map[byte]struct{})
	left := 0
	count := 0
	for right := 0; right < len(s); right++ {
		for {
			if _, ok := maps[s[right]]; ok {
				delete(maps, s[left])
				left++
			} else {
				break
			}
		}
		maps[s[right]] = struct{}{}
		count = max(count, right-left+1)
	}
	return count
}

func main() {
	s := "abcabcbb"
	res := lengthOfLongestSubstring(s)
	fmt.Println(res)
}
