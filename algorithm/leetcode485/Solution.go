package main

func findMaxConsecutiveOnes(nums []int) int {
	var ans = 0
	var tmp = 0
	var length = len(nums)
	for i := 0; i < length; i++ {
		if nums[i] == 1 {
			tmp++
			if tmp > ans {
				ans = tmp
			}
		} else {
			tmp = 0
		}
	}
	return ans
}
