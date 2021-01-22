package main

func fib(n int) int {
	a, b, sum := 0, 1, 0
	for i := 0; i < n; i++ {
		sum = (a + b) % (1e9 + 7)
		a, b = b, sum
	}
	return a
}

func numWays(n int) int {
	a, b, sum := 1, 1, 0
	for i := 0; i < n; i++ {
		sum = (a + b) % (1e9 + 7)
		a, b = b, sum
	}
	return a
}