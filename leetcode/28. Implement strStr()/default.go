package main

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	var i, j int
	var index = -1
	for i < len(haystack) && j < len(needle) {
		if haystack[i] == needle[j] {
			if index == -1 {
				index = i
			}
			if j == len(needle)-1 {
				return index
			}
			j++
		} else if index != -1 {
			i = index
			index = -1
			j = 0
		}
		i++
	}
	return -1
}
