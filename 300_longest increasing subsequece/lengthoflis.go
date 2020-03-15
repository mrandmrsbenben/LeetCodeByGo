package main

import "fmt"

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println("Output: ", lengthOfLIS(nums))
}

//执行用时 :8 ms, 在所有 Go 提交中击败了70.04%的用户
//内存消耗 :2.4 MB, 在所有 Go 提交中击败了59.06%的用户
func lengthOfLIS(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	dp := make([]int, l)
	dp[0] = 1
	ans, max := 1, 0
	for i := 1; i < l; i++ {
		max = 0
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && max < dp[j] {
				max = dp[j]
			}
		}
		dp[i] = max + 1
		if ans < dp[i] {
			ans = dp[i]
		}
	}

	return ans
}
