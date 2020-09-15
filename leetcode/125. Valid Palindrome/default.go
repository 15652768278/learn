package main

import (
	"strings"
)

func isPalindrome(s string) bool {
	strSlice := strings.Split(s, "")
	left, right := 0, len(strSlice)-1
	for left < right {
		if !isValid(strSlice[left]) {
			left++
			continue
		}
		if !isValid(strSlice[right]) {
			right--
			continue
		}
		if !strings.EqualFold(strSlice[left], strSlice[right]) {
			return false
		}
		left++
		right--
	}
	return true
}

func isValid(s string) bool {
	char := []byte(strings.ToLower(s))[0]
	if char < '0' || (char > '9' && char < 'a') || char > 'z' {
		return false
	}
	return true
}
