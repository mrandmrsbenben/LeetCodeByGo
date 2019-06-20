package main

import "fmt"

func main() {
	// nums := []int{1, 12, -5, -6, 50, 3}
	// k := 4
	nums := []int{0, 1, 1, 3, 3}
	k := 4
	fmt.Println("Output:", findMaxAverage(nums, k))
}

//执行用时 :200 ms, 在所有Go提交中击败了84.44%的用户
//内存消耗 :6.7 MB, 在所有Go提交中击败了100.00%的用户
func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	max := sum
	for i := 1; i <= len(nums)-k; i++ {
		sum = sum - nums[i-1] + nums[i+k-1]
		if sum > max {
			max = sum
		}
	}
	return float64(max) / float64(k)
}
