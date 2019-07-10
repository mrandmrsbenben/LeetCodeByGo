package main

import "fmt"

func main() {
	board := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0}}
	gameOfLife(board)
	fmt.Println("Output:", board)
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2.2 MB, 在所有 Go 提交中击败了25.00%的用户
func gameOfLife(board [][]int) {
	countLiveCells := func(r, c int) int {
		cnt := 0
		for i := r - 1; i < len(board) && i < r+2; i++ {
			if i < 0 {
				continue
			}
			for j := c - 1; j < len(board[i]) && j < c+2; j++ {
				if j < 0 || (i == r && j == c) {
					continue
				}
				if board[i][j] == 1 {
					cnt++
				}
			}
		}
		return cnt
	}
	buf := make([][]int, len(board))
	for i := range board {
		buf[i] = make([]int, len(board[i]))
		for j := 0; j < len(board[i]); j++ {
			cnt := countLiveCells(i, j)
			if board[i][j] == 1 {
				if cnt < 2 || cnt > 3 {
					buf[i][j] = 0
				} else {
					buf[i][j] = 1
				}
			} else {
				if cnt == 3 {
					buf[i][j] = 1
				}
			}
		}
	}
	copy(board, buf)
}
