package main

func partitionLabels(s string) []int {
	var ans []int
	last := [26]int{}
	for i, c := range s {
		last[c-'a'] = i
	}
	start, end := 0, 0
	for i, c := range s {
		if last[c-'a'] > end {
			end = last[c-'a']
		}
		if i == end {
			ans = append(ans, end-start+1)
			start = end + 1
		}
	}
	return ans
}
