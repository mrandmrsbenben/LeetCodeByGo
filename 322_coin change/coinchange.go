package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	// coins := []int{1, 2, 5}
	// amount := 11
	// coins := []int{186, 419, 83, 408}
	// amount := 6249
	coins := []int{288, 160, 10, 249, 40, 77, 314, 429}
	amount := 9208
	fmt.Println("Output: ", coinChange(coins, amount))
}

//执行用时 :4 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2 MB, 在所有 Go 提交中击败了100.00%的用户
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	sort.Sort(sort.Reverse(sort.IntSlice(coins)))
	min := math.MaxInt64

	var change func(c []int, a int, cnt int)
	change = func(c []int, a int, cnt int) {
		x := a / c[0]
		y := a % c[0]
		if y == 0 {
			if min > x+cnt {
				min = x + cnt
			}
			return
		} else if len(c) == 1 {
			return
		}
		for i := 0; i <= x && x-i+cnt < min; i++ {
			change(c[1:], i*c[0]+y, x-i+cnt)
		}
	}
	change(coins, amount, 0)
	if min == math.MaxInt64 {
		return -1
	}
	return min
}
