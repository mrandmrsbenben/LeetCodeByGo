package main

import "fmt"

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'}}
	word := "SEE"
	// board := [][]byte{{'a', 'a'}}
	// word := "aaa"
	fmt.Println("Output: ", exist(board, word))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :3.5 MB, 在所有 Go 提交中击败了91.55%的用户
func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return false
	}

	var nextstep func(pi, pj, n int) bool
	nextstep = func(pi, pj, n int) bool {
		if n == len(word) {
			return true
		}
		i, j := pi, pj

		if i > 0 && board[i-1][j] == word[n] {
			pi, pj = i-1, j
			board[pi][pj] = ' '
			if !nextstep(pi, pj, n+1) {
				board[pi][pj] = word[n]
			} else {
				return true
			}
		}
		if j > 0 && board[i][j-1] == word[n] {
			pi, pj = i, j-1
			board[pi][pj] = ' '
			if !nextstep(pi, pj, n+1) {
				board[pi][pj] = word[n]
			} else {
				return true
			}
		}
		if i < len(board)-1 && board[i+1][j] == word[n] {
			pi, pj = i+1, j
			board[pi][pj] = ' '
			if !nextstep(pi, pj, n+1) {
				board[pi][pj] = word[n]
			} else {
				return true
			}
		}
		if j < len(board[i])-1 && board[i][j+1] == word[n] {
			pi, pj = i, j+1
			board[pi][pj] = ' '
			if !nextstep(pi, pj, n+1) {
				board[pi][pj] = word[n]
			} else {
				return true
			}
		}
		return false
	}

	for r := range board {
		for c := range board[r] {
			if board[r][c] == word[0] {
				board[r][c] = '0'
				if len(word) == 1 {
					return true
				} else if nextstep(r, c, 1) {
					return true
				} else {
					board[r][c] = word[0]
				}
			}
		}
	}

	return false
}
