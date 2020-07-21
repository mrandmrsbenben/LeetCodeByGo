package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "PAYPALISHIRING"
	numRows := 3
	// s := "AB"
	// numRows := 1
	fmt.Println("Output: ", convert(s, numRows))
}

//执行用时：28 ms, 在所有 Go 提交中击败了11.33%的用户
//内存消耗：6.6 MB, 在所有 Go 提交中击败了25.00%的用户
func convert(s string, numRows int) string {
	if len(s) <= 1 || numRows == 1 {
		return s
	}
	strs := make([]string, 0)
	i, j := 0, 0
	row := ""
	for i < len(s) {
		if j == 0 {
			if i+numRows < len(s) {
				row = s[i : i+numRows]
			} else {
				row = s[i:] + strings.Repeat(" ", numRows-len(s[i:]))
			}
			i += numRows
		} else {
			row = strings.Repeat(" ", numRows-1-j) + string(s[i]) + strings.Repeat(" ", j)
			i++
		}
		strs = append(strs, row)
		j++
		if j == numRows-1 {
			j = 0
		}
	}

	res := ""
	for i := 0; i < numRows; i++ {
		for j := range strs {
			if strs[j][i] != ' ' {
				res += string(strs[j][i])
			}
		}
	}

	return res
}

//执行用时：112 ms, 在所有 Go 提交中击败了5.21%的用户
//内存消耗：12.4 MB, 在所有 Go 提交中击败了25.00%的用户
func convertV1(s string, numRows int) string {
	if len(s) <= 1 || numRows == 1 {
		return s
	}
	strs := make([][]string, 0)
	i, j := 0, 0
	for i < len(s) {
		row := make([]string, numRows)
		if j == 0 {
			for k := 0; k < numRows && i < len(s); k++ {
				row[k] = string(s[i])
				i++
			}
		} else {
			row[numRows-1-j] = string(s[i])
			i++
		}
		strs = append(strs, row)
		j++
		if j == numRows-1 {
			j = 0
		}
	}

	res := ""
	for i := 0; i < numRows; i++ {
		for j := range strs {
			res += strs[j][i]
		}
	}

	return res
}
