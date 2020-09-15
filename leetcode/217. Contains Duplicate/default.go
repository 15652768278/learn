package main

func containsDuplicate(nums []int) bool {
	cache := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := cache[num]; ok {
			return true
		}
		cache[num] = struct{}{}
	}
	return false
}
