package main

import (
	"fmt"
	"strings"
)

func main() {
	A := "abcd"
	B := "cdabcdab"
	// A := "aaaaaaaaaaaaaaaaaaaaaab"
	// B := "ba"
	fmt.Println("Output:", repeatedStringMatch(A, B))
}

//执行用时 :356 ms, 在所有 Go 提交中击败了10.00%的用户
//内存消耗 :6.7 MB, 在所有 Go 提交中击败了50.00%的用户
func repeatedStringMatch(A string, B string) int {
	var start, end int
	if len(B) > len(A) {
		if len(B)%len(A) == 0 {
			end = len(B) / len(A) * 2
		} else {
			end = (len(B)/len(A) + 1) * 2
		}
		start = len(B) / len(A)
	} else {
		end = 2
		start = 1
	}
	s := strings.Repeat(A, start)
	for i := start; i <= end; i++ {
		if strings.Index(s, B) != -1 {
			return i
		}
		s += A
	}
	return -1
}
