package main

import "fmt"

func main() {
	A := []int{8, 1, 5, 2, 6}
	// A := []int{3, 7, 2, 3}
	// A := []int{7, 2, 6, 6, 9, 4, 3}
	fmt.Println("Output: ", maxScoreSightseeingPair(A))
}

//执行用时 :60 ms, 在所有 Go 提交中击败了96.77%的用户
//内存消耗 :6.8 MB, 在所有 Go 提交中击败了100.00%的用户
func maxScoreSightseeingPair(A []int) int {
	max, val := 0, 0
	l, r := 0, 1
	for r < len(A) {
		val = A[l] + A[r] + l - r
		if val > max {
			max = val
		}
		if A[l]-A[r] <= r-l {
			l = r
		}
		r++
	}
	return max
}
