package main

import "fmt"

func main() {
	// A, B := []int{1, 4, 2}, []int{1, 2, 4}
	A, B := []int{2, 5, 1, 2, 5}, []int{10, 5, 2, 1, 5, 2}
	fmt.Println("Output: ", maxUncrossedLines(A, B))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :6.1 MB, 在所有 Go 提交中击败了100.00%的用户
func maxUncrossedLines(A []int, B []int) int {
	dp := make([][]int, len(A)+1)
	for i := range dp {
		dp[i] = make([]int, len(B)+1)
	}

	for i := range A {
		for j := range B {
			if A[i] == B[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[len(A)][len(B)]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
