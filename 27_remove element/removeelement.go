package main

import "fmt"

func main() {
	nums := []int{3, 2, 2, 3}
	nums = nums[0:]
	val := 3
	fmt.Printf("Output: %d, %v\n", removeElement(nums, val), nums)
}

func removeElement(nums []int, val int) int {
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == val {
			nums = append(nums[0:i], nums[i+1:]...)
		}
		fmt.Println(nums)
	}
	return len(nums)
}
