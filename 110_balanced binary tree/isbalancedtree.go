package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type TreeNode = common.TreeNode

func main() {
	// vals := []int{3, 9, 20, -1, -1, 15, 7}
	vals := []int{1, 2, 2, 3, 3, -1, -1, 4, 4}
	fmt.Printf("Output: %v\n", isBalanced(common.MakeTree(vals)))
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var depth func(t *TreeNode) int
	depth = func(t *TreeNode) int {
		if t == nil {
			return 0
		}
		d := 1
		leftD := depth(t.Left)
		rightD := depth(t.Right)
		if leftD > rightD {
			d = d + leftD
		} else {
			d = d + rightD
		}
		return d
	}
	ld := depth(root.Left)
	rd := depth(root.Right)
	if ld-rd > 1 || ld-rd < -1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}
