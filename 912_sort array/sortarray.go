package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	nums := []int{5, 1, 1, 2, 0, 0}
	fmt.Println("Output: ", sortArray(nums))
}

//执行用时 :36 ms, 在所有 Go 提交中击败了27.55%的用户
//内存消耗 :6.9 MB, 在所有 Go 提交中击败了5.13%的用户
type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func sortArray(nums []int) []int {
	h := &intHeap{}
	for i := range nums {
		heap.Push(h, nums[i])
	}
	for i := range nums {
		nums[i] = heap.Pop(h).(int)
	}
	return nums
}

//执行用时 :28 ms, 在所有 Go 提交中击败了58.33%的用户
//内存消耗 :6.9 MB, 在所有 Go 提交中击败了5.13%的用户
func sortArrayV2(nums []int) []int {
	var divide func(n []int) []int
	divide = func(n []int) []int {
		l := len(n)
		if len(n) < 2 {
			return n
		}
		return merge(divide(n[0:l/2]), divide(n[l/2:]))
	}
	return divide(nums)
}

func merge(l, r []int) []int {
	ll, lr := len(l), len(r)
	res := make([]int, ll+lr)
	var i, j, k int
	for i < ll && j < lr {
		if l[i] < r[j] {
			res[k] = l[i]
			i++
		} else {
			res[k] = r[j]
			j++
		}
		k++
	}
	for ; i < ll; i++ {
		res[k] = l[i]
		k++
	}
	for ; j < lr; j++ {
		res[k] = r[j]
		k++
	}
	return res
}

//执行用时 :28 ms, 在所有 Go 提交中击败了58.33%的用户
//内存消耗 :6.5 MB, 在所有 Go 提交中击败了5.13%的用户
func sortArrayV1(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return nums
}
