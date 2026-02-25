package main

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1) // dp[i],金额为i所需的硬币数
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for _, num := range coins {
		for j := num; j <= amount; j++ {
			dp[j] = min(dp[j], dp[j-num]+1)
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
