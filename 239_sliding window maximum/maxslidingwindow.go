package main

import "fmt"

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println("Output: ", maxSlidingWindow(nums, k))
}

//执行用时 :20 ms, 在所有 Go 提交中击败了84.60%的用户
//内存消耗 :6.9 MB, 在所有 Go 提交中击败了25.00%的用户
func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0)
	l, r := make([]int, len(nums)), make([]int, len(nums))

	var cnt, max int
	// from left to right
	for i := range nums {
		if i%k == 0 || nums[i] > max {
			max = nums[i]
		}
		l[i] = max
	}

	// max value from right to left
	max = nums[len(nums)-1]
	for j := len(nums) - 1; j >= 0; j-- {
		if (j+1)%k == 0 || nums[j] > max {
			max = nums[j]
		}
		cnt++
		r[j] = max
	}

	for i := 0; i+k-1 < len(nums); i++ {
		res = append(res, maxVal(r[i], l[i+k-1]))
	}
	return res
}

func maxVal(x, y int) int {
	if x > y {
		return x
	}
	return y
}
