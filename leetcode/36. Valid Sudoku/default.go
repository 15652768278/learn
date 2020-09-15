package main

func isValidSudoku(board [][]byte) bool {
	var rowCache, colCache, areaCache [9][9]int
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == '.' {
				continue
			}
			num := board[row][col] - '1'
			// row
			if rowCache[row][num] == 1 {
				return false
			}
			rowCache[row][num] = 1
			// col
			if colCache[col][num] == 1 {
				return false
			}
			colCache[col][num] = 1
			// area
			area := row/3*3 + col/3
			if areaCache[area][num] == 1 {
				return false
			}
			areaCache[area][num] = 1
		}
	}
	return true
}
