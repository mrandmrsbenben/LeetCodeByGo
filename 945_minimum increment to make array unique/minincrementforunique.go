package main

import (
	"fmt"
	"sort"
)

func main() {
	A := []int{1, 2, 2}
	// A := []int{3, 2, 1, 2, 1, 7} //1,1,2,2,3,7
	fmt.Println("Output: ", minIncrementForUnique(A))
}

//执行用时 :80 ms, 在所有 Go 提交中击败了57.69%的用户
//内存消耗 :6.4 MB, 在所有 Go 提交中击败了100.00%的用户
func minIncrementForUnique(A []int) int {
	sort.Ints(A)
	res := 0
	for i := 1; i < len(A); i++ {
		if A[i] <= A[i-1] {
			res += A[i-1] + 1 - A[i]
			A[i] = A[i-1] + 1
		}
	}
	return res
}

// Time Limit Exceeded
func minIncrementForUniqueV0(A []int) int {
	m := make(map[int]int)
	res := 0
	for _, n := range A {
		if _, ok := m[n]; ok {
			k := n
			for ok {
				res++
				k++
				_, ok = m[k]
			}
			m[k] = 1
		} else {
			m[n] = 1
		}
	}
	return res
}
