package main

func longestPalindrome(s string) string {
	var ans string
	length := 0
	expand := func(left, right int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right-left+1 > length {
				length = right - left + 1
				ans = s[left : right+1]
			}
			left--
			right++
		}
	}
	for i := 0; i < len(s); i++ {
		expand(i, i)
		expand(i, i+1)
	}
	return ans
}

func longestPalindrome1(s string) string {
	var ans string
	n := 0
	var expand = func(left, right int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right-left+1 > n {
				ans = s[left : right+1]
				n = right - left + 1
			}
			left--
			right++
		}
	}
	for i := 0; i < len(s); i++ {
		expand(i, i)
		expand(i, i+1)
	}
	return ans
}
