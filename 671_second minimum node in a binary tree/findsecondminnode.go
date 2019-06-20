package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
	"sort"
)

type TreeNode = common.TreeNode

func main() {
	vals := []int{1, 1, 3, 1, 1, 3, 4, 3, 1, 1, 1, 3, 8, 4, 8, 3, 3, 1, 6, 2, 1}
	fmt.Printf("Output: %d\n", findSecondMinimumValue(common.MakeTree(vals)))
}

func findSecondMinimumValue(root *TreeNode) int {
	if root == nil || root.Left == nil || root.Right == nil {
		return -1
	}
	vals := []int{root.Val, root.Left.Val, root.Right.Val}
	var getNodeVals func(node *TreeNode)
	getNodeVals = func(node *TreeNode) {
		if node.Left != nil && node.Right != nil {
			vals = append(vals, node.Left.Val)
			vals = append(vals, node.Right.Val)
			getNodeVals(node.Left)
			getNodeVals(node.Right)
		}
	}
	getNodeVals(root.Left)
	getNodeVals(root.Right)
	sort.Ints(vals)
	if vals[0] == vals[len(vals)-1] {
		return -1
	}
	for _, v := range vals {
		if v > vals[0] {
			return v
		}
	}
	return -1
}
