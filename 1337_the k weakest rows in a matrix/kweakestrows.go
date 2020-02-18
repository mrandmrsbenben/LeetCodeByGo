package main

import "fmt"

func main() {
	// mat := [][]int{{1, 1, 0, 0, 0}, {1, 1, 1, 1, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 0, 0}, {1, 1, 1, 1, 1}}
	// k := 5
	// mat := [][]int{{1, 0, 0, 0}, {1, 1, 1, 1}, {1, 0, 0, 0}, {1, 0, 0, 0}}
	// k := 4
	mat := [][]int{{1, 0}, {1, 0}, {1, 0}, {1, 1}}
	k := 4
	// mat := [][]int{{1, 1, 1, 1, 1}, {1, 0, 0, 0, 0}, {1, 1, 0, 0, 0}, {1, 1, 1, 1, 0}, {1, 1, 1, 1, 1}}
	// k := 5
	fmt.Println("Output: ", kWeakestRows(mat, k))
}

//执行用时 :12 ms, 在所有 Go 提交中击败了97.44%的用户
//内存消耗 :5.1 MB, 在所有 Go 提交中击败了100.00%的用户
func kWeakestRows(mat [][]int, k int) []int {
	rows := make([]int, k)
	cnt := 0
	rowmap := make(map[int]int)
	for i := 0; i < len(mat[0]) && cnt < k; i++ {
		for j := 0; j < len(mat); j++ {
			if rowmap[j] == 0 {
				if mat[j][i] == 0 {
					rows[cnt] = j
					cnt++
					if cnt == k {
						break
					}
					rowmap[j] = 1
				}
			}
		}
	}
	for i := 0; i < len(mat) && cnt < k; i++ {
		if mat[i][len(mat[i])-1] == 1 {
			rows[cnt] = i
			cnt++
		}
	}

	return rows
}
