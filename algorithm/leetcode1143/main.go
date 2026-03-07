package main

func longestCommonSubsequence(text1 string, text2 string) int {
	ans := 0
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
			ans = max(ans, dp[i][j])
		}
	}
	return ans
}

//      a b c d e
//    a 1 1 1 1 1
//    c 1 1 2 2 2
//    e 1 1 2 2 3
