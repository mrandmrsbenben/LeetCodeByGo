package main

import "fmt"

func main() {
	grid := [][]int{{0}, {0}, {1}, {1}, {2}, {2}, {1}, {0}, {2}, {2}}
	fmt.Printf("Output: %d\n", orangesRotting(grid))
}

func orangesRotting(grid [][]int) int {
	time := 0
	rotting := 0
	var cnt int
	for {
		cnt = 0
		for i := range grid {
			for j := range grid[i] {
				if grid[i][j] == 2 {
					if j > 0 && grid[i][j-1] == 1 {
						grid[i][j-1] = -2
						rotting++
					}
					if j < len(grid[i])-1 && grid[i][j+1] == 1 {
						grid[i][j+1] = -2
						rotting++
					}
					if i > 0 && grid[i-1][j] == 1 {
						grid[i-1][j] = -2
						rotting++
					}
					if i < len(grid)-1 && grid[i+1][j] == 1 {
						grid[i+1][j] = -2
						rotting++
					}
				} else if grid[i][j] == -2 {
					grid[i][j] = 2
				} else if grid[i][j] == 1 {
					cnt = cnt + 1
				}
			}
		}
		for i := range grid {
			for j := range grid[i] {
				if grid[i][j] == -2 {
					grid[i][j] = 2
				}
			}
		}
		if rotting == 0 {
			break
		}
		time = time + 1
		rotting = 0
	}
	if cnt > 0 {
		return -1
	}
	return time
}
