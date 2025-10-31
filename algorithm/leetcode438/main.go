package main

func findAnagrams(s string, p string) []int {
	var res []int
	if len(p) > len(s) {
		return res
	}
	sCount := make([]int, 26)
	pCount := make([]int, 26)
	for _, ch := range p {
		pCount[ch-'a']++
	}
	left := 0
	for right := 0; right < len(s); right++ {
		sCount[s[right]-'a']++
		if right-left+1 > len(p) {
			sCount[s[left]-'a']--
			left++
		}
		if right-left+1 == len(p) && equal(sCount, pCount) {
			res = append(res, left)
		}
	}
	return res
}

func equal(s, p []int) bool {
	for i := 0; i < 26; i++ {
		if s[i] != p[i] {
			return false
		}
	}
	return true
}
