package main

func firstUniqChar(s string) int {
	cache := make(map[int32]int)
	for _, c := range s {
		cache[c]++
	}
	for i, c := range s {
		if cache[c] == 1 {
			return i
		}
	}
	return -1
}
