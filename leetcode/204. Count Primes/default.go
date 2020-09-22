package main

func countPrimes(n int) int {
	ans := 0
	cache := make([]bool, n+1)
	for i := 2; i < n; i++ {
		if cache[i] {
			continue
		}
		ans++
		for j := i * 2; j < n; j += i {
			cache[j] = true
		}
	}
	return ans
}
