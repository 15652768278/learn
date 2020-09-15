package main

func isAnagram(s string, t string) bool {
	cache := make(map[int32]int)
	for _, c := range s {
		cache[c]++
	}
	for _, c := range t {
		cache[c]--
	}
	for _, val := range cache {
		if val != 0 {
			return false
		}
	}
	return true
}
