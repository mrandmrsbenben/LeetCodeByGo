package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

// TreeNode Binary Tree Node
type TreeNode = common.TreeNode

// https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/
func main() {
	vals := []int{3, 9, 20, -1, -1, 15, 7}
	root := common.MakeTree(vals)
	fmt.Println("Output: ", levelOrder(root))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2.7 MB, 在所有 Go 提交中击败了100.00%的用户
func levelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{root.Val}
	nodes := []*TreeNode{root}
	leaves := []*TreeNode{}
	for len(nodes) > 0 {
		leaves = []*TreeNode{}
		for _, n := range nodes {
			if n.Left != nil {
				leaves = append(leaves, n.Left)
				res = append(res, n.Left.Val)
			}
			if n.Right != nil {
				leaves = append(leaves, n.Right)
				res = append(res, n.Right.Val)
			}
		}
		nodes = leaves
	}
	return res
}
