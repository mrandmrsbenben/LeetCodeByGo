package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(numSquares(26))
}

func numSquares(n int) int {
	cnt := 0
	numsMap := make(map[int]int)
	nums := []int{n}
	numsMap[n] = 1
	var l, x, y int
	for len(nums) > 0 {
		l = len(nums)
		cnt++
		for i := range nums {
			x = int(math.Sqrt(float64(nums[i])))
			for j := x; j > 0; j-- {
				if j*j == nums[i] {
					return cnt
				}
				y = nums[i] - j*j
				if _, ok := numsMap[y]; !ok {
					numsMap[y] = 1
					nums = append(nums, y)
				}
			}
		}
		nums = nums[l:]
	}
	return cnt
}
