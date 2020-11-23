package main

import (
	"fmt"
	"sort"
)

func main() {
	// points := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
	// points := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	points := [][]int{{1, 5}, {2, 3}, {2, 3}, {3, 4}, {3, 5}}
	fmt.Println("Output: ", findMinArrowShots(points))
}

//执行用时：88 ms, 在所有 Go 提交中击败了60.31%的用户
//内存消耗：7.3 MB, 在所有 Go 提交中击败了43.85%的用户
func findMinArrowShots(points [][]int) int {
	if len(points) <= 1 {
		return len(points)
	}

	// sort points array
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] < points[j][0] {
			return true
		} else if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		}
		return false
	})

	// merge crossed element
	cur := points[0]
	cnt := 1
	for i := 1; i < len(points); i++ {
		if cur[1] < points[i][0] {
			cnt++
			cur = points[i]
		} else if cur[1] == points[i][0] {
			cur[0] = cur[1]
		} else {
			cur[0] = points[i][0]
			if cur[1] > points[i][1] {
				cur[1] = points[i][1]
			}
		}
	}

	return cnt
}
