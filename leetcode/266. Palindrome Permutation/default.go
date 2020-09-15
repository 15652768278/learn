package main

func canPermutePalindrome(s string) bool {
	cache := make(map[int32]int)
	for _, c := range s {
		cache[c]++
	}
	flag := false
	for _, val := range cache {
		if val%2 == 0 {
			continue
		}
		if flag {
			return false
		}
		flag = true
	}
	return true
}
