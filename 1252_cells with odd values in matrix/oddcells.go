package main

import "fmt"

func main() {
	// n := 2
	// m := 3
	// indices := [][]int{{0, 1}, {1, 1}}
	n := 2
	m := 2
	indices := [][]int{{1, 1}, {0, 0}}
	fmt.Println("Output: ", oddCells(n, m, indices))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2.7 MB, 在所有 Go 提交中击败了41.67%的用户
func oddCells(n int, m int, indices [][]int) int {
	cells := make([][]int, n)
	for i := range cells {
		cells[i] = make([]int, m)
	}

	cnt := 0
	var ri, ci int
	for i := range indices {
		ri = indices[i][0]
		ci = indices[i][1]

		//add to row
		for j := 0; j < m; j++ {
			cells[ri][j]++
		}

		//add to col
		for j := 0; j < n; j++ {
			cells[j][ci]++
		}
	}

	for i := range cells {
		for j := range cells[i] {
			if cells[i][j]%2 == 1 {
				cnt++
			}
		}
	}

	return cnt
}
