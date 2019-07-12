package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type TreeNode = common.TreeNode

func main() {
	vals := []int{3, 9, 20, -1, -1, 15, 7}
	fmt.Println("Output:", levelOrder(common.MakeTree(vals)))
}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :5.9 MB, 在所有 Go 提交中击败了88.03%的用户
func levelOrder(root *TreeNode) [][]int {
	levels := make([][]int, 0)
	if root == nil {
		return levels
	}
	nodes := make([]*TreeNode, 1)
	nodes[0] = root
	for len(nodes) > 0 {
		buf := make([]*TreeNode, 0)
		vals := make([]int, 0)
		for _, n := range nodes {
			vals = append(vals, n.Val)
			if n.Left != nil {
				buf = append(buf, n.Left)
			}
			if n.Right != nil {
				buf = append(buf, n.Right)
			}
		}
		levels = append(levels, vals)
		nodes = buf
	}
	return levels
}
