package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

// TreeNode Binary Tree Node
type TreeNode = common.TreeNode

// https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/
func main() {
	vals := []int{3, 9, 20, -1, -1, 15, 7}
	root := common.MakeTree(vals)
	fmt.Println("Output: ", levelOrder(root))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2.9 MB, 在所有 Go 提交中击败了100.00%的用户
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := [][]int{{root.Val}}
	nodes := []*TreeNode{root}
	var leaves []*TreeNode
	var vals []int
	for len(nodes) > 0 {
		leaves = []*TreeNode{}
		vals = []int{}
		for _, n := range nodes {
			if n.Left != nil {
				leaves = append(leaves, n.Left)
				vals = append(vals, n.Left.Val)
			}
			if n.Right != nil {
				leaves = append(leaves, n.Right)
				vals = append(vals, n.Right.Val)
			}
		}
		if len(vals) > 0 {
			res = append(res, vals)
		}
		nodes = leaves
	}
	return res
}
