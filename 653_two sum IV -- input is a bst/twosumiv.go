package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type TreeNode = common.TreeNode

func main() {
	root := common.MakeTree([]int{5, 3, 6, 2, 4, -1, 7})
	k := 28
	fmt.Printf("Output: %v\n", findTarget(root, k))
}

func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	vals := make(map[int]int)
	var f func(t *TreeNode, v int) bool
	f = func(t *TreeNode, v int) bool {
		if t == nil {
			return false
		}
		if !f(t.Left, v) {
			if _, ok := vals[v-t.Val]; ok {
				return true
			}
			vals[t.Val] = 1
			return f(t.Right, v)
		}
		return true
	}
	return f(root, k)
}
