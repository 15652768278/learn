package main

var directGroup = [][]int{
	{-1, 0},
	{0, 1},
	{0, -1},
	{1, 0},
}

func exist(board [][]byte, word string) bool {
	mask := make([][]bool, len(board))
	for i := range mask {
		mask[i] = make([]bool, len(board[i]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if search(board, i, j, 0, word) {
				return true
			}
		}
	}
	return false
}

func search(board [][]byte, i, j, k int, word string) bool {
	if len(word) == k {
		return true
	}
	if i < 0 || j < 0 || i > len(board)-1 || j > len(board[0])-1 {
		return false
	}
	if board[i][j] == word[k] {
		cache := board[i][j]
		board[i][j] = ' '
		var flag bool
		for _, direct := range directGroup {
			flag = flag || search(board, i+direct[0], j+direct[1], k+1, word)
		}
		if flag {
			return true
		}
		board[i][j] = cache
	}
	return false
}
