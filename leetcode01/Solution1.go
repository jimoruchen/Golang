package main

func twoSum1(nums []int, target int) []int {
	var Mymap = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if val, ok := Mymap[target-nums[i]]; ok {
			return []int{val, i}
		} else {
			Mymap[nums[i]] = i
		}
	}
	return nil
}
