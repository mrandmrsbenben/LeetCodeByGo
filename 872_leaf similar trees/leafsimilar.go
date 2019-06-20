package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type TreeNode = common.TreeNode

func main() {
	root1 := common.MakeTree([]int{3, 5, 1, 6, 2, 9, 8, -1, -1, 7, 4})
	root2 := common.MakeTree([]int{3, 5, 1, 6, 7, 4, 2, -1, -1, -1, -1, -1, -1, 9, 8})
	// root1 := common.MakeTree([]int{3})
	// root2 := common.MakeTree([]int{1})
	fmt.Printf("Output :%v\n", leafSimilar(root1, root2))
}
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaf := make([]int, 0)
	var f func(root *TreeNode)
	f = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			leaf = append(leaf, root.Val)
			return
		}
		f(root.Left)
		f(root.Right)
	}
	f(root1)
	leaf1 := make([]int, len(leaf))
	copy(leaf1, leaf)
	leaf = make([]int, 0)
	f(root2)
	if len(leaf1) != len(leaf) {
		return false
	}
	for i := range leaf1 {
		if leaf1[i] != leaf[i] {
			return false
		}
	}
	return true
}
