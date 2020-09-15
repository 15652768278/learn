package main

import (
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	return strings.Join(r(n), "")
}

func r(n int) []string {
	if n == 1 {
		return []string{"1"}
	}
	if n == 2 {
		return []string{"1", "1"}
	}
	strSlice := r(n - 1)
	var ans []string
	var count = 1
	for i := 1; i < len(strSlice); i++ {
		if strSlice[i] != strSlice[i-1] {
			ans = append(ans, strconv.Itoa(count), strSlice[i-1])
			count = 1
		} else {
			count++
		}
		if i+1 == len(strSlice) {
			ans = append(ans, strconv.Itoa(count), strSlice[i])
		}
	}
	return ans
}
