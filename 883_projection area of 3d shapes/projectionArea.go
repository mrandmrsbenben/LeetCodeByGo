package main

import (
	"fmt"
)

func main() {
	// grid := [][]int{{2}}
	// fmt.Printf("Output: %d\n", projectionArea(grid))
	// common.MakeTestCases("projectionArea")
	input1 := [][]int{{2}}
	fmt.Printf("Expect1: 5\n")
	fmt.Printf("Output1: %v\n", projectionArea(input1))
	input2 := [][]int{{1, 2}, {3, 4}}
	fmt.Printf("Expect2: 17\n")
	fmt.Printf("Output2: %v\n", projectionArea(input2))
	input3 := [][]int{{1, 0}, {0, 2}}
	fmt.Printf("Expect3: 8\n")
	fmt.Printf("Output3: %v\n", projectionArea(input3))
	input4 := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	fmt.Printf("Expect4: 14\n")
	fmt.Printf("Output4: %v\n", projectionArea(input4))
	input5 := [][]int{{2, 2, 2}, {2, 1, 2}, {2, 2, 2}}
	fmt.Printf("Expect5: 21\n")
	fmt.Printf("Output5: %v\n", projectionArea(input5))
}

func projectionArea(grid [][]int) int {
	area := 0
	maxX, maxY := 0, 0
	maxZ := make([]int, len(grid))
	for i := range grid {
		maxY = 0
		for j := range grid[i] {
			if grid[i][j] > 0 {
				maxX = maxX + 1
			}
			if grid[i][j] > maxY {
				maxY = grid[i][j]
			}
			if grid[i][j] > maxZ[j] {
				maxZ[j] = grid[i][j]
			}
		}
		area = area + maxY
	}
	area = area + maxX
	for _, z := range maxZ {
		area = area + z
	}
	return area
}
