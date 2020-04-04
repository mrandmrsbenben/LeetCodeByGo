package main

import "fmt"

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	// height := []int{4, 2, 3}
	// height := []int{5, 4, 1, 2}
	fmt.Println("Output: ", trap(height))
}

//执行用时 :4 ms, 在所有 Go 提交中击败了86.97%的用户
//内存消耗 :2.8 MB, 在所有 Go 提交中击败了29.63%的用户
func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	res := 0
	l, r := 0, len(height)-1
	lmax, rmax := 0, 0
	for l < r {
		if height[l] < height[r] {
			if height[l] >= lmax {
				lmax = height[l]
			} else {
				res += lmax - height[l]
			}
			l++
		} else {
			if height[r] >= rmax {
				rmax = height[r]
			} else {
				res += rmax - height[r]
			}
			r--
		}
	}
	return res
}
