package main

import (
	"fmt"
)

func main() {
	// nums := []int{4, 2, 1}
	// nums := []int{3, 4, 2, 3}
	nums := []int{2, 3, 3, 2, 4}
	fmt.Println("Output:", checkPossibility(nums))
}

//执行用时 :28 ms, 在所有 Go 提交中击败了98.28%的用户
//内存消耗 :6.4 MB, 在所有 Go 提交中击败了14.81%的用户
func checkPossibility(nums []int) bool {
	index := -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			index = i
			break
		}
	}
	if index == -1 {
		return true
	}
	buf := make([]int, index)
	copy(buf, nums[0:index])
	buf = append(buf, nums[index+1:]...)
	if isNonDecArray(buf) {
		return true
	}
	buf = make([]int, index+1)
	copy(buf, nums[0:index+1])
	buf = append(buf, nums[index+2:]...)
	return isNonDecArray(buf)
}

func isNonDecArray(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}
	return true
}
