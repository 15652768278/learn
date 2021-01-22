package main

func hammingWeight(num uint32) int {
	var ans uint32
	for ; num > 0; num = num >> 1 {
		ans += num % 2
	}
	return int(ans)
}
