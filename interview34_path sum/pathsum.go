package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

// TreeNode Binary Tree Node
type TreeNode = common.TreeNode

// https://leetcode-cn.com/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/
func main() {
	vals := []int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, -1, -1, 5, 1}
	root := common.MakeTree(vals)
	sum := 22
	fmt.Println("Output: ", pathSum(root, sum))
}

//执行用时 :4 ms, 在所有 Go 提交中击败了94.38%的用户
//内存消耗 :6.3 MB, 在所有 Go 提交中击败了30.53%的用户
func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	var l, r [][]int
	if root.Left == nil && root.Right == nil && root.Val == sum {
		return [][]int{{root.Val}}
	}
	if root.Left != nil {
		l = pathSum(root.Left, sum-root.Val)
		for i := range l {
			l[i] = append([]int{root.Val}, l[i]...)
		}
	}
	if root.Right != nil {
		r = pathSum(root.Right, sum-root.Val)
		for i := range r {
			r[i] = append([]int{root.Val}, r[i]...)
		}
	}
	return append(l, r...)
}
