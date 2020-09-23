package main

func getRow(rowIndex int) []int {
	return help(rowIndex, []int{1})
}

func help(n int, lastRow []int) []int {
	if n == 0 {
		return lastRow
	}
	ans := []int{1}
	for i := 1; i < len(lastRow); i++ {
		ans = append(ans, lastRow[i]+lastRow[i-1])
	}
	ans = append(ans, 1)
	return help(n-1, ans)
}
