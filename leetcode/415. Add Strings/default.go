package main

import "strconv"

func addStrings(num1 string, num2 string) string {
	var ans, carry = "", 0
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || carry > 0; i, j = i-1, j-1 {
		var x, y int
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		res := x + y + carry
		ans = strconv.Itoa(res%10) + ans
		carry = res / 10
	}
	return ans
}
