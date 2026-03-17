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

func lengthOfLongestSubstring1(s string) int {
	maps := make(map[byte]bool)
	left := 0
	res := 0
	for right := 0; right < len(s); right++ {
		for maps[s[right]] {
			maps[s[left]] = false
			left++
		}
		maps[s[right]] = true
		res = max(res, right-left+1)
	}
	return res
}

func lengthOfLongestSubstring3(s string) int {
	maps := make(map[byte]bool)
	left, ans := 0, 0
	for right := 0; right < len(s); right++ {
		for maps[s[right]] {
			maps[s[left]] = false
			left++
		}
		maps[s[right]] = true
		ans = max(ans, right-left+1)
	}
	return ans
}

func main() {
	s := "abcabcbb"
	res := lengthOfLongestSubstring(s)
	fmt.Println(res)
}
