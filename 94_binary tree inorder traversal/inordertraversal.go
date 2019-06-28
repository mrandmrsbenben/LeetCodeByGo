package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type TreeNode = common.TreeNode

func main() {
	vals := []int{1, -1, 2, -1, -1, 3}
	fmt.Println("Output:", inorderTraversal(common.MakeTree(vals)))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2.2 MB, 在所有 Go 提交中击败了59.75%的用户
func inorderTraversal(root *TreeNode) []int {
	vals := make([]int, 0)
	nodes := make([]*TreeNode, 0)
	for {
		for root != nil {
			nodes = append(nodes, root)
			root = root.Left
		}
		if len(nodes) == 0 {
			break
		}
		root = nodes[len(nodes)-1]
		nodes = nodes[0 : len(nodes)-1]
		vals = append(vals, root.Val)
		root = root.Right
	}
	return vals
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2.2 MB, 在所有 Go 提交中击败了44.65%的用户
func inorderTraversalV1(root *TreeNode) []int {
	vals := make([]int, 0)
	if root == nil {
		return vals
	}
	vals = inorderTraversal(root.Left)
	vals = append(vals, root.Val)
	vals = append(vals, inorderTraversal(root.Right)...)
	return vals
}
