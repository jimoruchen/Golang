package main

import "fmt"

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

func twoSum2(nums []int, target int) []int {
	var maps = map[int]int{}
	for i, num := range nums {
		if _, ok := maps[target-num]; ok {
			return []int{i, maps[target-num]}
		}
		maps[num] = i
	}
	return nil
}

func main() {
	var nums = []int{2, 7, 11, 15}
	var target = 9
	ans := twoSum2(nums, target)
	fmt.Println(ans)
}
