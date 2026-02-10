package main

func fib(n int) int {
	// maps := make(map[int]int)
	// return fibDp1(n, maps)
	return fibDp3(n)
}

func fibDp(n int, maps map[int]int) int {
	if n <= 1 {
		return n
	}
	if v, ok := maps[n]; ok {
		return v
	}
	maps[n] = fibDp(n-1, maps) + fibDp(n-2, maps)
	return maps[n]
}

func fibDp2(n int) int {
	if n <= 1 {
		return n
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func fibDp3(n int) int {
	if n <= 1 {
		return n
	}
	prev, cur := 0, 1
	for i := 2; i <= n; i++ {
		prev, cur = cur, prev+cur
	}
	return cur
}
