package main

import "fmt"

func main() {
	// grid := [][]int{
	// 	{4, 3, 2, -1},
	// 	{3, 2, 1, -1},
	// 	{1, 1, -1, -2},
	// 	{-1, -1, -2, -3}}
	// grid := [][]int{{3, 2}, {1, 0}}
	// grid := [][]int{{1, -1}, {-1, -1}}
	grid := [][]int{{-1}}
	fmt.Println("Output: ", countNegatives(grid))
}

func countNegatives(grid [][]int) int {
	cnt := 0
	endC := 0
	for r := len(grid) - 1; r >= 0; r-- {
		for c := len(grid[r]) - 1; c >= endC; c-- {
			if grid[r][c] < 0 {
				cnt++
			} else {
				endC = c - 1
				break
			}
		}
	}
	return cnt
}
