package main

import "fmt"

func main() {
	M := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	fmt.Printf("Output: %v\n", imageSmoother(M))
}

func imageSmoother(M [][]int) [][]int {
	m := make([][]int, len(M))
	var sum, cnt int
	for i := range M {
		m[i] = make([]int, len(M[i]))
		for j := range M[i] {
			sum, cnt = M[i][j], 1
			if i > 0 {
				sum = sum + M[i-1][j]
				cnt = cnt + 1
				if j > 0 {
					sum = sum + M[i-1][j-1]
					cnt = cnt + 1
				}
				if j < len(M[i])-1 {
					sum = sum + M[i-1][j+1]
					cnt = cnt + 1
				}
			}
			if i < len(M)-1 {
				sum = sum + M[i+1][j]
				cnt = cnt + 1
				if j > 0 {
					sum = sum + M[i+1][j-1]
					cnt = cnt + 1
				}
				if j < len(M[i])-1 {
					sum = sum + M[i+1][j+1]
					cnt = cnt + 1
				}
			}
			if j > 0 {
				sum = sum + M[i][j-1]
				cnt = cnt + 1
			}
			if j < len(M[i])-1 {
				sum = sum + M[i][j+1]
				cnt = cnt + 1
			}
			m[i][j] = sum / cnt
		}
	}
	return m
}
