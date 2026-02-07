package main

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	count := 0
	for _, cookie := range s {
		if count < len(g) && cookie >= g[count] {
			count++
		}
	}
	return count
}
