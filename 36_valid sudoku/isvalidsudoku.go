package main

import "fmt"

func main() {
	board := [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	fmt.Println("Output:", isValidSudoku(board))
}

//执行用时 :4 ms, 在所有 Go 提交中击败了90.15%的用户
//内存消耗 :4.3 MB, 在所有 Go 提交中击败了10.11%的用户
func isValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]int, 9)
	cols := make([]map[byte]int, 9)
	squares := make([]map[byte]int, 9)
	for i := range rows {
		rows[i] = make(map[byte]int)
		cols[i] = make(map[byte]int)
		squares[i] = make(map[byte]int)
	}
	var m map[byte]int
	for i := range board {
		for j, n := range board[i] {
			if board[i][j] == '.' {
				continue
			}
			// Row Check
			m = rows[i]
			if _, ok := m[n]; ok {
				return false
			}
			m[n] = 0
			// Column Check
			m = cols[j]
			if _, ok := m[n]; ok {
				return false
			}
			m[n] = 0
			// Square Check
			m = squares[i/3*3+j/3]
			if _, ok := m[n]; ok {
				return false
			}
			m[n] = 0
		}
	}
	return true
}
