package main

func fib(N int) int {
	if N < 2 {
		return N
	}
	dp := make([]int, N+1)
	dp[0], dp[1] = 0, 1
	for i := 2; i <= N; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[N]
}
