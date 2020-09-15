package main

func groupAnagrams(strs []string) [][]string {
	cache := make(map[[26]int][]string)
	for _, str := range strs {
		key := strArray(str)
		last, _ := cache[key]
		last = append(last, str)
		cache[key] = last
	}
	var ans [][]string
	for _, val := range cache {
		ans = append(ans, val)
	}
	return ans
}

func strArray(s string) [26]int {
	res := [26]int{}
	for _, v := range s {
		res[v-'a']++
	}
	return res
}
