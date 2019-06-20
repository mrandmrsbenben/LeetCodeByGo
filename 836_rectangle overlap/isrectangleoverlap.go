package main

import "fmt"

func main() {
	rec1 := []int{2, 0, 0, 2}
	rec2 := []int{1, 1, 3, 3}
	// rec1 := []int{1, 1, 0, 0}
	// rec2 := []int{3, 0, 2, 1}
	// rec1 := []int{0, 2, 2, 0}
	// rec2 := []int{0, 0, 2, 2}
	// rec1 := []int{7, 8, 13, 15}
	// rec2 := []int{10, 8, 12, 20}
	// rec1 := []int{4, 4, 14, 7}
	// rec2 := []int{4, 3, 8, 8}
	// rec1 := []int{-7, -3, 10, 5}
	// rec2 := []int{-6, -5, 5, 10}
	// rec1 := []int{-6, -10, 9, 2}
	// rec2 := []int{0, 5, 4, 8}
	fmt.Printf("Output: %v\n", isRectangleOverlap(rec1, rec2))
}

// 执行用时 : 0 ms, 在Rectangle Overlap的Go提交中击败了100.00% 的用户
// 内存消耗 : 2.1 MB, 在Rectangle Overlap的Go提交中击败了10.00% 的用户
func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	rec1 = sort(rec1)
	rec2 = sort(rec2)

	x1 := rec1[0]
	y1 := rec1[1]
	width1 := rec1[2] - rec1[0]
	height1 := rec1[3] - rec1[1]

	x2 := rec2[0]
	y2 := rec2[1]
	width2 := rec2[2] - rec2[0]
	height2 := rec2[3] - rec2[1]

	endx := max(x1+width1, x2+width2)
	startx := min(x1, x2)
	width := width1 + width2 - (endx - startx)

	endy := max(y1+height1, y2+height2)
	starty := min(y1, y2)
	height := height1 + height2 - (endy - starty)

	if width <= 0 || height <= 0 {
		return false
	}
	var area, area1, area2, ratio float32
	area = float32(width * height)
	area1 = float32(width1 * height1)
	area2 = float32(width2 * height2)
	ratio = area / (area1 + area2 - area)
	if ratio > 0 {
		return true
	}
	return false
}

func sort(rec []int) []int {
	if rec[0] > rec[2] {
		rec[0], rec[2] = rec[2], rec[0]
	}
	if rec[1] > rec[3] {
		rec[1], rec[3] = rec[3], rec[1]
	}
	return rec
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
