package main

import "fmt"

func main() {
	nums := []int{2, 0, 2}
	k := 2
	fmt.Println("Output: ", subarraySum(nums, k))
}

//执行用时 :20 ms, 在所有 Go 提交中击败了84.88%的用户
//内存消耗 :6.3 MB, 在所有 Go 提交中击败了100.00%的用户
func subarraySum(nums []int, k int) int {
	cnt, sum := 0, 0
	m := make(map[int]int)
	m[0] = 1

	for _, n := range nums {
		sum += n
		if m[sum-k] > 0 {
			cnt += m[sum-k]
		}
		m[sum]++
	}
	return cnt
}

//执行用时 :280 ms, 在所有 Go 提交中击败了5.86%的用户
//内存消耗 :5.9 MB, 在所有 Go 提交中击败了100.00%的用户
func subarraySumV1(nums []int, k int) int {
	cnt := 0
	sum := make([]int, len(nums)+1)

	for i := range nums {
		sum[i+1] = sum[i] + nums[i]
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= len(nums); j++ {
			if sum[j]-sum[i] == k {
				cnt++
			}
		}
	}
	return cnt
}
