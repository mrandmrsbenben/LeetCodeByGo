package main

import "fmt"

func main() {
	nums := []int{0, 2, 3}
	// nums := []int{0, 1, 2, 3, 4, 5, 7, 8}
	fmt.Println("Output: ", missingNumber(nums))
}

//执行用时 :16 ms, 在所有 Go 提交中击败了94.98%的用户
//内存消耗 :6 MB, 在所有 Go 提交中击败了100.00%的用户
func missingNumber(nums []int) int {
	if nums[0] != 0 {
		return 0
	} else if nums[len(nums)-1] == len(nums)-1 {
		return len(nums)
	}
	l, r := 0, len(nums)-1
	var i int
	for l < r {
		i = l + (r-l)/2
		if nums[i] == i {
			l = i + 1
		} else {
			r = i
		}
	}
	return l
}
