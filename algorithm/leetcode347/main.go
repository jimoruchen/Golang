package main

import "sort"

// func topKFrequent(nums []int, k int) []int {
//     maps := make(map[int]int)
//     for _, v := range nums {
//         maps[v]++
//     }
//     var keys []int
//     for k := range maps {
//         keys = append(keys, k)
//     }
//     sort.Slice(keys, func(i, j int) bool {
//         return maps[keys[i]] > maps[keys[j]]
//     })
//     return keys[:k]
// }

func topKFrequent(nums []int, k int) []int {
	var ans []int
	maps := make(map[int]int)
	for _, v := range nums {
		maps[v]++
	}
	type kv struct {
		k int
		v int
	}
	var KV []kv
	for k, v := range maps {
		KV = append(KV, kv{k, v})
	}
	sort.Slice(KV, func(i, j int) bool {
		return KV[i].v > KV[j].v
	})
	for i := 0; i < k; i++ {
		ans = append(ans, KV[i].k)
	}
	return ans
}
