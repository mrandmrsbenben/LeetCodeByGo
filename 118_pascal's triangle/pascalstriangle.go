package main

import "fmt"

func main() {
	numRows := 6
	fmt.Printf("Output:%v\n", generate(numRows))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2.4 MB, 在所有 Go 提交中击败了22.95%的用户
func generate(numRows int) [][]int {
	tri := make([][]int, numRows)
	row := 0
	var gen func(rows [][]int)
	gen = func(rows [][]int) {
		row++
		if row == numRows {
			return
		}
		rows[row] = make([]int, row+1)
		rows[row][0], rows[row][row] = 1, 1
		prevrow := rows[row-1]
		for i := 1; i < row; i++ {
			rows[row][i] = prevrow[i-1] + prevrow[i]
		}
		gen(rows)
	}
	if numRows > 0 {
		tri[0] = []int{1}
		gen(tri)
	}
	return tri
}

func generateV1(numRows int) [][]int {
	tri := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		tri[i] = make([]int, i+1)
		tri[i][0], tri[i][i] = 1, 1
		for j := 1; j < i; j++ {
			tri[i][j] = tri[i-1][j-1] + tri[i-1][j]
		}
	}
	return tri
}
