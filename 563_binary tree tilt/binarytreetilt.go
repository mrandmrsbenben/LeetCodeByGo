package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type TreeNode = common.TreeNode

func main() {
	vals := []int{1, 2, 3, 4, -1, 5, 6}
	fmt.Printf("Output: %d\n", findTilt(common.MakeTree(vals)))
}

func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var fval func(t *TreeNode) int
	fval = func(t *TreeNode) int {
		if t == nil {
			return 0
		}
		return t.Val + fval(t.Left) + fval(t.Right)
	}
	tilt := fval(root.Left) - fval(root.Right)
	if tilt < 0 {
		tilt = -1 * tilt
	}
	return tilt + findTilt(root.Left) + findTilt(root.Right)
}
