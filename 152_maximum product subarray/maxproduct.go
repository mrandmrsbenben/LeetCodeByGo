package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 3, -2, 4}
	// nums := []int{-1, 0, -2}
	// nums := []int{-1, -2, -9, -6}
	// nums := []int{2, -5, -2, -4, 3}
	fmt.Println("Output: ", maxProduct(nums))
}

//执行用时 :4 ms, 在所有 Go 提交中击败了88.22%的用户
//内存消耗 :2.9 MB, 在所有 Go 提交中击败了50.00%的用户
func maxProduct(nums []int) int {

	dpMax, dpMin := make([]int, len(nums)), make([]int, len(nums))
	dpMax[0], dpMin[0] = nums[0], nums[0]
	res := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			if dpMax[i-1] > 0 {
				dpMax[i] = nums[i] * dpMax[i-1]
			} else if dpMax[i-1] <= 0 {
				dpMax[i] = nums[i]
			}
			if dpMin[i-1] > 0 {
				dpMin[i] = nums[i]
			} else if dpMin[i-1] <= 0 {
				dpMin[i] = nums[i] * dpMin[i-1]
			}
		} else if nums[i] < 0 {
			if dpMin[i-1] <= 0 {
				dpMax[i] = nums[i] * dpMin[i-1]
				if dpMax[i-1] <= 0 {
					dpMin[i] = nums[i]
				} else {
					dpMin[i] = nums[i] * dpMax[i-1]
				}
			} else {
				dpMin[i] = nums[i] * dpMax[i-1]
				dpMax[i] = nums[i]
			}
		} else {
			dpMax[i], dpMin[i] = 0, 0
		}
		if dpMax[i] > res {
			res = dpMax[i]
		}
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
