package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type TreeNode = common.TreeNode

func main() {
	vals := []int{3, 9, 20, -1, -1, 15, 7}
	root := common.MakeTree(vals)
	fmt.Printf("Output: %v\n", sumOfLeftLeaves(root))
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	if root.Left != nil && root.Left.Left == nil &&
		root.Left.Right == nil {
		sum = sum + root.Left.Val
	}
	sum = sum + sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
	return sum
}
