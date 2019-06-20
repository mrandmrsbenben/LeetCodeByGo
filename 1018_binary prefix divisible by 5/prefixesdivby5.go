package main

import "fmt"

func main() {
	A := []int{0, 1, 1}
	// A := []int{1, 1, 1}
	// A := []int{0, 1, 1, 1, 1, 1}
	// A := []int{1, 1, 1, 0, 1}
	// A := []int{1, 1, 0, 0, 0, 1, 0, 0, 1}
	fmt.Printf("Output: %v\n", prefixesDivBy5(A))
}

// 执行用时 : 752 ms, 在Binary Prefix Divisible By 5的Go提交中击败了98.36% 的用户
// 内存消耗 : 17.8 MB, 在Binary Prefix Divisible By 5的Go提交中击败了28.57% 的用户
func prefixesDivBy5(A []int) []bool {
	b := make([]bool, len(A))
	add := 0
	for i := range A {
		if A[i] == 1 {
			add = add*2 + 1
		} else {
			add = add * 2
		}
		if add%5 == 0 {
			b[i] = true
		} else {
			b[i] = false
		}
		add %= 10
	}
	return b
}
