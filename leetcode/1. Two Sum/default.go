package main

func twoSum(nums []int, target int) []int {
	cache := make(map[int]int)
	for k, v := range nums {
		if last, ok := cache[v]; ok {
			return []int{last, k}
		}
		cache[target-v] = k
	}
	return nil
}
