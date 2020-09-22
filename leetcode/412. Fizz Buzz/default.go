package main

import "fmt"

func fizzBuzz(n int) []string {
	var ans []string
	for i := 1; i <= n; i++ {
		elem := fmt.Sprintf("%d", i)
		if i%3 == 0 && i%5 == 0 {
			elem = "FizzBuzz"
		} else if i%3 == 0 {
			elem = "Fizz"
		} else if i%5 == 0 {
			elem = "Buzz"
		}
		ans = append(ans, elem)
	}
	return ans
}
