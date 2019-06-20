package main

import (
	"fmt"
	"sort"
)

func main() {
	// A := []int{1, 2, 3, 4}
	// A := []int{2, 2, 2, 2}
	// A := []int{9, 2, 3, 5}
	// A := []int{0, 2, 4, 0}
	// A := []int{4, 4, 4, 0}
	// A := []int{1, 9, 6, 0}
	// A := []int{7, 9, 7, 0}
	// A := []int{2, 6, 6, 0}
	// A := []int{5, 5, 5, 5}
	A := []int{0, 0, 1, 0}
	fmt.Println("Output:", largestTimeFromDigits(A))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2 MB, 在所有 Go 提交中击败了100.00%的用户
func largestTimeFromDigits(A []int) string {
	sort.Ints(A)
	var h, m int
	for i := len(A) - 1; i >= 0; i-- {
		if A[i] > 2 {
			continue
		}
		for j := len(A) - 1; j >= 0; j-- {
			h = A[i]*10 + A[j]
			if j == i || h > 23 {
				continue
			}
			m = getMinutes(A, i, j)
			if m != -1 {
				return fmt.Sprintf("%02d", h) + ":" + fmt.Sprintf("%02d", m)
			}
		}
	}
	return ""
}

func getMinutes(A []int, i, j int) int {
	m1, m2 := -1, -1
	for k := len(A) - 1; k >= 0; k-- {
		if k == i || k == j {
			continue
		}
		if m1 == -1 {
			m1 = A[k]
		} else if m2 == -1 {
			m2 = A[k]
			break
		}
	}
	if m1*10+m2 < 60 {
		return m1*10 + m2
	} else if m2*10+m1 < 60 {
		return m2*10 + m1
	}
	return -1
}
