package main

func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}
	ans := [][]int{
		{1},
	}
	for i := 1; i < numRows; i++ {
		lastRow := ans[len(ans)-1]
		row := []int{1}
		for j := 1; j < len(lastRow); j++ {
			row = append(row, lastRow[j]+lastRow[j-1])
		}
		row = append(row, 1)
		ans = append(ans, row)
	}
	return ans
}
