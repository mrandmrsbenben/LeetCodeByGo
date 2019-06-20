package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type TreeNode = common.TreeNode

func main() {
	root := makeTree()
	fmt.Println(searchBST(root, 5))
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	node := searchBST(root.Left, val)
	if node == nil {
		node = searchBST(root.Right, val)
	}
	return node
}

func makeTree() *TreeNode {
	t := &TreeNode{2, nil, nil}
	t.Left = &TreeNode{1, nil, nil}
	t.Right = &TreeNode{3, nil, nil}
	t.Left.Right = &TreeNode{4, nil, nil}
	t.Right.Right = &TreeNode{7, nil, nil}
	return t
}
