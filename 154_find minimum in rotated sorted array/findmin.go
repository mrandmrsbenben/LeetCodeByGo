package main

import "fmt"

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println("Output: ", findMin(nums))
}

//执行用时 :4 ms, 在所有 Go 提交中击败了92.25%的用户
//内存消耗 :3.1 MB, 在所有 Go 提交中击败了55.77%的用户
func findMin(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	} else if l == 2 {
		if nums[1] < nums[0] {
			return nums[1]
		}
		return nums[0]
	} else if nums[0] < nums[l-1] {
		return nums[0]
	}

	leftN := findMin(nums[0 : l/2])
	rightN := findMin(nums[l/2:])
	if leftN < rightN {
		return leftN
	}
	return rightN
}
