package main

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([]int, len(prices))
	dp[0] = prices[0]
	ans := 0
	for i := 1; i < len(prices); i++ {
		ans = max(ans, prices[i]-dp[i-1])
		dp[i] = min(dp[i-1], prices[i])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
