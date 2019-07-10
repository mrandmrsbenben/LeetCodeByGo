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

//执行用时 :4 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :7.1 MB, 在所有 Go 提交中击败了17.86%的用户
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	var node, buf *TreeNode
	node = root
	for node != nil {
		if node.Left != nil {
			buf = node.Left
			for buf.Right != nil {
				buf = buf.Right
			}
			buf.Right, node.Right = node.Right, node.Left
			node.Left = nil
		}
		node = node.Right
	}
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :6.8 MB, 在所有 Go 提交中击败了46.43%的用户
func flattenV2(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	r := root.Right
	root.Right, root.Left = root.Left, nil
	for root.Right != nil {
		root = root.Right
	}
	root.Right = r
}

//执行用时 :8 ms, 在所有 Go 提交中击败了91.78%的用户
//内存消耗 :6.6 MB, 在所有 Go 提交中击败了71.43%的用户
func flattenV1(root *TreeNode) {
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
