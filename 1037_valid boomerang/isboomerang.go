package main

import "fmt"

func main() {
	// points := [][]int{{1, 1}, {2, 3}, {3, 2}}
	points := [][]int{{1, 1}, {2, 2}, {3, 3}}
	fmt.Printf("Output: %v\n", isBoomerang(points))
}

// 执行用时 : 0 ms, 在Valid Boomerang的Go提交中击败了100.00% 的用户
// 内存消耗 : 2.1 MB, 在Valid Boomerang的Go提交中击败了100.00% 的用户
func isBoomerang(points [][]int) bool {
	if isEqualPoints(points[0], points[1]) || isEqualPoints(points[0], points[2]) ||
		isEqualPoints(points[1], points[2]) {
		return false
	}
	angle := float64(points[1][1]-points[0][1]) / float64(points[1][0]-points[0][0])
	for i := 1; i < len(points)-1; i++ {
		if angle == float64(points[i+1][1]-points[i][1])/float64(points[i+1][0]-points[i][0]) {
			return false
		}
	}
	return true
}

func isEqualPoints(p1, p2 []int) bool {
	if p1[0] == p2[0] && p1[1] == p2[1] {
		return true
	}
	return false
}
