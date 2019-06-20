package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type TreeNode = common.TreeNode

func main() {
	vals := []int{5, 3, 6, 2, 4, -1, 8, 1, -1, -1, -1, -1, -1, 7, 9}
	root := common.MakeTree(vals)
	fmt.Printf("Output: %v\n", increasingBST(root))
}

func increasingBST(root *TreeNode) *TreeNode {
	nodes := make([]*TreeNode, 0)
	var f func(r *TreeNode)
	f = func(r *TreeNode) {
		if r.Left != nil {
			f(r.Left)
		}
		nodes = append(nodes, &TreeNode{r.Val, nil, nil})
		if len(nodes) > 1 {
			nodes[len(nodes)-2].Right = nodes[len(nodes)-1]
		}
		if r.Right != nil {
			f(r.Right)
		}
	}
	f(root)
	return nodes[0]
}
