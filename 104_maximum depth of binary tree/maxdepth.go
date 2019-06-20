package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type TreeNode = common.TreeNode

func main() {
	root := makeTree()
	fmt.Println(root)
	fmt.Printf("Depth:%d\n", maxDepth(root))
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 1
	leftD := maxDepth(root.Left)
	rightD := maxDepth(root.Right)
	if leftD > rightD {
		depth = depth + leftD
	} else {
		depth = depth + rightD
	}
	return depth
}

func makeTree() *TreeNode {
	t := &TreeNode{Val: 4, Left: nil, Right: nil}
	t.Left = &TreeNode{Val: 2, Left: nil, Right: nil}
	t.Right = &TreeNode{Val: 7, Left: nil, Right: nil}
	t.Left.Left = &TreeNode{Val: 1, Left: nil, Right: nil}
	t.Left.Right = &TreeNode{Val: 3, Left: nil, Right: nil}
	t.Right.Left = &TreeNode{Val: 6, Left: nil, Right: nil}
	t.Right.Right = &TreeNode{Val: 9, Left: nil, Right: nil}
	t.Right.Right.Right = &TreeNode{Val: 11, Left: nil, Right: nil}
	return t
}
