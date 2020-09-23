package main

func hammingWeight(num uint32) int {
	ans := 0
	for ; num > 0; num /= 2 {
		if num&1 == 1 {
			ans++
		}
	}
	return ans
}
