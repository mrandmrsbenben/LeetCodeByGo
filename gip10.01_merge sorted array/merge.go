package main

import (
	"fmt"
	"sort"
)

func main() {
	// A := []int{1, 2, 3, 0, 0, 0}
	// m := 3
	// B := []int{2, 5, 6}
	// n := 3
	// A := []int{0}
	// m := 0
	// B := []int{1}
	// n := 1
	A := []int{4, 0, 0, 0, 0, 0}
	m := 1
	B := []int{1, 2, 3, 5, 6}
	n := 5
	merge(A, m, B, n)
	fmt.Println("Output: ", A)
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2.3 MB, 在所有 Go 提交中击败了100.00%的用户
func merge(A []int, m int, B []int, n int) {
	sort.Ints(append(A[0:m], B...))
}
