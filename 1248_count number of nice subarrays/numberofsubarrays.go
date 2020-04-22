package main

import "fmt"

func main() {
	nums := []int{1, 1, 2, 1, 1}
	k := 3
	// nums := []int{2, 4, 6}
	// k := 1
	// nums := []int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}
	// k := 2
	// nums := []int{1, 1, 1, 1, 1}
	// k := 1
	// nums := []int{2, 1, 1}
	// k := 1
	// nums := []int{91473, 45388, 24720, 35841, 29648, 77363, 86290, 58032, 53752, 87188, 34428, 85343, 19801, 73201}
	// k := 4
	fmt.Println("Output: ", numberOfSubarrays(nums, k))
}

//执行用时 :152 ms, 在所有 Go 提交中击败了47.83%的用户
//内存消耗 :7.3 MB, 在所有 Go 提交中击败了100.00%的用户
func numberOfSubarrays(nums []int, k int) int {
	cnt := 0
	l, r := 0, k-1
	c := 0
	for i := 0; i < k; i++ {
		if nums[i]%2 == 1 {
			c++
		}
	}
	for r < len(nums) && r-l >= k-1 {
		if c == k {
			cnt++
			for i := r + 1; i < len(nums); i++ {
				if nums[i]%2 == 1 {
					break
				}
				cnt++
			}
			if nums[l]%2 == 1 {
				c--
			}
			l++
			if l > r {
				r = l
				if r < len(nums) && nums[r]%2 == 1 {
					c++
				}
			}
		} else {
			if r == len(nums)-1 {
				l++
				if l < len(nums) && nums[l]%2 == 1 {
					c--
				}
			} else {
				r++
				if r < len(nums) && nums[r]%2 == 1 {
					c++
				}
			}
		}
	}
	return cnt
}
