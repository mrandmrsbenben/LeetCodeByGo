package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type TreeNode = common.TreeNode

func main() {
	// vals := []int{1, 1, 1, 1, 1, -1, 1}
	vals := []int{2, 2, 2, 5, 2}
	root := common.MakeTree(vals)
	fmt.Println(vals, isUnivalTree(root))
}

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if (root.Left != nil && root.Val != root.Left.Val) ||
		(root.Right != nil && root.Val != root.Right.Val) {
		return false
	}
	return isUnivalTree(root.Left) && isUnivalTree(root.Right)
}
