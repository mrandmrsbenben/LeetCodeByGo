package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

// TreeNode binary tree node
type TreeNode = common.TreeNode

func main() {
	// vals := []int{1, 2, 3, -1, 5, -1, 4, 6}
	vals := []int{1, -1, 2, -1, 5, 4, 6, 3}
	fmt.Println("Output: ", rightSideView(common.MakeTree2(vals)))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :2.3 MB, 在所有 Go 提交中击败了100.00%的用户
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{root.Val}
	nodes := []*TreeNode{root}
	var buf []*TreeNode
	for len(nodes) > 0 {
		buf = make([]*TreeNode, 0)
		for _, n := range nodes {
			if n.Left != nil {
				buf = append(buf, n.Left)
			}
			if n.Right != nil {
				buf = append(buf, n.Right)
			}
		}
		if len(buf) > 0 {
			res = append(res, buf[len(buf)-1].Val)
		}
		nodes = make([]*TreeNode, len(buf))
		copy(nodes, buf)
	}
	return res
}
