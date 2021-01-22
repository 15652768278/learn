package main

func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func climbStairs2(n int) int {
	a, b, sum := 1, 1, 0
	for i := 0; i < n; i++ {
		sum = a + b
		a, b = b, sum
	}
	return a
}
