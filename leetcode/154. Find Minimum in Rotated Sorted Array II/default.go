package main

func findMin(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	var min = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			return nums[i]
		}
	}
	return min
}