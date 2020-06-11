package main

import "fmt"

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println("Output: ", nums)
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2.1 MB, 在所有 Go 提交中击败了100.00%的用户
func sortColors(nums []int) {
	ridx, bidx := 0, len(nums)-1
	i := 0
	for i < len(nums) {
		if nums[i] == 2 && i < bidx {
			nums[i], nums[bidx] = nums[bidx], nums[i]
			bidx--
		} else if nums[i] == 0 && i > ridx {
			nums[i], nums[ridx] = nums[ridx], nums[i]
			ridx++
		} else {
			i++
		}
	}
}
