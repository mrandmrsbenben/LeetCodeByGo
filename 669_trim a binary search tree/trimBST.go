package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type TreeNode = common.TreeNode

func main() {
	// input := []int{1, 0, 2}
	// L, R := 1, 2
	// input := []int{2, 1, 3}
	// L, R := 3, 4
	input := []int{3, 0, 4, -1, 2, -1, -1, -1, -1, 1}
	L, R := 1, 3
	root := common.MakeTree(input)
	fmt.Printf("Input:%v Output:%v\n", root, trimBST(root, L, R))
}

func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val >= L && root.Val <= R {
		root.Left = trimBST(root.Left, L, R)
		root.Right = trimBST(root.Right, L, R)
	} else {
		if root.Left == nil && root.Right == nil {
			return nil
		}
		node := trimBST(root.Left, L, R)
		if node != nil {
			root = node
		} else {
			root = trimBST(root.Right, L, R)
		}
	}
	return root
}
