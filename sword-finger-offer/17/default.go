package main

import (
	"fmt"
	"math"
)

func printNumbers(n int) []int {
	if n == 0 {
		return nil
	}
	end := int(math.Pow(10, float64(n)))
	ans := make([]int, end-1)
	for i := 1; i < end; i++ {
		ans[i-1] = i
	}
	return ans
}

func printNumbers2(n int) []string {
	if n == 0 {
		return nil
	}
	arr := make([]string, 10)
	for i := 0; i < 10; i++ {
		arr[i] = fmt.Sprintf("%v", i)
	}
	for i := 1; i < n; i++ {
		arr = handle(arr)
	}
	return arr[1:]
}

func dealZero(str string) string {
	if len(str) == 0 {
		return str
	}
	if str[1] == '0' {
		return dealZero(str[1:])
	}
	return str
}

func handle(arr []string) []string {
	var ans []string
	for i := 0; i < 10; i++ {
		for _, item := range arr {
			num := fmt.Sprintf("%d%s", i, item)
			ans = append(ans, dealZero(num))
		}
	}
	return ans
}
