package main

func maxProfit(prices []int) int {
	ans := 0
	for k, v := range prices {
		if k > 0 && v > prices[k-1] {
			ans += v - prices[k-1]
		}
	}
	return ans
}
