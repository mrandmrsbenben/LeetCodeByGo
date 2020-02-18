package main

import (
	"fmt"
	"sort"
)

func main() {
	// arr := []int{40, 10, 20, 30}
	arr := []int{100, 100, 100}
	// arr := []int{37, 12, 28, 9, 100, 56, 80, 5, 12}
	fmt.Println("Output: ", arrayRankTransform(arr))
}

//执行用时 :96 ms, 在所有 Go 提交中击败了34.78%的用户
//内存消耗 :8.8 MB, 在所有 Go 提交中击败了100.00%的用户
type num struct {
	Num   int
	Index int
}

func arrayRankTransform(arr []int) []int {
	rank := make([]int, len(arr))
	l := len(arr)
	if l == 0 {
		return rank
	}

	nums := make([]*num, l)
	for i := range arr {
		nums[i] = &num{arr[i], i}
	}
	sort.Slice(nums, func(i, j int) bool { return nums[i].Num < nums[j].Num })

	cnt := 0
	rank[nums[0].Index] = 1
	for i := 1; i < l; i++ {
		if nums[i].Num == nums[i-1].Num {
			rank[nums[i].Index] = rank[nums[i-1].Index]
			cnt++
		} else {
			rank[nums[i].Index] = i + 1 - cnt
		}
	}

	return rank
}

//执行用时 :92 ms, 在所有 Go 提交中击败了52.17%的用户
//内存消耗 :10 MB, 在所有 Go 提交中击败了100.00%的用户
func arrayRankTransformV1(arr []int) []int {
	rank := make([]int, len(arr))
	if len(arr) == 0 {
		return rank
	}

	copy(rank, arr)
	sort.Ints(arr)
	cnt := 1
	mRank := make(map[int]int)
	mRank[arr[0]] = cnt
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			cnt++
		}
		mRank[arr[i]] = cnt
	}

	for i := range rank {
		rank[i] = mRank[rank[i]]
	}

	return rank
}
