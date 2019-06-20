package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type TreeNode = common.TreeNode

func main() {
	vals := []int{1, 2, 3, -1, 4, 6, 5}
	x, y := 5, 6
	root := common.MakeTree(vals)
	fmt.Printf("Output: %v\n", isCousins(root, x, y))
}

func isCousins(root *TreeNode, x int, y int) bool {
	var depth func(t *TreeNode, n int) (int, int)
	depth = func(t *TreeNode, n int) (int, int) {
		if t == nil {
			return 0, 0
		}
		if t.Val == n {
			return 1, 0
		}
		d, p := depth(t.Left, n)
		if d == 0 {
			d, p = depth(t.Right, n)
		}
		if d == 1 {
			return 2, t.Val
		} else if d > 1 {
			return d + 1, p
		}
		return 0, 0
	}
	dx, px := depth(root, x)
	dy, py := depth(root, y)
	return dx == dy && px != py
}
