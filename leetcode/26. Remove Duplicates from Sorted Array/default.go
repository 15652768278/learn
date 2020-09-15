package main

func removeDuplicates(nums []int) int {
	lastValue, lastCursor := 0, 0
	for k, v := range nums {
		if k == 0 {
			lastValue = v
			lastCursor++
			continue
		}
		if lastValue != v {
			lastValue = v
			nums[lastCursor] = v
			lastCursor++
		}
	}
	return lastCursor
}
