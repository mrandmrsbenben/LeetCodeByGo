package main

import "fmt"

func main() {
	board := [][]byte{{'X', '.', '.', 'X'}, {'.', '.', '.', 'X'}, {'.', '.', '.', 'X'}}
	fmt.Println("Output:", countBattleships(board))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2.4 MB, 在所有 Go 提交中击败了100.00%的用户
func countBattleships(board [][]byte) int {
	count := 0
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'X' {
				if i > 0 && j > 0 {
					if board[i-1][j] == '.' && board[i][j-1] == '.' {
						count++
					}
				} else if j > 0 {
					if board[i][j-1] == '.' {
						count++
					}
				} else if i > 0 {
					if board[i-1][j] == '.' {
						count++
					}
				} else {
					count++
				}
			}
		}
	}
	return count
}
