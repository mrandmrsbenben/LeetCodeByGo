package main

import (
	"fmt"
	"strconv"
)

func main() {
	// nums := []int{0, 1, 2, 4, 5, 7}
	// nums := []int{0, 2, 3, 4, 6, 8, 9}
	nums := []int{0}
	fmt.Println("Output: ", summaryRanges(nums))
}

//执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗：1.9 MB, 在所有 Go 提交中击败了75.86%的用户
func summaryRanges(nums []int) []string {
	res := make([]string, 0)
	if len(nums) == 0 {
		return res
	}
	begin, end := nums[0], nums[0]+1
	for i := 1; i < len(nums); i++ {
		if nums[i] == end {
			end++
			continue
		}

		if begin == end-1 {
			res = append(res, strconv.Itoa(begin))
		} else {
			res = append(res, strconv.Itoa(begin)+"->"+strconv.Itoa(end-1))
		}
		begin, end = nums[i], nums[i]+1
	}

	if begin+1 == end {
		res = append(res, strconv.Itoa(begin))
	} else {
		res = append(res, strconv.Itoa(begin)+"->"+strconv.Itoa(end-1))
	}
	return res
}
