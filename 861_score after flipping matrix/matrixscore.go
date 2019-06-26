package main

import "fmt"

func main() {
	A := [][]int{{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0}}
	// A := [][]int{{0, 1, 1}, {1, 1, 1}, {0, 1, 0}}
	fmt.Println("Output:", matrixScore(A))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2.3 MB, 在所有 Go 提交中击败了20.00%的用户
func matrixScore(A [][]int) int {
	for i := range A[0] {
		if A[0][i] == 0 {
			for j := range A {
				if A[j][i] == 0 {
					A[j][i] = 1
				} else {
					A[j][i] = 0
				}
			}
		}
	}
	for i := range A {
		if A[i][0] == 0 {
			for j := range A[i] {
				if A[i][j] == 0 {
					A[i][j] = 1
				} else {
					A[i][j] = 0
				}
			}
		}
	}
	sum, cnt := 0, 0
	base := 1
	for i := len(A[0]) - 1; i >= 0; i-- {
		for j := range A {
			if A[j][i] == 1 {
				cnt++
			}
		}
		if cnt >= len(A)-cnt {
			sum += cnt * base
		} else {
			sum += (len(A) - cnt) * base
		}
		base *= 2
		cnt = 0
	}
	return sum
}
