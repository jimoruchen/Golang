package main

func twoSum(nums []int, target int) []int {
	mapAns := map[int]int{}
	for i, num := range nums {
		if value, ok := mapAns[target-num]; ok {
			return []int{value, i}
		} else {
			mapAns[num] = i
		}
	}
	return nil
}
