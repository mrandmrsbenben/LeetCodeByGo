package main

import (
	"fmt"
	"sort"
)

func main() {
	// a := 1
	// b := 2
	// c := 5
	// a, b, c := 4, 3, 2
	a, b, c := 3, 5, 1
	fmt.Println("Output:", numMovesStones(a, b, c))
}

//执行用时 :0 ms, 在所有Go提交中击败了100.00%的用户
//内存消耗 :2 MB, 在所有Go提交中击败了100.00%的用户
func numMovesStones(a int, b int, c int) []int {
	min, max := 0, 0
	nums := []int{a, b, c}
	sort.Ints(nums)
	if nums[1]-nums[0] == 1 && nums[2]-nums[1] == 1 {
		return []int{0, 0}
	}
	// find min
	atob := nums[1] - nums[0]
	btoc := nums[2] - nums[1]
	if atob <= 2 || btoc <= 2 {
		min = 1
	} else {
		min = 2
	}
	// find max
	max = atob - 1 + btoc - 1
	return []int{min, max}
}
