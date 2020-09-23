package main

func reverseBits(num uint32) uint32 {
	var ans uint32
	for bits := uint32(31); num > 0; num, bits = num/2, bits-1 {
		ans = ans | (num%2)<<bits
	}
	return ans
}
