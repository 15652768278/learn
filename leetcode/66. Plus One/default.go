package main

func plusOne(digits []int) []int {
	addon := 1
	for i := len(digits) - 1; i >= 0; i-- {
		sum := digits[i] + addon
		digits[i] = sum % 10
		addon = sum / 10
	}
	var ans []int
	if addon == 1 {
		ans = make([]int, 1)
		ans[0] = 1
		ans = append(ans, digits...)
	} else {
		ans = digits
	}
	return ans
}
