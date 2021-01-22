package main

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	rowLen, colLen := len(matrix), len(matrix[0])
	var ans []int
	for offset := 0; offset*2 < rowLen && offset*2 < colLen; offset++ {
		ans = append(ans, oneCycle(matrix, rowLen, colLen, offset)...)
	}
	return ans
}

func oneCycle(matrix [][]int, rowLen, colLen, offset int) []int {
	endX, endY := colLen-1-offset, rowLen-1-offset
	var ans []int
	for i := offset; i <= endX; i++ {
		ans = append(ans, matrix[offset][i])
	}
	if offset < endY {
		for i := offset + 1; i <= endY; i++ {
			ans = append(ans, matrix[i][endX])
		}
	}
	if offset < endX && offset < endY {
		for i := endX - 1; i >= offset; i-- {
			ans = append(ans, matrix[endY][i])
		}
	}
	if offset < endX && offset < endY-1 {
		for i := endY - 1; i > offset; i-- {
			ans = append(ans, matrix[i][offset])
		}
	}
	return ans
}
