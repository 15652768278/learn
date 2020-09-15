package main

func reverseStr(s string, k int) string {
	byteSlice := []byte(s)
	for i := 0; i < len(byteSlice); i += 2 * k {
		left := i
		right := min(i+k, len(byteSlice)) - 1
		for left < right {
			byteSlice[left], byteSlice[right] = byteSlice[right], byteSlice[left]
			left++
			right--
		}
	}
	return string(byteSlice)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
