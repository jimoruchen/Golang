package main

func wordBreak(s string, wordDict []string) bool {
	maps := make(map[string]bool)
	for _, str := range wordDict {
		maps[str] = true
	}
	dp := make([]bool, len(s)+1)
	// dp[i] 表示 s[0:i] 能否被拆分
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && maps[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
