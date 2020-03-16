package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

// TreeNode Binary Tree Node
type TreeNode = common.TreeNode

// https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof/
func main() {
	vals := []int{4, 2, 7, 1, 3, 6, 9}
	fmt.Println("Output: ", mirrorTree(common.MakeTree(vals)))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2.1 MB, 在所有 Go 提交中击败了100.00%的用户
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = mirrorTree(root.Right), mirrorTree(root.Left)
	return root
}
