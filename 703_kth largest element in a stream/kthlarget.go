package main

import (
	"fmt"
	"sort"
)

// 执行用时 : 424 ms, 在Kth Largest Element in a Stream的Go提交中击败了11.39% 的用户
// 内存消耗 : 62.7 MB, 在Kth Largest Element in a Stream的Go提交中击败了8.33% 的用户
type KthLargest struct {
	k    int
	nums []int
}

func Constructor(k int, nums []int) KthLargest {
	sort.Ints(nums)
	return KthLargest{k, nums}
}

func (this *KthLargest) Add(val int) int {
	index := sort.SearchInts(this.nums, val)
	if index == 0 {
		this.nums = append([]int{val}, this.nums...)
	} else if index == len(this.nums) {
		this.nums = append(this.nums, val)
	} else {
		this.nums = append(this.nums[0:index], append([]int{val}, this.nums[index:]...)...)
	}
	return this.nums[len(this.nums)-this.k]
}

func main() {
	k := 3
	nums := []int{4, 5, 8, 2}
	obj := Constructor(k, nums)
	fmt.Println(obj.Add(3))
	fmt.Println(obj.Add(5))
	fmt.Println(obj.Add(10))
	fmt.Println(obj.Add(9))
	fmt.Println(obj.Add(4))
}
