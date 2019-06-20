package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type TreeNode = common.TreeNode

func main() {
	vals := []int{1, 2, -1, 4, 5, -1, -1, 6, -1, 7}
	fmt.Printf("Output: %d\n", diameterOfBinaryTree(common.MakeTree(vals)))
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lmax := maxDepth(root.Left)
	rmax := maxDepth(root.Right)
	max := lmax + rmax
	var ld, rd int
	if lmax*2 > max {
		ld = diameterOfBinaryTree(root.Left)
	}
	if rmax*2 > max {
		rd = diameterOfBinaryTree(root.Right)
	}
	if ld > max || rd > max {
		if ld > rd {
			return ld
		}
		return rd
	}
	return max
}

func maxDepth(t *TreeNode) int {
	if t == nil {
		return 0
	}
	depth := 1
	leftD := maxDepth(t.Left)
	rightD := maxDepth(t.Right)
	if leftD > rightD {
		depth = depth + leftD
	} else {
		depth = depth + rightD
	}
	return depth
}
