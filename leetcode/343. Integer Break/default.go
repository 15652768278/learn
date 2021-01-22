package main

import "math"

func integerBreak(n int) int {
	if n < 4 {
		return n - 1
	}
	a, b := n/3, n%3
	if b == 0 {
		return int(math.Pow(3, float64(a)))
	}
	if b == 1 {
		return int(math.Pow(3, float64(a-1)) * 4)
	}
	if b == 2 {
		return int(math.Pow(3, float64(a)) * 2)
	}
	return 0
}
