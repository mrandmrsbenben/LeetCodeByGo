package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	// nums := []int{2, 1, 3, 2, 3}
	// nums := []int{2, 4, -2, -3}
	// nums := []int{1, 2, 1, 2, 1, 2, 1, 2, 1, 2}
	// nums := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	// 	0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	// 	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, -1, -1, -1, -1, 3}
	fmt.Println("Output:", increasingTriplet(nums))
}

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	i, j, k := 0, -1, len(nums)-1
	for n := range nums {
		if nums[len(nums)-1-n] > nums[k] {
			k = len(nums) - 1 - n
		}
		if nums[n] < nums[i] {
			i = n
			j = -1
		} else if j == -1 && nums[n] > nums[i] {
			j = n
		} else if j > i && nums[n] > nums[i] && nums[n] < nums[j] {
			j = n
		} else if j > i && nums[n] > nums[j] {
			return true
		}
		if i < j && j < k && nums[i] < nums[j] && nums[j] < nums[k] {
			return true
		}
	}
	return false
}
