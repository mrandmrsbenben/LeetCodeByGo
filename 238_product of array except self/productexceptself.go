package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println("Output: ", productExceptSelf(nums))
}

//执行用时 :12 ms, 在所有 Go 提交中击败了94.15%的用户
//内存消耗 :6.7 MB, 在所有 Go 提交中击败了100.00%的用户
func productExceptSelf(nums []int) []int {
	lp, rp := make([]int, len(nums)), make([]int, len(nums))
	lp[0] = 1
	for i := 1; i < len(nums); i++ {
		lp[i] = lp[i-1] * nums[i-1]
	}
	rp[len(rp)-1] = 1
	for j := len(nums) - 2; j >= 0; j-- {
		rp[j] = rp[j+1] * nums[j+1]
	}
	for i := range nums {
		nums[i] = lp[i] * rp[i]
	}
	return nums
}
