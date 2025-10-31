package main

func moveZeroes(nums []int) {
	var len = len(nums)
	var tmp = 0
	for i := 0; i < len; i++ {
		if nums[i] != 0 {
			nums[tmp] = nums[i]
			tmp++
		}
	}
	for i := tmp; i < len; i++ {
		nums[i] = 0
	}
}

func moveZeroes1(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}
