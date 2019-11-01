package main

import "fmt"

func main() {
	// chips := []int{1, 2, 3}
	chips := []int{2, 2, 2, 3, 3}
	fmt.Println("Output: ", minCostToMoveChips(chips))
}

//执行用时 :0 ms, 在所有 golang 提交中击败了100.00%的用户
//内存消耗 :2.1 MB, 在所有 golang 提交中击败了100.00%的用户
func minCostToMoveChips(chips []int) int {
	var c0, c1 int

	for _, c := range chips {
		if c%2 == 0 {
			c0++
		} else {
			c1++
		}
	}
	if c0 < c1 {
		return c0
	}
	return c1
}
