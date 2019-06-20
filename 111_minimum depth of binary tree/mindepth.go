package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type TreeNode = common.TreeNode

func main() {
	vals := []int{3, 9, 20, -1, -1, 15, 7}
	fmt.Printf("Output: %d\n", minDepth(common.MakeTree(vals)))
}

// 执行用时 : 12 ms, 在Minimum Depth of Binary Tree的Go提交中击败了94.95% 的用户
// 内存消耗 : 5.4 MB, 在Minimum Depth of Binary Tree的Go提交中击败了21.21% 的用户
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else if root.Left == nil && root.Right == nil {
		return 1
	}
	minLeft := minDepth(root.Left)
	minRight := minDepth(root.Right)
	if minLeft > minRight {
		if minRight > 0 {
			return minRight + 1
		}
		return minLeft + 1
	} else if minLeft > 0 {
		return minLeft + 1
	}
	return minRight + 1
}
