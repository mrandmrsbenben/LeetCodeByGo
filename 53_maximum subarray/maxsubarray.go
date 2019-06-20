package main

import (
	"fmt"
)

func main() {
	nums := []int{-2, 1, -3, -1, 2, 1, -5, 4}
	// nums := []int{-2, -1}
	fmt.Printf("Output: %d\n", maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	sum := 0
	allsum := nums[0]
	for _, n := range nums {
		if sum < 0 {
			sum = n
		} else {
			sum = sum + n
		}
		if sum > allsum {
			allsum = sum
		}
	}
	return allsum
}

func maxSubArraySlow(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}
	maxLSum := func(arr []int) int {
		maxl := 0
		suml := 0
		for i := len(arr) - 1; i >= 0; i-- {
			suml = suml + arr[i]
			if suml > maxl {
				maxl = suml
			}
		}
		return maxl
	}
	maxRSum := func(arr []int) int {
		maxr := 0
		sumr := 0
		for _, n := range arr {
			sumr = sumr + n
			if sumr > maxr {
				maxr = sumr
			}
		}
		return maxr
	}
	sum, max := 0, nums[0]
	for i := range nums {
		if nums[i] < 0 {
			if nums[i] > max {
				max = nums[i]
			}
			continue
		}
		sum = 0
		if i > 0 {
			sum = sum + maxLSum(nums[0:i])
		}
		sum = sum + nums[i]
		if i < len(nums)-1 {
			sum = sum + maxRSum(nums[i+1:])
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
