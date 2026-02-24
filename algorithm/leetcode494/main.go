package main

func findTargetSumWays(nums []int, target int) int {
	// P N  P+N=S P-N=T P=(S+T)/2
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if (sum+target)%2 != 0 || sum < abs(target) {
		return 0
	}
	ans := (sum + target) / 2
	dp := make([]int, ans+1)
	dp[0] = 1
	for _, num := range nums {
		for j := ans; j >= num; j-- {
			dp[j] = dp[j] + dp[j-num]
		}
	}
	return dp[ans]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
