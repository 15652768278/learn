package main

func maxSubArray(nums []int) int {
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > ans {
			ans = nums[i]
		}
	}
	return ans
}
