package main

func intersect(nums1 []int, nums2 []int) []int {
	cache := make(map[int]int)
	for _, num := range nums1 {
		cache[num]++
	}
	var ans []int
	for _, num := range nums2 {
		count, ok := cache[num]
		if ok && count > 0 {
			ans = append(ans, num)
			count--
		}
		cache[num] = count
	}
	return ans
}
