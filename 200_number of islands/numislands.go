package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := []string{
		"11110",
		"11010",
		"11001",
		"00101"}
	// "11000",
	// "11000",
	// "00100",
	// "00011"}
	// "111",
	// "010",
	// "111"}
	grid := make([][]byte, len(str))
	for i := range str {
		grid[i] = make([]byte, len(str[i]))
		for j := range str[i] {
			grid[i][j] = str[i][j]
		}
	}
	fmt.Println(numIslands(grid))
}

//执行用时 :4 ms, 在所有 Go 提交中击败了76.92%的用户
//内存消耗 :3.9 MB, 在所有 Go 提交中击败了15.29%的用户
func numIslands(grid [][]byte) int {
	cnt := 0
	var curIsland [][2]int
	var size, x, y int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				grid[i][j] = '0'
				curIsland = [][2]int{{i, j}}
				for len(curIsland) > 0 {
					size = len(curIsland)
					for k := 0; k < size; k++ {
						x, y = curIsland[k][0], curIsland[k][1]
						if x < len(grid)-1 && grid[x+1][y] == '1' {
							grid[x+1][y] = '0'
							curIsland = append(curIsland, [2]int{x + 1, y})
						}
						if y < len(grid[i])-1 && grid[x][y+1] == '1' {
							grid[x][y+1] = '0'
							curIsland = append(curIsland, [2]int{x, y + 1})
						}
						if x > 0 && grid[x-1][y] == '1' {
							grid[x-1][y] = '0'
							curIsland = append(curIsland, [2]int{x - 1, y})
						}
						if y > 0 && grid[x][y-1] == '1' {
							grid[x][y-1] = '0'
							curIsland = append(curIsland, [2]int{x, y - 1})
						}
					}
					curIsland = curIsland[size:]
				}
				cnt++
			}
		}
	}
	return cnt
}

// Slow Version
func numIslandsV1(grid [][]byte) int {
	cnt := 0
	islandMap := make(map[string]int)
	var key string
	var curIsland [][2]int
	var size, x, y int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				key = islandKey(i, j)
				if _, ok := islandMap[key]; ok {
					continue
				} else {
					curIsland = [][2]int{{i, j}}
					for len(curIsland) > 0 {
						size = len(curIsland)
						for k := 0; k < size; k++ {
							x, y = curIsland[k][0], curIsland[k][1]
							if x < len(grid)-1 && grid[x+1][y] == '1' {
								if _, ok := islandMap[islandKey(x+1, y)]; !ok {
									curIsland = append(curIsland, [2]int{x + 1, y})
									islandMap[islandKey(x+1, y)] = 1
								}
							}
							if y < len(grid[i])-1 && grid[x][y+1] == '1' {
								if _, ok := islandMap[islandKey(x, y+1)]; !ok {
									curIsland = append(curIsland, [2]int{x, y + 1})
									islandMap[islandKey(x, y+1)] = 1
								}
							}
							if x > 0 && grid[x-1][y] == '1' {
								if _, ok := islandMap[islandKey(x-1, y)]; !ok {
									curIsland = append(curIsland, [2]int{x - 1, y})
									islandMap[islandKey(x-1, y)] = 1
								}
							}
							if y > 0 && grid[x][y-1] == '1' {
								if _, ok := islandMap[islandKey(x, y-1)]; !ok {
									curIsland = append(curIsland, [2]int{x, y - 1})
									islandMap[islandKey(x, y-1)] = 1
								}
							}
						}
						curIsland = curIsland[size:]
					}
					islandMap[key] = 1
					cnt++
				}
			}
		}
	}
	return cnt
}

func islandKey(i, j int) string {
	return strconv.Itoa(i) + "," + strconv.Itoa(j)
}
