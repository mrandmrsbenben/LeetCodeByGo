package main

import "fmt"

func main() {
	nums := []int{1, 1, 2, 3, 3, 4, 4, 8, 8}
	// nums := []int{3, 3, 7, 7, 10, 11, 11}
	// nums := []int{1, 1, 2}
	fmt.Println("Output:", singleNonDuplicate(nums))
}

//执行用时 :12 ms, 在所有 Go 提交中击败了61.76%的用户
//内存消耗 :4.3 MB, 在所有 Go 提交中击败了9.09%的用户
func singleNonDuplicate(nums []int) int {
	var i int
	for {
		if len(nums) <= 2 {
			return nums[0]
		}
		i = len(nums) / 2
		if i%2 == 1 {
			i++
		}
		if nums[i-1] != nums[i-2] {
			nums = nums[0:i]
		} else {
			nums = nums[i:]
		}
	}
}

//执行用时 :12 ms, 在所有 Go 提交中击败了61.76%的用户
//内存消耗 :4.3 MB, 在所有 Go 提交中击败了9.09%的用户
func singleNonDuplicateV1(nums []int) int {
	n := nums[0]
	for i := 1; i < len(nums); i++ {
		n ^= nums[i]
	}
	return n
}
