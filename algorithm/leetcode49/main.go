package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	maps := make(map[string][]string)
	for _, str := range strs {
		chars := []byte(str)
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})
		sortStr := string(chars)
		maps[sortStr] = append(maps[sortStr], str)
	}
	var ans [][]string
	for _, v := range maps {
		ans = append(ans, v)
	}
	return ans
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	ans := groupAnagrams(strs)
	fmt.Println(ans)
}
