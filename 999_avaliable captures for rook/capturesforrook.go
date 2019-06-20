package main

import (
	"fmt"
	"strings"
)

func main() {
	board :=
		[][]byte{{'.', '.', '.', '.', '.', '.', '.', '.'},
			{'.', '.', '.', 'p', '.', '.', '.', '.'},
			{'.', '.', '.', 'B', '.', '.', '.', '.'},
			{'p', 'B', 'p', 'R', 'p', 'B', 'p', '.'},
			{'.', '.', '.', 'p', 'p', '.', '.', '.'},
			{'.', '.', '.', 'B', '.', '.', '.', '.'},
			{'.', '.', '.', 'p', '.', '.', '.', '.'},
			{'.', '.', '.', '.', '.', '.', '.', '.'}}
	fmt.Printf("Captures for Rook: %d\n", numRookCaptures(board))
}

func numRookCaptures(board [][]byte) int {
	num := 0
	R, line := -1, -1
	var p, B int
	for i := range board {
		str := string(board[i])
		if R == -1 {
			R = strings.IndexByte(str, 'R')
			if R != -1 {
				line = i - 1
				p = strings.LastIndexByte(str[0:R], 'p')
				B = strings.LastIndexByte(str[0:R], 'B')
				if (B == -1 || B < p) && p >= 0 {
					num = num + 1
				}
				p = strings.IndexByte(str[R:], 'p')
				B = strings.IndexByte(str[R:], 'B')
				if (B == -1 || B > p) && p >= 0 {
					num = num + 1
				}
			}
		} else {
			if board[i][R] == 'B' {
				break
			} else if board[i][R] == 'p' {
				num = num + 1
				break
			}
		}
	}
	if line > 0 {
		for i := line; i >= 0; i-- {
			if board[i][R] == 'B' {
				break
			} else if board[i][R] == 'p' {
				num = num + 1
				break
			}
		}
	}
	return num
}
