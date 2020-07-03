package main

import "fmt"

func main() {
	cells := []int{0, 1, 0, 1, 1, 0, 0, 1}
	N := 30
	fmt.Println("Output: ", prisonAfterNDays(cells, N))
}

//执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗：2.4 MB, 在所有 Go 提交中击败了100.00%的用户
func prisonAfterNDays(cells []int, N int) []int {
	if N == 0 {
		return cells
	}

	temp := make([]int, 8)
	i := 0
	if N%14 == 0 {
		N = 14
	} else {
		N %= 14
	}
	for i < N {
		for j := 1; j < len(cells)-1; j++ {
			if cells[j-1] == cells[j+1] {
				temp[j] = 1
			} else {
				temp[j] = 0
			}
		}
		cells = make([]int, 8)
		copy(cells, temp)
		temp = make([]int, 8)
		i++
	}
	return cells
}
