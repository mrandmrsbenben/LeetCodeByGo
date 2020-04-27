package main

import "fmt"

func main() {
	// matrix := [][]byte{
	// 	{'1', '0', '1', '0', '0'},
	// 	{'1', '0', '1', '1', '1'},
	// 	{'1', '1', '1', '1', '1'},
	// 	{'1', '0', '0', '1', '0'}}
	// matrix := [][]byte{
	// 	{'1', '1'},
	// 	{'1', '1'}}
	matrix := [][]byte{
		{'0', '0', '1', '0'},
		{'1', '1', '1', '1'},
		{'1', '1', '1', '1'},
		{'1', '1', '1', '0'},
		{'1', '1', '0', '0'},
		{'1', '1', '1', '1'},
		{'1', '1', '1', '0'}}
	// matrix := [][]byte{
	// 	{'0', '0', '0', '1', '0', '1', '1', '1'},
	// 	{'0', '1', '1', '0', '0', '1', '0', '1'},
	// 	{'1', '0', '1', '1', '1', '1', '0', '1'},
	// 	{'0', '0', '0', '1', '0', '0', '0', '0'},
	// 	{'0', '0', '1', '0', '0', '0', '1', '0'},
	// 	{'1', '1', '1', '0', '0', '1', '1', '1'},
	// 	{'1', '0', '0', '1', '1', '0', '0', '1'},
	// 	{'0', '1', '0', '0', '1', '1', '0', '0'},
	// 	{'1', '0', '0', '1', '0', '0', '0', '0'}}
	fmt.Println("Output: ", maximalSquare(matrix))
}

//执行用时 :4 ms, 在所有 Go 提交中击败了68.62%的用户
//内存消耗 :4.3 MB, 在所有 Go 提交中击败了33.33%的用户
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	dp := make([][]int, len(matrix)+1)
	for i := range dp {
		dp[i] = make([]int, len(matrix[0])+1)
	}

	max := 0
	for i := len(matrix) - 1; i >= 0; i-- {
		for j := len(matrix[i]) - 1; j >= 0; j-- {
			if matrix[i][j] == '0' {
				dp[i][j] = 0
			} else {
				dp[i][j] = 1 + min(dp[i+1][j], dp[i][j+1], dp[i+1][j+1])
				if dp[i][j] > max {
					max = dp[i][j]
				}
			}
		}
	}
	return max * max
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	} else if c < b {
		return c
	}
	return b
}
