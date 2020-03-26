package main

import "fmt"

// https://leetcode-cn.com/problems/li-wu-de-zui-da-jie-zhi-lcof/submissions/
func main() {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1}}
	fmt.Println("Output: ", maxValue(grid))
}

//执行用时 :8 ms, 在所有 Go 提交中击败了94.84%的用户
//内存消耗 :3.9 MB, 在所有 Go 提交中击败了100.00%的用户
func maxValue(grid [][]int) int {
	for i := range grid {
		max := 0
		for j := range grid[i] {
			if i > 0 && max < grid[i-1][j] {
				max = grid[i-1][j]
			}
			if j > 0 && max < grid[i][j-1] {
				max = grid[i][j-1]
			}
			grid[i][j] += max
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

//执行用时 :12 ms, 在所有 Go 提交中击败了26.45%的用户
//内存消耗 :4.6 MB, 在所有 Go 提交中击败了100.00%的用户
func maxValueV1(grid [][]int) int {
	m := make([][]int, len(grid))
	for i := range m {
		m[i] = make([]int, len(grid[0]))
	}
	max := 0
	l, w := len(grid), len(grid[0])

	var find func(i, j int) int
	find = func(i, j int) int {
		if m[i][j] > 0 {
			return m[i][j]
		}

		sum := grid[i][j]
		d, r := 0, 0

		if i+1 < l {
			r = find(i+1, j)
		}
		if j+1 < w {
			d = find(i, j+1)
		}
		if r > d {
			sum += r
		} else {
			sum += d
		}

		m[i][j] = sum
		if sum > max {
			max = sum
		}
		return sum
	}
	find(0, 0)
	fmt.Println(m)
	return max
}
