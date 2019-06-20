package main

import "fmt"

func main() {
	// common.MakeTestCases("surfaceArea")
	input1 := [][]int{{0, 2, 0}}
	fmt.Printf("Expect1: 10\n")
	fmt.Printf("Output1: %v\n", surfaceArea(input1))
	input2 := [][]int{{1, 2}, {3, 4}}
	fmt.Printf("Expect2: 34\n")
	fmt.Printf("Output2: %v\n", surfaceArea(input2))
	input3 := [][]int{{1, 0}, {0, 2}}
	fmt.Printf("Expect3: 16\n")
	fmt.Printf("Output3: %v\n", surfaceArea(input3))
	input4 := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	fmt.Printf("Expect4: 32\n")
	fmt.Printf("Output4: %v\n", surfaceArea(input4))
	input5 := [][]int{{2, 2, 2}, {2, 1, 2}, {2, 2, 2}}
	fmt.Printf("Expect5: 46\n")
	fmt.Printf("Output5: %v\n", surfaceArea(input5))
}

func surfaceArea(grid [][]int) int {
	area := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				continue
			}
			area = area + grid[i][j]*4 + 2
			if i > 0 {
				if grid[i][j] > grid[i-1][j] {
					area = area - grid[i-1][j]
				} else {
					area = area - grid[i][j]
				}
			}
			if i < len(grid)-1 {
				if grid[i][j] > grid[i+1][j] {
					area = area - grid[i+1][j]
				} else {
					area = area - grid[i][j]
				}
			}
			if j > 0 {
				if grid[i][j] > grid[i][j-1] {
					area = area - grid[i][j-1]
				} else {
					area = area - grid[i][j]
				}
			}
			if j < len(grid[i])-1 {
				if grid[i][j] > grid[i][j+1] {
					area = area - grid[i][j+1]
				} else {
					area = area - grid[i][j]
				}
			}
		}
	}
	return area
}
