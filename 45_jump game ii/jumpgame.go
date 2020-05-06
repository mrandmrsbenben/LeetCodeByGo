package main

import "fmt"

func main() {
	// nums := []int{2, 3, 1, 1, 4}
	nums := []int{1, 2, 1, 1, 1}
	// nums := []int{3, 2, 1}
	// nums := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 0}
	fmt.Println("Output: ", jump(nums))
}

//执行用时 :8 ms, 在所有 Go 提交中击败了95.93%的用户
//内存消耗 :4.2 MB, 在所有 Go 提交中击败了100.00%的用户
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	steps := 0
	i := 0
	maxV, maxI := 0, 0
	for i < len(nums) {
		if nums[i] >= len(nums)-1-i {
			steps++
			break
		}
		maxV = 0
		for j := 1; j <= nums[i] && i+j < len(nums); j++ {
			if nums[i+j]+j >= maxV {
				maxV, maxI = nums[i+j]+j, i+j
			}
		}
		steps++
		i = maxI
		if i == len(nums)-1 {
			break
		}
	}
	return steps
}

//执行用时 :304 ms, 在所有 Go 提交中击败了23.63%的用户
// 内存消耗 :4.7 MB, 在所有 Go 提交中击败了100.00%的用户
func jumpV2(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	dp := make([]int, len(nums))
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] >= len(nums)-1-i {
			dp[i] = 1
		} else {
			min := len(nums)
			for j := 1; j <= nums[i]; j++ {
				if dp[i+j] == 1 {
					min = 1
					break
				} else if min > dp[i+j] {
					min = dp[i+j]
				}
			}
			dp[i] = min + 1
		}
	}
	return dp[0]
}

//执行用时 :1340 ms, 在所有 Go 提交中击败了5.70%的用户
//内存消耗 :6.9 MB, 在所有 Go 提交中击败了100.00%的用户
func jumpV1(nums []int) int {
	dp := make([]int, len(nums))

	var nextStep func(int) int
	nextStep = func(index int) int {
		if index >= len(nums)-1 {
			return 0
		} else if dp[index] > 0 {
			return dp[index]
		}
		min := len(nums)
		var ns int
		for i := 1; i <= nums[index]; i++ {
			ns = nextStep(index + i)
			if ns < min {
				min = ns
			}
			if min == 0 {
				break
			}
		}
		dp[index] = min + 1
		return min + 1
	}

	return nextStep(0)
}
