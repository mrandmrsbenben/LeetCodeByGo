package main

import (
	"fmt"
	"sort"
)

func main() {
	// nums := []int{1, 2, 3, 4}
	nums := []int{0, 3, 5, 7, 9}
	fmt.Println("Output:", subsets(nums))
}

//执行用时 :4 ms, 在所有 Go 提交中击败了93.43%的用户
//内存消耗 :7.2 MB, 在所有 Go 提交中击败了21.74%的用户
func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{[]int{}}
	}
	sort.Ints(nums)
	subs := make([][]int, 2)
	subs[0] = []int{}
	subs[1] = nums
	if len(nums) == 1 {
		return subs
	}
	sets := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		sets[i] = []int{nums[i]}
	}
	subs = append(subs, sets...)
	for i := 2; i < len(nums); i++ {
		sets = getSubSetsByLen(nums, sets)
		subs = append(subs, sets...)
	}
	return subs
}

// get every length's subsets
func getSubSetsByLen(nums []int, subs [][]int) [][]int {
	length := combinsCount(len(subs[0])+1, len(nums))
	subsets := make([][]int, length)
	subsets[0] = make([]int, len(subs[0])+1)
	index := 0
	for i := range subs {
		for j := range nums {
			if nums[j] > subs[i][len(subs[i])-1] {
				subsets[index] = make([]int, len(subs[0])+1)
				copy(subsets[index], subs[i])
				subsets[index][len(subs[i])] = nums[j]
				index++
			}
		}
	}
	return subsets
}

// get count of combinations
func combinsCount(m, n int) int {
	switch m {
	case 2:
		return n * (n - 1) / 2
	case 3:
		return n * (n - 1) * (n - 2) / 6
	case n:
		return 1
	default:
		return combinsCount(m-1, n-1) + combinsCount(m, n-1)
	}
}
