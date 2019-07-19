package main

import "fmt"

func main() {
	fmt.Printf("Output: %v\n", getRow(4))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2.1 MB, 在所有 Go 提交中击败了51.22%的用户
func getRow(rowIndex int) []int {
	var getNextRow func(row []int) []int
	getNextRow = func(row []int) []int {
		row = append(row, 1)
		for i := (len(row)+1)/2 - 1; i > 0; i-- {
			v := row[i-1] + row[i]
			row[i], row[len(row)-1-i] = v, v
		}
		if rowIndex >= len(row) {
			return getNextRow(row)
		}
		return row
	}
	r := []int{1}
	if rowIndex > 0 {
		return getNextRow(r)
	}
	return r
}

func getRowV1(rowIndex int) []int {
	rows := make([][]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		rows[i] = make([]int, i+1)
		rows[i][0] = 1
		rows[i][len(rows[i])-1] = 1
		if i > 1 {
			for j := 1; j < len(rows[i])-1; j++ {
				rows[i][j] = rows[i-1][j-1] + rows[i-1][j]
			}
		}
	}
	return rows[rowIndex]
}
