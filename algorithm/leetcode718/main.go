package main

func findLength(nums1 []int, nums2 []int) int {
	ans := 0
	m := len(nums1)
	n := len(nums2)
	dp := make([][]int, m+1) //dp[i][j]表示以nums1[i-1],nums2[j-1]结尾的最长重复子数组
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 0
	for i := 1; i <= m; i++ {
		dp[i][0] = 0
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = 0
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				ans = max(ans, dp[i][j])
			}
		}
	}
	return ans
}
