package main

import (
	"fmt"
	"strings"
)

func main() {
	moves := [][]int{{0, 0}, {2, 0}, {1, 1}, {2, 1}, {2, 2}}
	// moves := [][]int{{0, 0}, {1, 1}, {0, 1}, {0, 2}, {1, 0}, {2, 0}}
	// moves := [][]int{{0, 0}, {1, 1}, {2, 0}, {1, 0}, {1, 2}, {2, 1}, {0, 1}, {0, 2}, {2, 2}}
	// moves := [][]int{{0, 0}, {1, 1}}
	fmt.Println("Output: ", tictactoe(moves))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2.1 MB, 在所有 Go 提交中击败了25.64%的用户
func tictactoe(moves [][]int) string {
	s := strings.Repeat(" ", 9)
	for i := range moves {
		s = move(s, moves[i][0]*3+moves[i][1], i%2 == 0)
	}
	return result(s)
}

func move(s string, p int, isX bool) string {
	first := s[0:p]
	second := s[p+1:]
	if isX {
		first = first + "X"
	} else {
		first = first + "O"
	}
	return first + second
}

func result(s string) string {
	if (s[0] == s[1] && s[1] == s[2]) || (s[0] == s[3] && s[3] == s[6]) || (s[0] == s[4] && s[4] == s[8]) {
		if s[0] == 'X' {
			return "A"
		} else if s[0] == 'O' {
			return "B"
		}
	}
	if (s[3] == s[4] && s[4] == s[5]) || (s[1] == s[4] && s[4] == s[7]) || (s[6] == s[4] && s[4] == s[2]) {
		if s[4] == 'X' {
			return "A"
		} else if s[4] == 'O' {
			return "B"
		}
	}
	if (s[6] == s[7] && s[7] == s[8]) || (s[2] == s[5] && s[5] == s[8]) {
		if s[8] == 'X' {
			return "A"
		} else if s[8] == 'O' {
			return "B"
		}
	}
	if strings.Count(s, " ") == 0 {
		return "Draw"
	}
	return "Pending"
}
