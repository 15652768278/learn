package main

func reverse(x int) int {
	var ans int
	for x != 0 {
		if temp := int32(ans); (temp*10)/10 != temp {
			return 0
		}
		ans = ans*10 + x%10
		x /= 10
	}
	return ans
}
