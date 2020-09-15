package main

func validPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] == s[right] {
			left++
			right--
			continue
		}
		flag1, flag2 := true, true
		for i, j := left, right-1; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				flag1 = false
				break
			}
		}
		for i, j := left+1, right; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				flag2 = false
				break
			}
		}
		return flag1 || flag2
	}
	return true
}
