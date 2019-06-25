package main

import "fmt"

func main() {
	fmt.Println("Output:", generateMatrix(4))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2.6 MB, 在所有 Go 提交中击败了17.50%的用户
func generateMatrix(n int) [][]int {
	if n == 1 {
		return [][]int{{1}}
	}
	// initialize matrix
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	v := 1
	d := 2 //Directions 0:W 1:N 2:E 3:S
	x, y := 0, 0
	for {
		matrix[y][x] = v
		switch d {
		case 0:
			if x == 0 || matrix[y][x-1] > 0 {
				d = 1
				y--
			} else {
				x--
			}
		case 1:
			if y == 0 || matrix[y-1][x] > 0 {
				d = 2
				x++
			} else {
				y--
			}
		case 2:
			if x == n-1 || matrix[y][x+1] > 0 {
				d = 3
				y++
			} else {
				x++
			}
		case 3:
			if y == n-1 || matrix[y+1][x] > 0 {
				d = 0
				x--
			} else {
				y++
			}
		}
		v++
		if v > n*n {
			break
		}
	}
	return matrix
}
