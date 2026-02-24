package main

func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, num := range stones {
		sum += num
	}
	target := sum / 2
	dp := make([]int, target+1)
	for _, num := range stones {
		for j := target; j >= num; j-- {
			dp[j] = max(dp[j], dp[j-num]+num)
		}
	}
	return sum - 2*dp[target]
}
