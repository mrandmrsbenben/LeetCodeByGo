package main

import "fmt"

func main() {
	trust := [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}}
	N := 4
	fmt.Printf("Output: %d\n", findJudge(N, trust))
}

func findJudge(N int, trust [][]int) int {
	tcount := make([]int, N)
	pcount := make([]int, N)
	for i := range trust {
		tcount[trust[i][1]-1] = tcount[trust[i][1]-1] + 1
		pcount[trust[i][0]-1] = tcount[trust[i][0]-1] + 1
	}
	for i := range tcount {
		if tcount[i] == N-1 && pcount[i] == 0 {
			return i + 1
		}
	}
	return -1
}
