package main

import "fmt"

// https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/
func main() {
	// matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16}}
	// matrix := [][]int{{}}
	fmt.Println("Output: ", spiralOrder(matrix))
}

//执行用时 :12 ms, 在所有 Go 提交中击败了95.07%的用户
//内存消耗 :6.1 MB, 在所有 Go 提交中击败了100.00%的用户
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	flags := make([][]int, len(matrix))
	for i := range matrix {
		flags[i] = make([]int, len(matrix[i]))
	}

	cur := 0 //0:Right, 1:Down, 2:Left, 3:Up
	res := make([]int, len(matrix)*len(matrix[0]))
	i, j := 0, 0
	index := 0
	for index < len(res) {
		res[index] = matrix[i][j]
		flags[i][j] = 1
		switch cur {
		case 0:
			if j == len(matrix[0])-1 || flags[i][j+1] == 1 {
				i++
				cur++
			} else {
				j++
			}
		case 1:
			if i == len(matrix)-1 || flags[i+1][j] == 1 {
				j--
				cur++
			} else {
				i++
			}
		case 2:
			if j == 0 || flags[i][j-1] == 1 {
				i--
				cur++
			} else {
				j--
			}
		case 3:
			if i == 0 || flags[i-1][j] == 1 {
				j++
				cur = 0
			} else {
				i--
			}
		}
		index++
	}
	return res
}
