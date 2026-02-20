package main

// func generate(numRows int) [][]int {
//     if numRows == 1 {
//         return [][]int{{1}}
//     }
//     if numRows == 2 {
//         return [][]int{{1}, {1, 1}}
//     }
//     var ans [][]int
//     dp := make([][]int, numRows)
//     for i := range dp {
//         dp[i] = make([]int, numRows)
//     }
//     dp[0][0] = 1
//     dp[1][0] = 1
//     dp[1][1] = 1
//     for i := 2; i < numRows; i++ {
//         dp[i][0] = 1
//         dp[i][i] = 1
//         for j := 1; j <= i - 1; j++ {
//             dp[i][j] = dp[i - 1][j - 1] + dp[i - 1][j]
//         }
//     }
//     for i := 0; i < numRows; i++ {
//         ans = append(ans, dp[i][:i + 1])
//     }
//     return ans
// }

func generate(numRows int) [][]int {
	dp := make([][]int, numRows)
	for i := range dp {
		dp[i] = make([]int, i+1)
		dp[i][0] = 1
		dp[i][i] = 1
		for j := 1; j < i; j++ {
			dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
		}
	}
	return dp
}
