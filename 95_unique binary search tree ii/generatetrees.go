package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

type TreeNode = common.TreeNode

func main() {
	fmt.Println("Output:", generateTrees(3))
}

//执行用时 :48 ms, 在所有 Go 提交中击败了62.07%的用户
//内存消耗 :26 MB, 在所有 Go 提交中击败了34.61%的用户
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return make([]*TreeNode, 0)
	}
	return generate(1, n)
}

func generate(start, end int) []*TreeNode {
	trees := make([]*TreeNode, 0)
	if start > end {
		trees = append(trees, nil)
		return trees
	}
	var root *TreeNode
	var left, right []*TreeNode
	for i := start; i <= end; i++ {
		left = generate(start, i-1)
		right = generate(i+1, end)
		for _, l := range left {
			for _, r := range right {
				root = &TreeNode{i, nil, nil}
				root.Left = l
				root.Right = r
				trees = append(trees, root)
			}
		}
	}
	return trees
}
