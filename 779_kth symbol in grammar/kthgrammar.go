package main

import (
	"fmt"
	"strconv"
)

func main() {
	// fmt.Println("Output:", kthGrammar(1, 1))
	// fmt.Println("Output:", kthGrammar(2, 1))
	// fmt.Println("Output:", kthGrammar(2, 2))
	fmt.Println("Output:", kthGrammar(4, 5), kthGrammarV1(4, 5))
	fmt.Println("Output:", kthGrammar(30, 434991989), kthGrammarV1(30, 434991989))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :1.9 MB, 在所有 Go 提交中击败了100.00%的用户
func kthGrammar(N int, K int) int {
	if N == 1 {
		return 0
	}
	prevK := kthGrammar(N-1, (K+1)/2)
	if prevK == 1 {
		return K % 2
	}
	return 1 - K%2
}

func kthGrammarV1(N int, K int) int {
	if N == 1 {
		return 0
	}
	if K <= 1<<uint(N-2) {
		return kthGrammarV1(N-1, K)
	}
	return kthGrammarV1(N-1, K-(1<<uint(N-2))) ^ 1
}

func kthGrammarSlowVersion(N int, K int) int {
	s := getKLineString(N)
	x, _ := strconv.Atoi(s[K-1 : K])
	return x
}

func getKLineString(N int) string {
	if N == 1 {
		return "0"
	}
	s := getKLineString(N - 1)
	var newS string
	for _, b := range s {
		if b == '0' {
			newS += "1"
		} else {
			newS += "0"
		}
	}
	return s + newS
}
