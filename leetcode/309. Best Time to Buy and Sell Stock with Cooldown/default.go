package main

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	// f0: 手上持有股票的最大收益
	// f1: 手上不持有股票,并且处于冷冻期中的累计最大收益
	// f2: 手上不持有股票,并且不在冷冻期中的累计最大收益
	f0, f1, f2 := -prices[0], 0, 0
	for i := 1; i < len(prices); i++ {
		f0, f1, f2 = max(f0, f2-prices[i]), f0+prices[i], max(f1, f2)
	}
	return max(f1, f2)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
