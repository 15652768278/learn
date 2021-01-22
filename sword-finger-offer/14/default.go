package main

import "math"

func cuttingRope(n int) int {
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

func cuttingRope2(n int) int {
	if n < 4 {
		return n - 1
	}
	a, b := n/3, n%3
	if b == 0 {
		return pow3N(a)
	}
	if b == 1 {
		return pow3N(a-1) * 4 % (1e9 + 7)
	}
	if b == 2 {
		return pow3N(a) * 2 % (1e9 + 7)
	}
	return 0
}

func pow3N(n int) int {
	o := 1
	for i := 0; i < n; i++ {
		o = (o * 3) % (1e9 + 7)
	}
	return o
}
