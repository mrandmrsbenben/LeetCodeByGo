package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type TreeNode = common.TreeNode

func main() {
	vals := []int{1, 2, 5, 3, 4, -1, 6}
	root := common.MakeTree(vals)
	fmt.Println("Input :", root)
	flatten(root)
	fmt.Println("Output:", root)
}

//执行用时 :8 ms, 在所有 Go 提交中击败了91.78%的用户
//内存消耗 :6.6 MB, 在所有 Go 提交中击败了71.43%的用户
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	vals := make([]int, 0)
	var getVals func(node *TreeNode)
	getVals = func(node *TreeNode) {
		if node == nil {
			return
		}
		vals = append(vals, node.Val)
		getVals(node.Left)
		getVals(node.Right)
	}
	getVals(root)
	var node, right *TreeNode
	for i := len(vals) - 1; i > 0; i-- {
		node = &TreeNode{vals[i], nil, right}
		right = node
	}
	root.Left = nil
	root.Right = right
}
