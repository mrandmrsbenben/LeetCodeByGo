package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type TreeNode = common.TreeNode

func main() {
	vals := []int{90, 69, -1, 49, 89, -1, 52}
	fmt.Printf("Output :%d\n", minDiffInBST(common.MakeTree(vals)))
}

func minDiffInBST(root *TreeNode) int {
	min := 0
	vals := make([]int, 0)
	var f func(r *TreeNode)
	f = func(r *TreeNode) {
		if r == nil {
			return
		}
		if r.Left != nil {
			f(r.Left)
		}
		if len(vals) > 0 && (min == 0 || r.Val-vals[len(vals)-1] < min) {
			min = r.Val - vals[len(vals)-1]
		}
		vals = append(vals, r.Val)
		if r.Right != nil {
			f(r.Right)
		}
	}
	f(root)
	return min
}
