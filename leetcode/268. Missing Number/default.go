package main

func missingNumber(nums []int) int {
	ans := len(nums)
	for i, num := range nums {
		ans = ans + i - num
	}
	return ans
}
