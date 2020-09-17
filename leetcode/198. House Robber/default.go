package main

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// f0: 未偷窃状态最大金额
	// f1: 偷窃状态最大金额
	f0, f1 := 0, nums[0]
	for i := 1; i < len(nums); i++ {
		f0, f1 = max(f0, f1), max(f0+nums[i], f1)
	}
	return max(f0, f1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
