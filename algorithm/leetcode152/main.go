package main

func maxProduct(nums []int) int {
	ans := nums[0]
	maxdp := make([]int, len(nums))
	mindp := make([]int, len(nums))
	maxdp[0] = nums[0]
	mindp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		curmax := maxdp
		curmin := mindp
		maxdp[i] = max(curmax[i-1]*nums[i], max(curmin[i-1]*nums[i], nums[i]))
		mindp[i] = min(curmax[i-1]*nums[i], min(curmin[i-1]*nums[i], nums[i]))
		ans = max(ans, maxdp[i])
	}
	return ans
}
