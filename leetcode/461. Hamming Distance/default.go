package main

func hammingDistance(x int, y int) int {
	ans := 0
	for xor := x ^ y; xor > 0; xor /= 2 {
		if xor&1 == 1 {
			ans++
		}
	}
	return ans
}
