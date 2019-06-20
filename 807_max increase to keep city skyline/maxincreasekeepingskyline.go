package main

import "fmt"

func main() {
	grid := [][]int{{3, 0, 8, 4}, {2, 4, 5, 7}, {9, 2, 6, 3}, {0, 3, 1, 0}}
	fmt.Println("Output:", maxIncreaseKeepingSkyline(grid))
}

//执行用时 :12 ms, 在所有 Go 提交中击败了72.00%的用户
//内存消耗 :3.5 MB, 在所有 Go 提交中击败了66.67%的用户
func maxIncreaseKeepingSkyline(grid [][]int) int {
	maxH := make([]int, len(grid))
	maxV := make([]int, len(grid))
	// get max in horizon and vertical
	for i := range grid {
		for j := range grid[i] {
			if maxH[i] < grid[i][j] {
				maxH[i] = grid[i][j]
			}
			if maxV[j] < grid[i][j] {
				maxV[j] = grid[i][j]
			}
		}
	}
	inc := 0
	// get increase summary
	for i := range grid {
		for j := range grid[i] {
			inc += getInc(grid[i][j], maxH[i], maxV[j])
		}
	}
	return inc
}

// get increase for each buliding
func getInc(n int, maxH, maxV int) int {
	if maxH > maxV {
		return maxV - n
	}
	return maxH - n
}
