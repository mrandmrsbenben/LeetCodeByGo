package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type TreeNode = common.TreeNode

func main() {
	// vals := []int{10, 5, -3, 3, 2, -1, 11, 3, -2, -1, 1}
	// sum := 8
	// vals := []int{0, 1, 1, 0, 0}
	// sum := 1
	vals := []int{-8, -2, 8, -1, -1, 8, 2, -1, -1, -1, -1, -1, -1, -1, -2}
	sum := -2
	root := common.MakeTree(vals)
	fmt.Printf("Output: %d\n", pathSum(root, sum))
}

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	cnt := 0
	m := make(map[int]int)
	var findPath func(t *TreeNode, target map[int]int)
	findPath = func(t *TreeNode, target map[int]int) {
		if t == nil {
			return
		}
		if t.Val == sum {
			cnt = cnt + 1
		}
		if _, ok := target[t.Val]; ok {
			cnt = cnt + target[t.Val]
		}
		mp := make(map[int]int)
		if _, ok := target[sum]; ok {
			mp[sum-t.Val] = target[sum]
		} else {
			mp[sum-t.Val] = 1
		}
		for k, v1 := range target {
			if v2, ok := mp[k-t.Val]; ok {
				mp[k-t.Val] = 1 + v2
			} else {
				mp[k-t.Val] = v1
			}
		}
		findPath(t.Left, mp)
		findPath(t.Right, mp)
	}
	findPath(root, m)
	return cnt
}
