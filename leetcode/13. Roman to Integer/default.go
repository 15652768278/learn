package main

func romanToInt(s string) int {
	ans := 0
	var last int
	for i := range s {
		num := toNumber(s[i])
		ans += num
		if i > 0 && num > last {
			ans -= last * 2
		}
		last = num
	}
	return ans
}

func toNumber(b uint8) int {
	switch b {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}
	return 0
}
