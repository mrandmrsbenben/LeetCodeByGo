package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

// TreeNode Tree Node
type TreeNode = common.TreeNode

func main() {
	root := []int{8, 3, 10, 1, 6, -1, 14, -1, -1, 4, 7, 13}
	// root := []int{1, -1, 2, -1, 0, 3}
	// root := []int{8, -1, 1, 5, 6, 2, 4, 0, -1, 7, 3}
	fmt.Println("Output: ", maxAncestorDiff(common.MakeTree2(root)))
}

//执行用时：4 ms, 在所有 Go 提交中击败了83.33%的用户
//内存消耗：4.2 MB, 在所有 Go 提交中击败了35.42%的用户
func maxAncestorDiff(root *TreeNode) int {
	res := 0

	var f func(*TreeNode, int, int)
	f = func(n *TreeNode, minv, maxv int) {

		if n.Left != nil {
			l := max(abs(n.Left.Val-minv), abs(n.Left.Val-maxv))
			if res < l {
				res = l
			}
			f(n.Left, min(minv, n.Left.Val), max(maxv, n.Left.Val))
		}

		if n.Right != nil {
			r := max(abs(n.Right.Val-minv), abs(n.Right.Val-maxv))
			if res < abs(r) {
				res = abs(r)
			}
			f(n.Right, min(minv, n.Right.Val), max(maxv, n.Right.Val))
		}
	}
	f(root, root.Val, root.Val)
	return res
}

func abs(v int) int {
	if v < 0 {
		return v * -1
	}
	return v
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
