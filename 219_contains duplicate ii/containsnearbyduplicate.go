package main

import "fmt"

func main() {
	// nums := []int{1, 2, 3, 1}
	// k := 3
	nums := []int{1, 0, 1, 1}
	k := 1
	// nums := []int{1, 2, 3, 1, 2, 3}
	// k := 2
	// nums := []int{99, 99}
	// k := 2
	fmt.Println("Output:", containsNearbyDuplicate(nums, k))
}

//执行用时 :20 ms, 在所有Go提交中击败了98.62%的用户
//内存消耗 :8.8 MB, 在所有Go提交中击败了39.02%的用户
func containsNearbyDuplicate(nums []int, k int) bool {
	nm := make(map[int]int)
	for i, n := range nums {
		if index, ok := nm[n]; ok {
			if i-index <= k {
				return true
			}
		}
		nm[n] = i
	}
	return false
}

//执行用时 :1008 ms, 在所有Go提交中击败了21.38%的用户
//内存消耗 :5.2 MB, 在所有Go提交中击败了87.81%的用户
func containsNearbyDuplicateV1(nums []int, k int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= i+k && j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}
