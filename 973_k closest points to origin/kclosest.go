package main

import (
	"fmt"
	"sort"
)

func main() {
	points := [][]int{{3, 3}, {5, -1}, {-2, 4}}
	K := 2
	fmt.Println("Output: ", kClosest(points, K))
}

//执行用时 :136 ms, 在所有 Go 提交中击败了72.60%的用户
//内存消耗 :7.5 MB, 在所有 Go 提交中击败了100.00%的用户
func kClosest(points [][]int, K int) [][]int {
	if len(points) == 0 || K == 0 {
		return [][]int{}
	}

	sort.Slice(points, func(i, j int) bool {
		return sqrt(points[i]) < sqrt(points[j])
	})

	return points[:K]
}

func sqrt(p []int) int {
	return p[0]*p[0] + p[1]*p[1]
}
