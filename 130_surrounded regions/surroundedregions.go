package main

import "fmt"

func main() {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'}}
	solve(board)
	fmt.Println("Output: ", board)
}

//执行用时 :24 ms, 在所有 Go 提交中击败了87.80%的用户
//内存消耗 :6 MB, 在所有 Go 提交中击败了100.00%的用户
func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	lr, lc := len(board), len(board[0])

	var setFlag func(int, int)
	setFlag = func(i, j int) {
		if i < 0 || j < 0 || i >= lr || j >= lc || board[i][j] != 'O' {
			return
		}
		board[i][j] = 'T'
		setFlag(i-1, j)
		setFlag(i+1, j)
		setFlag(i, j-1)
		setFlag(i, j+1)
	}

	for i := 0; i < lc; i++ {
		if board[0][i] == 'O' {
			setFlag(0, i)
		}
		if board[lr-1][i] == 'O' {
			setFlag(lr-1, i)
		}
	}
	for j := 0; j < lr; j++ {
		if board[j][0] == 'O' {
			setFlag(j, 0)
		}
		if board[j][lc-1] == 'O' {
			setFlag(j, lc-1)
		}
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'T' {
				board[i][j] = 'O'
			}
		}
	}
}

//执行用时 :24 ms, 在所有 Go 提交中击败了87.80%的用户
//内存消耗 :6 MB, 在所有 Go 提交中击败了100.00%的用户
func solveV1(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	lr, lc := len(board), len(board[0])
	flag := make([][]int, lr)
	for i := range board {
		flag[i] = make([]int, lc)
		for j := range flag[i] {
			if board[i][j] == 'X' {
				flag[i][j] = 1
			}
		}
	}

	var setFlag func(int, int)
	setFlag = func(i, j int) {
		if i < 0 || j < 0 || i >= lr || j >= lc || flag[i][j] > 0 {
			return
		}
		flag[i][j] = 2
		setFlag(i-1, j)
		setFlag(i+1, j)
		setFlag(i, j-1)
		setFlag(i, j+1)
	}

	for i := 0; i < lc; i++ {
		if flag[0][i] == 0 {
			setFlag(0, i)
		}
		if flag[lr-1][i] == 0 {
			setFlag(lr-1, i)
		}
	}
	for j := 0; j < lr; j++ {
		if flag[j][0] == 0 {
			setFlag(j, 0)
		}
		if flag[j][lc-1] == 0 {
			setFlag(j, lc-1)
		}
	}

	for i := range board {
		for j := range board[i] {
			if flag[i][j] == 0 {
				board[i][j] = 'X'
			}
		}
	}
}
