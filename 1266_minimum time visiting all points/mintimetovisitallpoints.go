package main

import (
	"fmt"
)

func main() {
	points := [][]int{{1, 1}, {3, 4}, {-1, 0}}
	fmt.Println("Output: ", minTimeToVisitAllPoints(points))
}

//执行用时 :4 ms, 在所有 Go 提交中击败了91.72%的用户
//内存消耗 :3.4 MB, 在所有 Go 提交中击败了100.00%的用户
func minTimeToVisitAllPoints(points [][]int) int {
	if len(points) == 1 {
		return 0
	}

	t := 0
	for i := 1; i < len(points); i++ {
		t += visitTime(points[i-1], points[i])
	}
	return t
}

func visitTime(p0, p1 []int) int {
	t := 0
	if p1[0] == p0[0] {
		t = abs(p1[1] - p0[1])
	} else if p1[1] == p0[1] {
		t = abs(p1[0] - p0[0])
	} else {
		if abs(p1[0]-p0[0]) < abs(p1[1]-p0[1]) {
			t = abs(p1[1] - p0[1])
		} else {
			t = abs(p1[0] - p0[0])
		}
	}
	return t
}

func abs(n int) int {
	if n < 0 {
		n = -1 * n
	}
	return n
}
