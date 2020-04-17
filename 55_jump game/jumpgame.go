package main

import "fmt"

func main() {
	// nums := []int{2, 3, 1, 1, 4}
	nums := []int{3, 2, 2, 0, 4}
	// nums := []int{3, 0, 0, 1, 0}
	fmt.Println("Output: ", canJump(nums))
}

//执行用时 :8 ms, 在所有 Go 提交中击败了97.02%的用户
//内存消耗 :4.2 MB, 在所有 Go 提交中击败了100.00%的用户
func canJump(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			canJ := false
			for j := i - 1; j >= 0 && !canJ; j-- {
				if nums[j] > i-j {
					canJ = true
				}
			}
			if !canJ {
				return false
			}
		}
	}
	return true
}
