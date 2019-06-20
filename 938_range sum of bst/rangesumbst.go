package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type TreeNode = common.TreeNode

func main() {
	vals := []int{10, 5, 15, 3, 7, 13, 18, 1, -1, 6}
	root := common.MakeTree(vals)
	L, R := 6, 10
	fmt.Printf("Tree: %v, L:%d, R:%d, Sum:%d\n",
		root, L, R, rangeSumBST(root, L, R))
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	sum := rangeSumBST(root.Left, L, R) + rangeSumBST(root.Right, L, R)
	if root.Val >= L && root.Val <= R {
		return sum + root.Val
	}
	return sum
}
