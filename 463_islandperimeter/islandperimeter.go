package main

import "fmt"

func main() {
	grid := [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}
	// grid := [][]int{{0, 1}}
	fmt.Printf("Output: %d\n", islandPerimeter(grid))
}

func islandPerimeter(grid [][]int) int {
	peri := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				if j == 0 || grid[i][j-1] == 0 {
					peri = peri + 1
				}
				if i == 0 || grid[i-1][j] == 0 {
					peri = peri + 1
				}
				if j == len(grid[i])-1 || grid[i][j+1] == 0 {
					peri = peri + 1
				}
				if i == len(grid)-1 || grid[i+1][j] == 0 {
					peri = peri + 1
				}
			}
		}
	}
	return peri
}
