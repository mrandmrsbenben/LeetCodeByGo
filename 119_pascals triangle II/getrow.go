package main

import "fmt"

func main() {
	fmt.Printf("Output: %v\n", getRow(4))
}

func getRow(rowIndex int) []int {
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
